package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

	awscdkapigw "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	awsapigwintegrations "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2"
	awscdklambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type TestGoCdkStackProps struct {
	awscdk.StackProps
}

func NewTestGoCdkStack(scope constructs.Construct, id string, props *TestGoCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	echoLambdaFunction := awscdklambdago.NewGoFunction(stack, jsii.String("EchoFunc"), &awscdklambdago.GoFunctionProps{
		FunctionName: jsii.String("EchoFunc"),
		Description:  jsii.String("an apigw handler that returns IP and User-Agent as JSON"),
		Entry:        jsii.String("../lambda-func/echo"),
	})

	echoApi := awscdkapigw.NewHttpApi(stack, jsii.String("EchoApi"), nil)

	echoApi.AddRoutes(&awscdkapigw.AddRoutesOptions{
		Path:        jsii.String("/"),
		Methods:     &[]awscdkapigw.HttpMethod{awscdkapigw.HttpMethod_GET},
		Integration: awsapigwintegrations.NewHttpLambdaIntegration(jsii.String("EchoApiIntegration"), echoLambdaFunction, nil),
	})

	awscdk.NewCfnOutput(stack, jsii.String("EchoApiURL"), &awscdk.CfnOutputProps{
		Value:       echoApi.ApiEndpoint(),
		Description: jsii.String("the URL to echoApi"),
		ExportName:  jsii.String("echoApiURL"),
	})

	// awss3.NewBucket(stack, jsii.String("MyFirstBucket"), &awss3.BucketProps{
	// 	Versioned:         jsii.Bool(true),
	// 	RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	// 	AutoDeleteObjects: jsii.Bool(true),
	// })
	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("TestGoCdkQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewTestGoCdkStack(app, "TestGoCdkStack", &TestGoCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
