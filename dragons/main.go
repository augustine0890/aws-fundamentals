package main

import (
	"fmt"

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

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
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
}
