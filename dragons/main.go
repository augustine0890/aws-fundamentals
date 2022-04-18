package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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
}
