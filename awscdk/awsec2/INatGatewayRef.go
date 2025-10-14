package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NatGateway.
// Experimental.
type INatGatewayRef interface {
	constructs.IConstruct
	// A reference to a NatGateway resource.
	// Experimental.
	NatGatewayRef() *NatGatewayReference
}

// The jsii proxy for INatGatewayRef
type jsiiProxy_INatGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INatGatewayRef) NatGatewayRef() *NatGatewayReference {
	var returns *NatGatewayReference
	_jsii_.Get(
		j,
		"natGatewayRef",
		&returns,
	)
	return returns
}

