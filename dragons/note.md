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
- API Gateway: 401 [Unauthorized](https://stackoverflow.com/questions/50778387/api-gateway-cognito-user-pool-authorizer-401-unauthorized?rq=1)
- The ID token contains claims about the identity of the authenticated user, such as name and email. The Access token contains claims about the authenticated user, a list of the user's groups, and a list of scopes.
- Using [tokens](https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html) with user pools. More [here](https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-the-id-token.html)
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

### Exercise 6: Optimizing Application
- Optimizing the [Application](https://aws-tc-largeobjects.s3.amazonaws.com/DEV-AWS-MO-BuildingRedux/node-exercise-6-optimizing.html)
- Add X-Ray to Lambda function
  - Create `.zip` package: `zip -r nodeListDragonsFunction.zip .`
  - Update `ListDragons` function: `aws lambda update-function-code  --function-name ListDragons --zip-file fileb://nodeListDragonsFunction.zip`

## The Serverless Framework
- Install the `serverless`: `npm install -g serverless`
- Go template
  - `serverless create -t aws-go-dep -p myservice`

## AWS Serverless Application Model (AWS SAM)
- The AWS Serverless Application Model (AWS SAM) is an open-source framework that you can use to build serverless applications on AWS.
- A serverless application is a combination of Lambda functions, event sources, and other resources that work together to perform tasks.
- Follow these [instructions](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html) to install the latest version of the AWS SAM CLI.
- The Makefile gives you three basic commands you will use to compile and package your Lambda function:
  - `make clean`: remove the compiled executable and zipped package
  - `make build`: compile the program to run in the Lambda go1.x runtime.
  - `make package`: zip the compiled function for upload to the Management Console