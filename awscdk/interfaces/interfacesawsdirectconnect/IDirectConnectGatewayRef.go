package interfacesawsdirectconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdirectconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectConnectGateway.
// Experimental.
type IDirectConnectGatewayRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DirectConnectGateway resource.
	// Experimental.
	DirectConnectGatewayRef() *DirectConnectGatewayReference
}

// The jsii proxy for IDirectConnectGatewayRef
type jsiiProxy_IDirectConnectGatewayRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDirectConnectGatewayRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDirectConnectGatewayRef) DirectConnectGatewayRef() *DirectConnectGatewayReference {
	var returns *DirectConnectGatewayReference
	_jsii_.Get(
		j,
		"directConnectGatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectConnectGatewayRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectConnectGatewayRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

