package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayPeering.
// Experimental.
type ITransitGatewayPeeringRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayPeering resource.
	// Experimental.
	TransitGatewayPeeringRef() *TransitGatewayPeeringReference
}

// The jsii proxy for ITransitGatewayPeeringRef
type jsiiProxy_ITransitGatewayPeeringRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayPeeringRef) TransitGatewayPeeringRef() *TransitGatewayPeeringReference {
	var returns *TransitGatewayPeeringReference
	_jsii_.Get(
		j,
		"transitGatewayPeeringRef",
		&returns,
	)
	return returns
}

