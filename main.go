package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file to load, skipping.")
	}

	cfg := aws.NewConfig()
	cfg.WithCredentialsChainVerboseErrors(true)

	session := session.Must(session.NewSession(cfg))
	awssts := sts.New(session)

	input := &sts.GetCallerIdentityInput{}
	result, err := awssts.GetCallerIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Fatal(aerr.Error())
			}
		} else {
			log.Fatal(err.Error())
		}
	} else {
		log.Print(result)
	}

}
