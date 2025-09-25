package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Gateway.
// Experimental.
type IGatewayRef interface {
	constructs.IConstruct
	// A reference to a Gateway resource.
	// Experimental.
	GatewayRef() *GatewayReference
}

// The jsii proxy for IGatewayRef
type jsiiProxy_IGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGatewayRef) GatewayRef() *GatewayReference {
	var returns *GatewayReference
	_jsii_.Get(
		j,
		"gatewayRef",
		&returns,
	)
	return returns
}

