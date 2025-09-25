package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
)

// Integrates an AWS Sagemaker Endpoint to an API Gateway method.
//
// Example:
//   var resource resource
//   var endpoint iEndpoint
//
//   resource.AddMethod(jsii.String("POST"), apigateway.NewSagemakerIntegration(endpoint))
//
type SagemakerIntegration interface {
	AwsIntegration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(method Method) *IntegrationConfig
}

// The jsii proxy struct for SagemakerIntegration
type jsiiProxy_SagemakerIntegration struct {
	jsiiProxy_AwsIntegration
}

func NewSagemakerIntegration(endpoint awssagemaker.IEndpoint, options *SagemakerIntegrationOptions) SagemakerIntegration {
	_init_.Initialize()

	if err := validateNewSagemakerIntegrationParameters(endpoint, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.SagemakerIntegration",
		[]interface{}{endpoint, options},
		&j,
	)

	return &j
}

func NewSagemakerIntegration_Override(s SagemakerIntegration, endpoint awssagemaker.IEndpoint, options *SagemakerIntegrationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.SagemakerIntegration",
		[]interface{}{endpoint, options},
		s,
	)
}

func (s *jsiiProxy_SagemakerIntegration) Bind(method Method) *IntegrationConfig {
	if err := s.validateBindParameters(method); err != nil {
		panic(err)
	}
	var returns *IntegrationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{method},
		&returns,
	)

	return returns
}

