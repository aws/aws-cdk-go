package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Device.
// Experimental.
type IDeviceRef interface {
	constructs.IConstruct
	// A reference to a Device resource.
	// Experimental.
	DeviceRef() *DeviceReference
}

// The jsii proxy for IDeviceRef
type jsiiProxy_IDeviceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDeviceRef) DeviceRef() *DeviceReference {
	var returns *DeviceReference
	_jsii_.Get(
		j,
		"deviceRef",
		&returns,
	)
	return returns
}

