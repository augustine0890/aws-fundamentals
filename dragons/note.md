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

## Install AWS SDK Go [V1](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html)
- [API](https://docs.aws.amazon.com/sdk-for-go/api/index.html) Referance V1

Reference: AWS [Code](https://docs.aws.amazon.com/code-samples/latest/catalog/code-catalog-go.html) Sample.
## API Authentication
- Amazon [Cognito](https://aws-tc-largeobjects.s3.amazonaws.com/DEV-AWS-MO-BuildingRedux/node-exercise-3-cognito.html) Authentication
- There are two main components of Amazon Cognito:
  - User Pools
  - Identity Pools
- User pools are user directories that probide sign-up and sign-in options for your app users.
- Identity pools enable you to grant your users temporay AWS credentials to access AWS services, such as Amazon S3 or DynamoDB.
- Use Cases:
  - Access resourses with API Gateway and Lambda with a User Pool. API Gateway validate the tokens from a successful user pool authentication, and uses them to grant your users access to resourses including Lambda functions, or your own API.
  - After a successful user pool authentication, your app will receive user pool tokens from Amazon Cognito. You can exchange them for temporay access to other AWS services with an identity pool.
- Temp [mail](https://tempmailo.com/) use cases for creating user (testing)
- Secure Password [Generator](https://passwordsgenerator.net/)
## Serverless Computing and Lamda
