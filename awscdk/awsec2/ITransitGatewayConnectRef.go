package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayConnect.
// Experimental.
type ITransitGatewayConnectRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayConnect resource.
	// Experimental.
	TransitGatewayConnectRef() *TransitGatewayConnectReference
}

// The jsii proxy for ITransitGatewayConnectRef
type jsiiProxy_ITransitGatewayConnectRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayConnectRef) TransitGatewayConnectRef() *TransitGatewayConnectReference {
	var returns *TransitGatewayConnectReference
	_jsii_.Get(
		j,
		"transitGatewayConnectRef",
		&returns,
	)
	return returns
}

