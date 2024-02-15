package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the SNS topic ARN (replace with your actual ARN)
	topicArn := os.Getenv("SNS_TOPIC_ARN")

	// Create a new SNS client
	sess, err := session.NewSession()
	if err != nil {
		log.Println("Error in creating session : ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}
	svc := sns.New(sess)

	// Prepare message input
	message := "This is a test message"
	input := &sns.PublishInput{
		TopicArn: aws.String(topicArn),
		Message:  aws.String(message),
	}

	// Publish the message to the topic
	_, err = svc.Publish(input)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Notifications sent successfully"}, nil
}

func main() {
	lambda.Start(Handler)
}
