package main

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/tugberkugurlu/event-sourcing-dynamodb-go"
)

func main() {
	cfg := aws.Config{}
	cfg.Region = aws.String("eu-west-2")
	cfg.Endpoint = aws.String("http://localhost:8000")
	cfg.Credentials = credentials.NewStaticCredentials("test", "test", "")
	sess := session.Must(session.NewSession())
	
	err := esdynamodb.EnsureInitialized(sess, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
}