package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPNGatewayRoutePropagation.
// Experimental.
type IVPNGatewayRoutePropagationRef interface {
	constructs.IConstruct
	// A reference to a VPNGatewayRoutePropagation resource.
	// Experimental.
	VpnGatewayRoutePropagationRef() *VPNGatewayRoutePropagationReference
}

// The jsii proxy for IVPNGatewayRoutePropagationRef
type jsiiProxy_IVPNGatewayRoutePropagationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPNGatewayRoutePropagationRef) VpnGatewayRoutePropagationRef() *VPNGatewayRoutePropagationReference {
	var returns *VPNGatewayRoutePropagationReference
	_jsii_.Get(
		j,
		"vpnGatewayRoutePropagationRef",
		&returns,
	)
	return returns
}

