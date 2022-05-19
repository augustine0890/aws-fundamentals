package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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

// Secret hash is not a client secret itself, but a base64 encoded hmac-sha256 hash.
// The actual AWS documentation on how to compute this hash is here:
// https://docs.aws.amazon.com/cognito/latest/developerguide/signing-up-users-in-your-app.html#cognito-user-pools-computing-secret-hash
func computeSecretHash(clientSecret string, username string, clientId string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(username + clientId))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (a *Auth) Register(u *User) (string, error) {
	// Compute secret hash based on client secret
	secretHash := computeSecretHash(a.AppClientSecret, u.Email, a.AppClientID)

	user := &cognito.SignUpInput{
		Username:   aws.String(u.Email),
		Password:   aws.String(u.Password),
		ClientId:   aws.String(a.AppClientID),
		SecretHash: aws.String(secretHash),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(u.Email),
			},
		},
	}

	result, err := a.CognitoClient.SignUp(user)
	if err != nil {
		log.Printf("Cognito SignUp Error - %v", err)
		return "", err
	}
	return result.String(), nil
}

func (a *Auth) ConfirmSignUp(uc *UserConfirm) (string, error) {
	secretHash := computeSecretHash(a.AppClientSecret, uc.Email, a.AppClientID)

	input := &cognito.ConfirmSignUpInput{
		Username:         aws.String(uc.Email),
		ConfirmationCode: aws.String(uc.ConfirmationCode),
		ClientId:         aws.String(a.AppClientID),
		SecretHash:       aws.String(secretHash),
	}

	result, err := a.CognitoClient.ConfirmSignUp(input)
	if err != nil {
		log.Printf("Cognito ConfirmSignUp Error - %v", err)
		return "", err
	}

	return result.String(), nil
}

func (a *Auth) Login(u *User) (string, error) {
	secretHash := computeSecretHash(a.AppClientSecret, u.Email, a.AppClientID)
	params := map[string]*string{
		"USERNAME": aws.String(u.Email),
		"PASSWORD": aws.String(u.Password),
		"SECRET_HASH": aws.String(secretHash),
	}

	input := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: params,
		ClientId: aws.String(a.AppClientID),
	}
	auth, err := a.CognitoClient.InitiateAuth(input)
	if err != nil {
		log.Printf("Cognito Login Error - %v", err)
		return "", err
	}
	u.Token = *auth.AuthenticationResult.AccessToken
	return u.Token, nil
}