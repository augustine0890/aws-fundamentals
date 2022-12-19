package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

func main() {
	// Create a Session with a custom region
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	// Create a Comprehend client from Session
	client := comprehend.New(sess)

	text := "Hello Zhang Wei, I am John. Your AnyCompany Financial Services, LLC credit card account 1111-0000-1111-0008 has a minimum payment of $24.53 that is due by July 31st. Based on your autopay settings, we will withdraw your payment on the due date from your bank account number XXXXXX1111 with the routing number XXXXX0000."

	params := comprehend.DetectSentimentInput{
		LanguageCode: aws.String("en"),
		Text:         aws.String(text),
	}

	req, resp := client.DetectSentimentRequest(&params)
	err := req.Send()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)

	{
		params := comprehend.DetectKeyPhrasesInput{}
		params.SetLanguageCode("en")
		params.SetText(text)

		req, resp := client.DetectKeyPhrasesRequest(&params)

		err := req.Send()
		if err == nil { // resp is now filled
			// fmt.Println(*resp.Sentiment)
			for _, s := range resp.KeyPhrases {
				if *s.Score >= 0.95 {
					fmt.Println(s)
				}
			}
			// fmt.Println(*resp)
		} else {
			fmt.Println(err)
		}
	}
}
