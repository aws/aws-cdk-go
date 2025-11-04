package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRouteTableAssociation.
// Experimental.
type ITransitGatewayRouteTableAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TransitGatewayRouteTableAssociation resource.
	// Experimental.
	TransitGatewayRouteTableAssociationRef() *TransitGatewayRouteTableAssociationReference
}

// The jsii proxy for ITransitGatewayRouteTableAssociationRef
type jsiiProxy_ITransitGatewayRouteTableAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayRouteTableAssociationRef) TransitGatewayRouteTableAssociationRef() *TransitGatewayRouteTableAssociationReference {
	var returns *TransitGatewayRouteTableAssociationReference
	_jsii_.Get(
		j,
		"transitGatewayRouteTableAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTableAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTableAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

