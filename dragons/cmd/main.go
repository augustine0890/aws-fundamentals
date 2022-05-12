package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"dragons/pkg/storage"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Application struct {
	storage *storage.Storage
}

func ReadDragons(sess *session.Session) {
	svc := s3.New(sess)
	params := &s3.SelectObjectContentInput{
		Bucket:         aws.String("dragons-app-20220417"),
		Key:            aws.String("dragon_stats_one.txt"),
		ExpressionType: aws.String(s3.ExpressionTypeSql),
		Expression:     aws.String("SELECT * FROM S3Object[*][*] s"),
		InputSerialization: &s3.InputSerialization{
			JSON: &s3.JSONInput{
				Type: aws.String("DOCUMENT"),
			},
		},
		OutputSerialization: &s3.OutputSerialization{
			JSON: &s3.JSONOutput{
				RecordDelimiter: aws.String(","),
			},
		},
	}
	resp, err := svc.SelectObjectContent(params)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.EventStream.Close()

	for event := range resp.EventStream.Events() {
		switch v := event.(type) {
		case *s3.RecordsEvent:
			// s3.RecordsEvent.Records is a byte slice of select records
			fmt.Println("Results: ", string(v.Payload))
		case *s3.StatsEvent:
			// s3.StatsEvent contains information on the data that’s processed
			fmt.Println("Processed", *v.Details.BytesProcessed, "bytes")
		case *s3.EndEvent:
			// s3.EndEvent
			fmt.Println("SelectObjectContent completed")
		}
	}

	if err := resp.EventStream.Err(); err != nil {
		fmt.Printf("failed to read from SelectObjectContent EventStream, %v", err)
	}
}

func main() {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
		// SharedConfigState: session.SharedConfigEnable,
		// Config: aws.Config{
			// Region: aws.String("us-east-1"),
		// },
	// }))

	// ReadDragons(sess)
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	app := &Application{
		storage: storage.NewStorage(""),
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
