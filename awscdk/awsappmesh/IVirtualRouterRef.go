package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualRouter.
// Experimental.
type IVirtualRouterRef interface {
	constructs.IConstruct
	// A reference to a VirtualRouter resource.
	// Experimental.
	VirtualRouterRef() *VirtualRouterReference
}

// The jsii proxy for IVirtualRouterRef
type jsiiProxy_IVirtualRouterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVirtualRouterRef) VirtualRouterRef() *VirtualRouterReference {
	var returns *VirtualRouterReference
	_jsii_.Get(
		j,
		"virtualRouterRef",
		&returns,
	)
	return returns
}

