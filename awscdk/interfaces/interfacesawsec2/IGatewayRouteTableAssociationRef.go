package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayRouteTableAssociation.
// Experimental.
type IGatewayRouteTableAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GatewayRouteTableAssociation resource.
	// Experimental.
	GatewayRouteTableAssociationRef() *GatewayRouteTableAssociationReference
}

// The jsii proxy for IGatewayRouteTableAssociationRef
type jsiiProxy_IGatewayRouteTableAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IGatewayRouteTableAssociationRef) GatewayRouteTableAssociationRef() *GatewayRouteTableAssociationReference {
	var returns *GatewayRouteTableAssociationReference
	_jsii_.Get(
		j,
		"gatewayRouteTableAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRouteTableAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRouteTableAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

