package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappmesh/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Interface which all Virtual Gateway based classes must implement.
// Experimental.
type IVirtualGateway interface {
	awscdk.IResource
	// Utility method to add a new GatewayRoute to the VirtualGateway.
	// Experimental.
	AddGatewayRoute(id *string, route *GatewayRouteBaseProps) GatewayRoute
	// Grants the given entity `appmesh:StreamAggregatedResources`.
	// Experimental.
	GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant
	// The Mesh which the VirtualGateway belongs to.
	// Experimental.
	Mesh() IMesh
	// The Amazon Resource Name (ARN) for the VirtualGateway.
	// Experimental.
	VirtualGatewayArn() *string
	// Name of the VirtualGateway.
	// Experimental.
	VirtualGatewayName() *string
}

// The jsii proxy for IVirtualGateway
type jsiiProxy_IVirtualGateway struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVirtualGateway) AddGatewayRoute(id *string, route *GatewayRouteBaseProps) GatewayRoute {
	if err := i.validateAddGatewayRouteParameters(id, route); err != nil {
		panic(err)
	}
	var returns GatewayRoute

	_jsii_.Invoke(
		i,
		"addGatewayRoute",
		[]interface{}{id, route},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVirtualGateway) GrantStreamAggregatedResources(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantStreamAggregatedResourcesParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStreamAggregatedResources",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVirtualGateway) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) VirtualGatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualGateway) VirtualGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualGatewayName",
		&returns,
	)
	return returns
}

