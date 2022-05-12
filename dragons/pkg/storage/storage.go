package storage

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const AWS_REGION_NAME = "us-east-1"

var (
	s3sessions = make(map[string]*s3.S3)
)

type Storage struct {
	region string
	s3sess *s3.S3
}

func NewStorage(region string) *Storage {
	if region == "" {
		region = AWS_REGION_NAME
	}
	// Connect Region
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	)
	if err != nil {
		log.Println(err)
	}
	s := s3.New(sess)

	return &Storage{
		region: region,
		s3sess: s,
	}
}

type Bucket struct {
	Name         string
	CreationDate string
}

// GetBuckets retrieves a list of all buckets.
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
// Output:
//     If success, the list of buckets and nil
//     Otherwise, nil and an error from the call to ListBuckets
func (s *Storage) GetBuckets() (b []Bucket, err error) {
	result, err := s.s3sess.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Printf("Error getting buckets: %v", err)
		return nil, err
	}
	for _, bucket := range result.Buckets {
		b = append(b, Bucket{
			Name:         *bucket.Name,
			CreationDate: bucket.CreationDate.Format(time.RFC1123),
		})
	}
	return b, nil
}

func (s *Storage) QueryBucket(bucketName string) (payload []byte, err error) {
	params := &s3.SelectObjectContentInput{
		Bucket:         aws.String(bucketName),
		Key:            aws.String("dragon_stats_one.txt"),
		ExpressionType: aws.String(s3.ExpressionTypeSql),
		Expression:     aws.String("SELECT * FROM S3Object[*][*] s"),
		InputSerialization: &s3.InputSerialization{
			JSON: &s3.JSONInput{
				Type: aws.String("DOCUMENT"),
			},
		},
		OutputSerialization: &s3.OutputSerialization{
			JSON: &s3.JSONOutput{
				RecordDelimiter: aws.String(","),
			},
		},
	}
	resp, err := s.s3sess.SelectObjectContent(params)
	if err != nil {
		return nil, err
	}
	defer resp.EventStream.Close()

	for event := range resp.EventStream.Events() {
		switch v := event.(type) {
		case *s3.RecordsEvent:
			// s3.RecordsEvent.Records is a byte slice of select records
			payload = v.Payload
		case *s3.StatsEvent:
			// s3.StatsEvent contains information on the data that’s processed
			log.Println("Processed", *v.Details.BytesProcessed, "bytes")
		case *s3.EndEvent:
			// s3.EndEvent
			log.Println("SelectObjectContent completed")
		}
	}

	if err := resp.EventStream.Err(); err != nil {
		log.Printf("failed to read from SelectObjectContent EventStream, %v", err)
		return nil, err
	}
	return payload, nil
}
