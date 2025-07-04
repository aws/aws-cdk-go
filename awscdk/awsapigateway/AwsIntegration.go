package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
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
//   getMessageIntegration := apigateway.NewAwsIntegration(&AwsIntegrationProps{
//   	Service: jsii.String("sqs"),
//   	Path: jsii.String("queueName"),
//   	Region: jsii.String("eu-west-1"),
//   })
//
type AwsIntegration interface {
	Integration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(method Method) *IntegrationConfig
}

// The jsii proxy struct for AwsIntegration
type jsiiProxy_AwsIntegration struct {
	jsiiProxy_Integration
}

func NewAwsIntegration(props *AwsIntegrationProps) AwsIntegration {
	_init_.Initialize()

	if err := validateNewAwsIntegrationParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AwsIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAwsIntegration_Override(a AwsIntegration, props *AwsIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AwsIntegration",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AwsIntegration) Bind(method Method) *IntegrationConfig {
	if err := a.validateBindParameters(method); err != nil {
		panic(err)
	}
	var returns *IntegrationConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{method},
		&returns,
	)

	return returns
}

