package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Integrates an AWS Lambda function to an API Gateway method.
//
// Example:
//   var resource resource
//   var handler function
//
//   resource.AddMethod(jsii.String("GET"), apigateway.NewLambdaIntegration(handler))
//
// Experimental.
type LambdaIntegration interface {
	AwsIntegration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	// Experimental.
	Bind(method Method) *IntegrationConfig
}

// The jsii proxy struct for LambdaIntegration
type jsiiProxy_LambdaIntegration struct {
	jsiiProxy_AwsIntegration
}

// Experimental.
func NewLambdaIntegration(handler awslambda.IFunction, options *LambdaIntegrationOptions) LambdaIntegration {
	_init_.Initialize()

	if err := validateNewLambdaIntegrationParameters(handler, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigateway.LambdaIntegration",
		[]interface{}{handler, options},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaIntegration_Override(l LambdaIntegration, handler awslambda.IFunction, options *LambdaIntegrationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.LambdaIntegration",
		[]interface{}{handler, options},
		l,
	)
}

func (l *jsiiProxy_LambdaIntegration) Bind(method Method) *IntegrationConfig {
	if err := l.validateBindParameters(method); err != nil {
		panic(err)
	}
	var returns *IntegrationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{method},
		&returns,
	)

	return returns
}

