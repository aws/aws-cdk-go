package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
)

// Interface for which all GatewayRoute based classes MUST implement.
type IGatewayRoute interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) for the GatewayRoute.
	GatewayRouteArn() *string
	// The name of the GatewayRoute.
	GatewayRouteName() *string
	// The VirtualGateway the GatewayRoute belongs to.
	VirtualGateway() IVirtualGateway
}

// The jsii proxy for IGatewayRoute
type jsiiProxy_IGatewayRoute struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IGatewayRoute) GatewayRouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) GatewayRouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayRouteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRoute) VirtualGateway() IVirtualGateway {
	var returns IVirtualGateway
	_jsii_.Get(
		j,
		"virtualGateway",
		&returns,
	)
	return returns
}

