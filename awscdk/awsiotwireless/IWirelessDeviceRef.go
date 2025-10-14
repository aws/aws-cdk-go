package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WirelessDevice.
// Experimental.
type IWirelessDeviceRef interface {
	constructs.IConstruct
	// A reference to a WirelessDevice resource.
	// Experimental.
	WirelessDeviceRef() *WirelessDeviceReference
}

// The jsii proxy for IWirelessDeviceRef
type jsiiProxy_IWirelessDeviceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWirelessDeviceRef) WirelessDeviceRef() *WirelessDeviceReference {
	var returns *WirelessDeviceReference
	_jsii_.Get(
		j,
		"wirelessDeviceRef",
		&returns,
	)
	return returns
}

