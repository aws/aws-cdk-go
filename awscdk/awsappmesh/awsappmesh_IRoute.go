package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappmesh/internal"
)

// Interface for which all Route based classes MUST implement.
// Experimental.
type IRoute interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) for the route.
	// Experimental.
	RouteArn() *string
	// The name of the route.
	// Experimental.
	RouteName() *string
	// The VirtualRouter the Route belongs to.
	// Experimental.
	VirtualRouter() IVirtualRouter
}

// The jsii proxy for IRoute
type jsiiProxy_IRoute struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRoute) RouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoute) VirtualRouter() IVirtualRouter {
	var returns IVirtualRouter
	_jsii_.Get(
		j,
		"virtualRouter",
		&returns,
	)
	return returns
}

