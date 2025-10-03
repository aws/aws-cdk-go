package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WirelessGateway.
// Experimental.
type IWirelessGatewayRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWirelessGatewayRef
type jsiiProxy_IWirelessGatewayRef struct {
	internal.Type__constructsIConstruct
}

