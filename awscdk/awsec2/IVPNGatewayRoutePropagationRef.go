package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPNGatewayRoutePropagation.
// Experimental.
type IVPNGatewayRoutePropagationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VPNGatewayRoutePropagation resource.
	// Experimental.
	VpnGatewayRoutePropagationRef() *VPNGatewayRoutePropagationReference
}

// The jsii proxy for IVPNGatewayRoutePropagationRef
type jsiiProxy_IVPNGatewayRoutePropagationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IVPNGatewayRoutePropagationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPNGatewayRoutePropagationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

