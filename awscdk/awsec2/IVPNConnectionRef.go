package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPNConnection.
// Experimental.
type IVPNConnectionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VPNConnection resource.
	// Experimental.
	VpnConnectionRef() *VPNConnectionReference
}

// The jsii proxy for IVPNConnectionRef
type jsiiProxy_IVPNConnectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVPNConnectionRef) VpnConnectionRef() *VPNConnectionReference {
	var returns *VPNConnectionReference
	_jsii_.Get(
		j,
		"vpnConnectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPNConnectionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPNConnectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

