package main

import (
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

type CloudyMessengerStackProps struct {
	awscdk.StackProps
}

func NewCloudyMessengerStack(scope constructs.Construct, id string, props *CloudyMessengerStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	// Lambda function - fetch service
	// Lambda function - notfiy service
	// SNS service for sending emails to subscribers
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCloudyMessengerStack(app, "CloudyMessengerStack", &CloudyMessengerStackProps{
		awscdk.StackProps{
			StackName: jsii.String("CloudyMessengerStack"),
			Env:       env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file : ", err)
	}

	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
