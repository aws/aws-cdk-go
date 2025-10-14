package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceFleet.
// Experimental.
type IDeviceFleetRef interface {
	constructs.IConstruct
	// A reference to a DeviceFleet resource.
	// Experimental.
	DeviceFleetRef() *DeviceFleetReference
}

// The jsii proxy for IDeviceFleetRef
type jsiiProxy_IDeviceFleetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDeviceFleetRef) DeviceFleetRef() *DeviceFleetReference {
	var returns *DeviceFleetReference
	_jsii_.Get(
		j,
		"deviceFleetRef",
		&returns,
	)
	return returns
}

