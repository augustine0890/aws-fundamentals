package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetAllBuckets retrieves a list of all buckets.
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
// Output:
//     If success, the list of buckets and nil
//     Otherwise, nil and an error from the call to ListBuckets
func GetAllBuckets(sess *session.Session) (*s3.ListBucketsOutput, error) {
	svc := s3.New(sess)

	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ReadDragons(sess *session.Session) {
	svc := s3.New(sess)
	params := &s3.SelectObjectContentInput{
		Bucket:         aws.String("dragons-app-20220417"),
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
	resp, err := svc.SelectObjectContent(params)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.EventStream.Close()

	for event := range resp.EventStream.Events() {
		switch v := event.(type) {
		case *s3.RecordsEvent:
			// s3.RecordsEvent.Records is a byte slice of select records
			fmt.Println("Results: ", string(v.Payload))
		case *s3.StatsEvent:
			// s3.StatsEvent contains information on the data that’s processed
			fmt.Println("Processed", *v.Details.BytesProcessed, "bytes")
		case *s3.EndEvent:
			// s3.EndEvent
			fmt.Println("SelectObjectContent completed")
		}
	}

	if err := resp.EventStream.Err(); err != nil {
		fmt.Printf("failed to read from SelectObjectContent EventStream, %v", err)
	}
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	}))

	result, err := GetAllBuckets(sess)
	if err != nil {
		fmt.Printf("Error getting buckets: %v", err)
		return
	}

	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name + " " + bucket.CreationDate.Format("2016-01-02 15:04:05 Monday"))
	}

	ReadDragons(sess)
}
