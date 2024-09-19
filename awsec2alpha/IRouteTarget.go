package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface to define a routing target, such as an egress-only internet gateway or VPC endpoint.
// Experimental.
type IRouteTarget interface {
	constructs.IDependable
	// The ID of the route target.
	// Experimental.
	RouterTargetId() *string
	// The type of router used in the route.
	// Experimental.
	RouterType() awsec2.RouterType
}

// The jsii proxy for IRouteTarget
type jsiiProxy_IRouteTarget struct {
	internal.Type__constructsIDependable
}

func (j *jsiiProxy_IRouteTarget) RouterTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteTarget) RouterType() awsec2.RouterType {
	var returns awsec2.RouterType
	_jsii_.Get(
		j,
		"routerType",
		&returns,
	)
	return returns
}

