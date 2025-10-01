package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRouteTableAssociation.
// Experimental.
type ITransitGatewayRouteTableAssociationRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayRouteTableAssociation resource.
	// Experimental.
	TransitGatewayRouteTableAssociationRef() *TransitGatewayRouteTableAssociationReference
}

// The jsii proxy for ITransitGatewayRouteTableAssociationRef
type jsiiProxy_ITransitGatewayRouteTableAssociationRef struct {
	internal.Type__constructsIConstruct
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

