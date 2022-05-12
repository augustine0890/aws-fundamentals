package storage

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const AWS_REGION_NAME = "us-east-1"

var (
	sessions   = make(map[string]*session.Session)
	s3sessions = make(map[string]*s3.S3)
)

type Storage struct {
	region string
}

func NewStorage() *Storage {
	return &Storage{
		region: AWS_REGION_NAME,
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
	result, err := gets3clientRegion("").ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Printf("Error getting buckets: %v", err)
		return nil, err
	}
	for _, bucket := range result.Buckets {
		b = append(b, Bucket{
			Name:         *bucket.Name,
			CreationDate: bucket.CreationDate.Format("2016-01-02 15:04:05 Monday"),
		})
	}
	return b, nil
}

func connectRegion(region string) *session.Session {
	if region == "" {
		region = AWS_REGION_NAME
	}

	if val, ok := sessions[region]; ok {
		return val
	} else {
		sess, err := session.NewSession(
			&aws.Config{
				Region: aws.String(region),
			},
		)
		if err != nil {
			log.Println(err)
		}
		sessions[region] = sess
		return sess
	}
}

func gets3clientRegion(region string) *s3.S3 {
	if region == "" {
		region = AWS_REGION_NAME
	}
	if val, ok := s3sessions[region]; ok {
		return val
	} else {
		s := s3.New(connectRegion(region))
		s3sessions[region] = s
		return s
	}
}
