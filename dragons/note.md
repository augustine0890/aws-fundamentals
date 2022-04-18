## Simple Storage Service
- Create a bucket
  - `aws s3 mb <target> [--options]`
  - `aws s3 mb s3://dragons-app-20220417`
- List buckets and objects
  - `aws s3 ls <target> [--options]`
  - `aws s3 ls`
- Delete buckets
  - `aws s3 rb <target> [--options]`
  - `aws s3 rb s3://dragons-app-20220417`

- High-level (s3) commands with AWS [CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-services-s3-commands.html)

## Install the AWS SDK for Go V2
- Core SDK module
  - `go get github.com/aws/aws-sdk-go-v2`
- The config module
  - `go get github.com/aws/aws-sdk-go-v2/config`
- Amazon S3 API client
  - `go get github.com/aws/aws-sdk-go-v2/service/s3`