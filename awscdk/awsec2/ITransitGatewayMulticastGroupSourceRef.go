package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayMulticastGroupSource.
// Experimental.
type ITransitGatewayMulticastGroupSourceRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayMulticastGroupSource resource.
	// Experimental.
	TransitGatewayMulticastGroupSourceRef() *TransitGatewayMulticastGroupSourceReference
}

// The jsii proxy for ITransitGatewayMulticastGroupSourceRef
type jsiiProxy_ITransitGatewayMulticastGroupSourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayMulticastGroupSourceRef) TransitGatewayMulticastGroupSourceRef() *TransitGatewayMulticastGroupSourceReference {
	var returns *TransitGatewayMulticastGroupSourceReference
	_jsii_.Get(
		j,
		"transitGatewayMulticastGroupSourceRef",
		&returns,
	)
	return returns
}

