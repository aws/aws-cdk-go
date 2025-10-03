package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WirelessDevice.
// Experimental.
type IWirelessDeviceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWirelessDeviceRef
type jsiiProxy_IWirelessDeviceRef struct {
	internal.Type__constructsIConstruct
}

