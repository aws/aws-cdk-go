package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRouteTable.
// Experimental.
type ITransitGatewayRouteTableRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayRouteTable resource.
	// Experimental.
	TransitGatewayRouteTableRef() *TransitGatewayRouteTableReference
}

// The jsii proxy for ITransitGatewayRouteTableRef
type jsiiProxy_ITransitGatewayRouteTableRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayRouteTableRef) TransitGatewayRouteTableRef() *TransitGatewayRouteTableReference {
	var returns *TransitGatewayRouteTableReference
	_jsii_.Get(
		j,
		"transitGatewayRouteTableRef",
		&returns,
	)
	return returns
}

