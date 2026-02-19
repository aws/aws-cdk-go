package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayPeering.
// Experimental.
type ITransitGatewayPeeringRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayPeering resource.
	// Experimental.
	TransitGatewayPeeringRef() *TransitGatewayPeeringReference
}

// The jsii proxy for ITransitGatewayPeeringRef
type jsiiProxy_ITransitGatewayPeeringRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITransitGatewayPeeringRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITransitGatewayPeeringRef) TransitGatewayPeeringRef() *TransitGatewayPeeringReference {
	var returns *TransitGatewayPeeringReference
	_jsii_.Get(
		j,
		"transitGatewayPeeringRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPeeringRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPeeringRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

