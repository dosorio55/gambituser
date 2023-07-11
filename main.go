package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/dosorio55/gambituser/awsgo"
	"github.com/dosorio55/gambituser/db"
	"github.com/dosorio55/gambituser/models"
)

func main() {
	// Create a new instance
	lambda.Start(ExecutingLambda)
}

func ExecutingLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	// Your code here
	awsgo.InitialiceAws()

	// Get the secret from AWS Secrets Manager
	if !ValidateParams() {
		fmt.Println("Missing parameter SecretName")
		error := errors.New("Missing parameter SecretName")
		return event, error
	}

	var data models.SignUp

	// data.Username = event.Request.UserAttributes["email"]
	// data.UserUUID = event.Request.UserAttributes["sub"]

	// fmt.Println("data: ", data)

	for key, value := range event.Request.UserAttributes {
		switch key {
		case "email":
			data.Username = value
		case "sub":
			data.UserUUID = value
		}
	}

	err := db.ReadSecret()

	if err != nil {
		fmt.Println("Error reading secret: ", err)
		return event, err
	}

	err = db.SignUp(data)

	return event, err
}

func ValidateParams() bool {
	_, getParam := os.LookupEnv("SecretName")
	return getParam
}
