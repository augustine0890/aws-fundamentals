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
### Exercise 4: AWS Lambda Functions
- AWS Lambda Functions [Exercise](https://aws-tc-largeobjects.s3.amazonaws.com/DEV-AWS-MO-BuildingRedux/node-exercise-4-lambda.html)
- In this exercise, you deploy the `ListDragon`, `AddDragon`, and `ValidateDragon` Lambda functions. After you implement the functions, you then update the `GET` method in the application programming interface (API) to use the Lambda function.
- Save the ARNs for the Lambda roles in environment variables:
  - `export ROLE_ARN_READWRITE=aws iam get-role --role-name dragons-readwrite-lambda-role  --query 'Role.Arn' --output text`
  -`export ROLE_ARN_READ=aws iam get-role --role-name dragons-read-lambda-role  --query 'Role.Arn' --output text`

### Exercise 5: AWS Step Functions
- AWS Step Functions [Exercise](https://aws-tc-largeobjects.s3.amazonaws.com/DEV-AWS-MO-BuildingRedux/node-exercise-5-stepfunctions.html)
- List the ARNs:
  - `aws lambda list-functions --query 'Functions[].FunctionArn'`
- Step Functions
  - ARN of the new state machine:
    - `arn:aws:iam::427591574411:role/service-role/StepFunctions-DragonsStateMachine-role-a700c3ef`
- Role in IAM (role's ARN)
  - `arn:aws:iam::427591574411:role/dragons-apig-role`
- Update the POST (REST API)
  - Remove mock integration
  - Connect the method to a Step Functions workflow for the Lambda functions
- Configure the mapping template
  - The mapping template converts the `POST` body that the application sends into a body that Step Functions `StartExecution` method understands.
- Re-enabling CORS and deploying the API