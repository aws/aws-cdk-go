package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
)

// Interface to define a route.
// Experimental.
type IRouteV2 interface {
	awscdk.IResource
	// The IPv4 or IPv6 CIDR block used for the destination match.
	//
	// Routing decisions are based on the most specific match.
	// TODO: Look for strong IP type implementation here.
	// Experimental.
	Destination() *string
	// The ID of the route table for the route.
	// Experimental.
	RouteTable() awsec2.IRouteTable
	// The gateway or endpoint targeted by the route.
	// Experimental.
	Target() RouteTargetType
}

// The jsii proxy for IRouteV2
type jsiiProxy_IRouteV2 struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRouteV2) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteV2) RouteTable() awsec2.IRouteTable {
	var returns awsec2.IRouteTable
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteV2) Target() RouteTargetType {
	var returns RouteTargetType
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

