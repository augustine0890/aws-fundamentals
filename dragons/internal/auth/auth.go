package auth

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
)

type Auth struct {
	CognitoClient   *cognito.CognitoIdentityProvider
	UserPoolID      string
	AppClientID     string
	AppClientSecret string
}

func NewAuth(region string) *Auth {
	if region == "" {
		region = "us-east-1"
	}

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		},
	)
	if err != nil {
		log.Printf("Error creating session: %v", err)
	}
	s := cognito.New(sess)

	err = godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading variables environment: %v", err)
	}

	return &Auth{
		CognitoClient:   s,
		UserPoolID:      os.Getenv("COGNITO_USER_POOL_ID"),
		AppClientID:     os.Getenv("COGNITO_APP_CLIENT_ID"),
		AppClientSecret: os.Getenv("COGNITO_APP_CLIENT_SECRET"),
	}
}
