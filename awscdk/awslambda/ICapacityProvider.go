package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
)

// Represents a Lambda capacity provider.
type ICapacityProvider interface {
	awscdk.IResource
	// The ARN of the capacity provider.
	CapacityProviderArn() *string
	// The name of the capacity provider.
	CapacityProviderName() *string
}

// The jsii proxy for ICapacityProvider
type jsiiProxy_ICapacityProvider struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ICapacityProvider) CapacityProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityProvider) CapacityProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityProviderName",
		&returns,
	)
	return returns
}

