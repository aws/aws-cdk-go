package interfacesawsdirectconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdirectconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DirectConnectGatewayAssociation.
// Experimental.
type IDirectConnectGatewayAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DirectConnectGatewayAssociation resource.
	// Experimental.
	DirectConnectGatewayAssociationRef() *DirectConnectGatewayAssociationReference
}

// The jsii proxy for IDirectConnectGatewayAssociationRef
type jsiiProxy_IDirectConnectGatewayAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDirectConnectGatewayAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDirectConnectGatewayAssociationRef) DirectConnectGatewayAssociationRef() *DirectConnectGatewayAssociationReference {
	var returns *DirectConnectGatewayAssociationReference
	_jsii_.Get(
		j,
		"directConnectGatewayAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectConnectGatewayAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDirectConnectGatewayAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

