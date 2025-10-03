package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayRouteTableAssociation.
// Experimental.
type IGatewayRouteTableAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGatewayRouteTableAssociationRef
type jsiiProxy_IGatewayRouteTableAssociationRef struct {
	internal.Type__constructsIConstruct
}

