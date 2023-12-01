package awscdkapigatewayv2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/internal"
)

// Represents a route.
// Deprecated.
type IRoute interface {
	awscdk.IResource
	// Id of the Route.
	// Deprecated.
	RouteId() *string
}

// The jsii proxy for IRoute
type jsiiProxy_IRoute struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRoute) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
		&returns,
	)
	return returns
}

