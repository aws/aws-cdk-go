package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayPeering.
// Experimental.
type ITransitGatewayPeeringRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TransitGatewayPeering resource.
	// Experimental.
	TransitGatewayPeeringRef() *TransitGatewayPeeringReference
}

// The jsii proxy for ITransitGatewayPeeringRef
type jsiiProxy_ITransitGatewayPeeringRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITransitGatewayPeeringRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

