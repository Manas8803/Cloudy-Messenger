package roles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/jsii-runtime-go"
)

func CreateInvocationRole(stack awscdk.Stack, notify_handler awslambda.Function) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("Invoke-Role"), &awsiam.RoleProps{
    AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
    Actions:   &[]*string{jsii.String("lambda:InvokeFunction")},
    Resources: &[]*string{jsii.String(*notify_handler.FunctionArn())},
	}))

	return role
}
func CreatePublishRole(stack awscdk.Stack, sns_topic awssns.Topic) awsiam.Role {

	role := awsiam.NewRole(stack, jsii.String("SNSPublishRole"), &awsiam.RoleProps{
    AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
  })

  // Grant permission to publish to the specific SNS topic
  role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
    Actions:   &[]*string{jsii.String("sns:Publish")},
    Resources: &[]*string{jsii.String(*sns_topic.TopicArn())},
  }))

	return role
}

func CreateScheduler_InvokeRole(stack awscdk.Stack, invoke_handler awslambda.Function) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("IAMRoleForExecutingLambdaFunction"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("scheduler.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
		RoleName:  jsii.String("scheduler-role-for-executing-lambda-function"),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("lambda:InvokeFunction"),
		},
		Resources: &[]*string{
			invoke_handler.FunctionArn(),
		},
	}))

	return role
}