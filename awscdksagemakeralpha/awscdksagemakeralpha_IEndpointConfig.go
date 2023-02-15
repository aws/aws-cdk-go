// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/internal"
)

// The interface for a SageMaker EndpointConfig resource.
// Experimental.
type IEndpointConfig interface {
	awscdk.IResource
	// The ARN of the endpoint configuration.
	// Experimental.
	EndpointConfigArn() *string
	// The name of the endpoint configuration.
	// Experimental.
	EndpointConfigName() *string
}

// The jsii proxy for IEndpointConfig
type jsiiProxy_IEndpointConfig struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEndpointConfig) EndpointConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEndpointConfig) EndpointConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointConfigName",
		&returns,
	)
	return returns
}

