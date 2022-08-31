package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// This type of integration lets an API expose AWS service actions.
//
// It is
// intended for calling all AWS service actions, but is not recommended for
// calling a Lambda function, because the Lambda custom integration is a legacy
// technology.
//
// Example:
//   getMessageIntegration := apigateway.NewAwsIntegration(&awsIntegrationProps{
//   	service: jsii.String("sqs"),
//   	path: jsii.String("queueName"),
//   	region: jsii.String("eu-west-1"),
//   })
//
// Experimental.
type AwsIntegration interface {
	Integration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	// Experimental.
	Bind(method Method) *IntegrationConfig
}

// The jsii proxy struct for AwsIntegration
type jsiiProxy_AwsIntegration struct {
	jsiiProxy_Integration
}

// Experimental.
func NewAwsIntegration(props *AwsIntegrationProps) AwsIntegration {
	_init_.Initialize()

	j := jsiiProxy_AwsIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigateway.AwsIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIntegration_Override(a AwsIntegration, props *AwsIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.AwsIntegration",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AwsIntegration) Bind(method Method) *IntegrationConfig {
	var returns *IntegrationConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{method},
		&returns,
	)

	return returns
}

