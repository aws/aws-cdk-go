package interfacesawsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DevicePool.
// Experimental.
type IDevicePoolRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DevicePool resource.
	// Experimental.
	DevicePoolRef() *DevicePoolReference
}

// The jsii proxy for IDevicePoolRef
type jsiiProxy_IDevicePoolRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDevicePoolRef) DevicePoolRef() *DevicePoolReference {
	var returns *DevicePoolReference
	_jsii_.Get(
		j,
		"devicePoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDevicePoolRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDevicePoolRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

