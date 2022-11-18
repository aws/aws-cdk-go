package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalaccelerator/internal"
)

// The interface of the EndpointGroup.
// Experimental.
type IEndpointGroup interface {
	awscdk.IResource
	// EndpointGroup ARN.
	// Experimental.
	EndpointGroupArn() *string
}

// The jsii proxy for IEndpointGroup
type jsiiProxy_IEndpointGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEndpointGroup) EndpointGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointGroupArn",
		&returns,
	)
	return returns
}

