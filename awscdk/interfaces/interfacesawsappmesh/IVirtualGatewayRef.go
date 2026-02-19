package interfacesawsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualGateway.
// Experimental.
type IVirtualGatewayRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VirtualGateway resource.
	// Experimental.
	VirtualGatewayRef() *VirtualGatewayReference
}

// The jsii proxy for IVirtualGatewayRef
type jsiiProxy_IVirtualGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IVirtualGatewayRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVirtualGatewayRef) VirtualGatewayRef() *VirtualGatewayReference {
	var returns *VirtualGatewayReference
	_jsii_.Get(
		j,
		"virtualGatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGatewayRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

