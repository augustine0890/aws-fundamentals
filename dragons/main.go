package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// S3ListBucketsAPI
type S3ListBucketsAPI interface {
	ListBuckets(ctx context.Context,
		params *s3.ListBucketsInput,
		optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

// GetAllBuckets retrieves all buckets
// Inputs:
//  c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
// Outputs:
//	ListBucketsOutput object containing the result of the service call and nil.
func GetAllBuckets(c context.Context, api S3ListBucketsAPI, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return api.ListBuckets(c, input)
}

func main() {
	ctx := context.TODO()
	// Load the shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1")) // specify region
	if err != nil {
		log.Fatalf("failed to load AWS configuration: %v", err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// create new context with a timeout, e.g. 10 seconds
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	input := &s3.ListBucketsInput{}
	// List buckets
	result, err := GetAllBuckets(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving buckets: ", err)
		fmt.Println(err)
		return
	}

	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name + ": " + bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
	}

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("dragons-app-20220417"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("The First Page Results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}

	params := &s3.SelectObjectContentInput{
		Bucket:         aws.String("dragons-app-20220417"),
		Key:            aws.String("dragon_stats_one.txt"),
		ExpressionType: "SQL",
		Expression:     aws.String("SELECT * FROM S3Object[*][*] s"),
		InputSerialization: &types.InputSerialization{
			JSON: &types.JSONInput{
				Type: "DOCUMENT",
			},
		},
		OutputSerialization: &types.OutputSerialization{
			JSON: &types.JSONOutput{
				RecordDelimiter: aws.String(","),
			},
		},
	}

	resp, err := client.SelectObjectContent(context.TODO(), params)
	if err != nil {
		fmt.Println("Got an error selecting object:")
		fmt.Println(err)
		return
	}
	defer resp.GetStream().Close()

	for evt := range resp.GetStream().Events() {
		switch v := evt.(type) {
		case *types.RecordsEvent:
			fmt.Println(v)
		}
	}
}
