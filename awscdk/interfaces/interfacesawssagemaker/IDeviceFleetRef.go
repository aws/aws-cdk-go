package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceFleet.
// Experimental.
type IDeviceFleetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DeviceFleet resource.
	// Experimental.
	DeviceFleetRef() *DeviceFleetReference
}

// The jsii proxy for IDeviceFleetRef
type jsiiProxy_IDeviceFleetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IDeviceFleetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeviceFleetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

