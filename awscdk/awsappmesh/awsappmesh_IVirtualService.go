package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappmesh/internal"
)

// Represents the interface which all VirtualService based classes MUST implement.
// Experimental.
type IVirtualService interface {
	awscdk.IResource
	// The Mesh which the VirtualService belongs to.
	// Experimental.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the virtual service.
	// Experimental.
	VirtualServiceArn() *string
	// The name of the VirtualService.
	// Experimental.
	VirtualServiceName() *string
}

// The jsii proxy for IVirtualService
type jsiiProxy_IVirtualService struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IVirtualService) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualService) VirtualServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualService) VirtualServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualServiceName",
		&returns,
	)
	return returns
}

