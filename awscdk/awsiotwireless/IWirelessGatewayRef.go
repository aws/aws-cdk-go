package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WirelessGateway.
// Experimental.
type IWirelessGatewayRef interface {
	constructs.IConstruct
	// A reference to a WirelessGateway resource.
	// Experimental.
	WirelessGatewayRef() *WirelessGatewayReference
}

// The jsii proxy for IWirelessGatewayRef
type jsiiProxy_IWirelessGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWirelessGatewayRef) WirelessGatewayRef() *WirelessGatewayReference {
	var returns *WirelessGatewayReference
	_jsii_.Get(
		j,
		"wirelessGatewayRef",
		&returns,
	)
	return returns
}

