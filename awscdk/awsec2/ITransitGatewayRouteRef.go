package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayRoute.
// Experimental.
type ITransitGatewayRouteRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayRoute resource.
	// Experimental.
	TransitGatewayRouteRef() *TransitGatewayRouteReference
}

// The jsii proxy for ITransitGatewayRouteRef
type jsiiProxy_ITransitGatewayRouteRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayRouteRef) TransitGatewayRouteRef() *TransitGatewayRouteReference {
	var returns *TransitGatewayRouteReference
	_jsii_.Get(
		j,
		"transitGatewayRouteRef",
		&returns,
	)
	return returns
}

