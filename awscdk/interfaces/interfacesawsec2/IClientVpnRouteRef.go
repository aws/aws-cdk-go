package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientVpnRoute.
// Experimental.
type IClientVpnRouteRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ClientVpnRoute resource.
	// Experimental.
	ClientVpnRouteRef() *ClientVpnRouteReference
}

// The jsii proxy for IClientVpnRouteRef
type jsiiProxy_IClientVpnRouteRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IClientVpnRouteRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IClientVpnRouteRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnRouteRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

