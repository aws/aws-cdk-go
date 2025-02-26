package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
)

// Represents a Transit Gateway Route Table Propagation.
// Experimental.
type ITransitGatewayRouteTablePropagation interface {
	awscdk.IResource
	// The ID of the transit gateway route table propagation.
	// Experimental.
	TransitGatewayRouteTablePropagationId() *string
}

// The jsii proxy for ITransitGatewayRouteTablePropagation
type jsiiProxy_ITransitGatewayRouteTablePropagation struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITransitGatewayRouteTablePropagation) TransitGatewayRouteTablePropagationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayRouteTablePropagationId",
		&returns,
	)
	return returns
}

