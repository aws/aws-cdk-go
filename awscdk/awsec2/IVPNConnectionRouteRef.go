package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPNConnectionRoute.
// Experimental.
type IVPNConnectionRouteRef interface {
	constructs.IConstruct
	// A reference to a VPNConnectionRoute resource.
	// Experimental.
	VpnConnectionRouteRef() *VPNConnectionRouteReference
}

// The jsii proxy for IVPNConnectionRouteRef
type jsiiProxy_IVPNConnectionRouteRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPNConnectionRouteRef) VpnConnectionRouteRef() *VPNConnectionRouteReference {
	var returns *VPNConnectionRouteReference
	_jsii_.Get(
		j,
		"vpnConnectionRouteRef",
		&returns,
	)
	return returns
}

