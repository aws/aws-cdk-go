package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRouteTablePropagation.
// Experimental.
type ITransitGatewayRouteTablePropagationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayRouteTablePropagation resource.
	// Experimental.
	TransitGatewayRouteTablePropagationRef() *TransitGatewayRouteTablePropagationReference
}

// The jsii proxy for ITransitGatewayRouteTablePropagationRef
type jsiiProxy_ITransitGatewayRouteTablePropagationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayRouteTablePropagationRef) TransitGatewayRouteTablePropagationRef() *TransitGatewayRouteTablePropagationReference {
	var returns *TransitGatewayRouteTablePropagationReference
	_jsii_.Get(
		j,
		"transitGatewayRouteTablePropagationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTablePropagationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTablePropagationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

