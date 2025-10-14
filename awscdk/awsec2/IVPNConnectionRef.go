package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPNConnection.
// Experimental.
type IVPNConnectionRef interface {
	constructs.IConstruct
	// A reference to a VPNConnection resource.
	// Experimental.
	VpnConnectionRef() *VPNConnectionReference
}

// The jsii proxy for IVPNConnectionRef
type jsiiProxy_IVPNConnectionRef struct {
	internal.Type__constructsIConstruct
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

