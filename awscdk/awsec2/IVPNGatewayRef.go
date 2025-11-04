package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPNGateway.
// Experimental.
type IVPNGatewayRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VPNGateway resource.
	// Experimental.
	VpnGatewayRef() *VPNGatewayReference
}

// The jsii proxy for IVPNGatewayRef
type jsiiProxy_IVPNGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IVPNGatewayRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPNGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

