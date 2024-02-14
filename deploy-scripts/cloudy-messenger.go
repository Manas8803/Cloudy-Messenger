package main

import (
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
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

	//* SNS service for sending emails to subscribers
	sns_topic := awssns.NewTopic(stack, jsii.String("SnsTopic"), &awssns.TopicProps{
		TopicName:   jsii.String("CloudyMessengerTopic"),
		DisplayName: jsii.String("Cloudy Messenger Notifications"),
	})
	emailAddresses := []string{"manasahavegeta@gmail.com", "sahamm@rknec.edu", "mishraas_3@rknec.edu", "manassaha0803@gmail.com"}

	//* Subscribe each email address
	for _, email := range emailAddresses {
		awssns.NewSubscription(stack, jsii.String("EmailSubscription-"+email), &awssns.SubscriptionProps{
			Topic:    sns_topic,
			Protocol: awssns.SubscriptionProtocol_EMAIL,
			Endpoint: jsii.String(email),
		})
	}

	//* Lambda function - notfiy service
	notify_handler := awslambda.NewFunction(stack, jsii.String("Notify-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("notify-service.zip"), nil),
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("/notify-service/build/main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(6)),
		Environment: &map[string]*string{
			"SNS_TOPIC_ARN": jsii.String(*sns_topic.TopicArn()),
		},
	})

	//* Lambda function - fetch service
	awslambda.NewFunction(stack, jsii.String("fetch-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("fetch-service.zip"), nil),
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("/fetch-service/build/main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(6)),
		Environment: &map[string]*string{
			"NOTIFY_FUNC_ARN": jsii.String(*notify_handler.FunctionArn()),
		},
	})

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
