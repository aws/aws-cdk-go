package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualGateway.
// Experimental.
type IVirtualGatewayRef interface {
	constructs.IConstruct
	// A reference to a VirtualGateway resource.
	// Experimental.
	VirtualGatewayRef() *VirtualGatewayReference
}

// The jsii proxy for IVirtualGatewayRef
type jsiiProxy_IVirtualGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVirtualGatewayRef) VirtualGatewayRef() *VirtualGatewayReference {
	var returns *VirtualGatewayReference
	_jsii_.Get(
		j,
		"virtualGatewayRef",
		&returns,
	)
	return returns
}

