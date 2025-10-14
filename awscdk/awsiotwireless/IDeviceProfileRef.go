package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceProfile.
// Experimental.
type IDeviceProfileRef interface {
	constructs.IConstruct
	// A reference to a DeviceProfile resource.
	// Experimental.
	DeviceProfileRef() *DeviceProfileReference
}

// The jsii proxy for IDeviceProfileRef
type jsiiProxy_IDeviceProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDeviceProfileRef) DeviceProfileRef() *DeviceProfileReference {
	var returns *DeviceProfileReference
	_jsii_.Get(
		j,
		"deviceProfileRef",
		&returns,
	)
	return returns
}

