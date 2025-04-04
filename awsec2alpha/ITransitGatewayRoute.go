package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
)

// Represents a Transit Gateway Route.
// Experimental.
type ITransitGatewayRoute interface {
	awscdk.IResource
	// The destination CIDR block for this route.
	//
	// Destination Cidr cannot overlap for static routes but is allowed for propagated routes.
	// When overlapping occurs, static routes take precedence over propagated routes.
	// Experimental.
	DestinationCidrBlock() *string
	// The transit gateway route table this route belongs to.
	// Experimental.
	RouteTable() ITransitGatewayRouteTable
}

// The jsii proxy for ITransitGatewayRoute
type jsiiProxy_ITransitGatewayRoute struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITransitGatewayRoute) DestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRoute) RouteTable() ITransitGatewayRouteTable {
	var returns ITransitGatewayRouteTable
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

