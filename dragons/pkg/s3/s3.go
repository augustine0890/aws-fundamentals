package s3

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

type Bucket struct {
	Name string
}

func GetBuckets() (b Bucket) {
	return Bucket{}
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