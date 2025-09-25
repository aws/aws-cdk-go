package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientVpnRoute.
// Experimental.
type IClientVpnRouteRef interface {
	constructs.IConstruct
	// A reference to a ClientVpnRoute resource.
	// Experimental.
	ClientVpnRouteRef() *ClientVpnRouteReference
}

// The jsii proxy for IClientVpnRouteRef
type jsiiProxy_IClientVpnRouteRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IClientVpnRouteRef) ClientVpnRouteRef() *ClientVpnRouteReference {
	var returns *ClientVpnRouteReference
	_jsii_.Get(
		j,
		"clientVpnRouteRef",
		&returns,
	)
	return returns
}

