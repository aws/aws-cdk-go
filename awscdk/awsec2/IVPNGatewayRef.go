package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPNGateway.
// Experimental.
type IVPNGatewayRef interface {
	constructs.IConstruct
	// A reference to a VPNGateway resource.
	// Experimental.
	VpnGatewayRef() *VPNGatewayReference
}

// The jsii proxy for IVPNGatewayRef
type jsiiProxy_IVPNGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPNGatewayRef) VpnGatewayRef() *VPNGatewayReference {
	var returns *VPNGatewayReference
	_jsii_.Get(
		j,
		"vpnGatewayRef",
		&returns,
	)
	return returns
}

