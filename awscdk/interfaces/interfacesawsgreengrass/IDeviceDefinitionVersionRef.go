package interfacesawsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceDefinitionVersion.
// Experimental.
type IDeviceDefinitionVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DeviceDefinitionVersion resource.
	// Experimental.
	DeviceDefinitionVersionRef() *DeviceDefinitionVersionReference
}

// The jsii proxy for IDeviceDefinitionVersionRef
type jsiiProxy_IDeviceDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDeviceDefinitionVersionRef) DeviceDefinitionVersionRef() *DeviceDefinitionVersionReference {
	var returns *DeviceDefinitionVersionReference
	_jsii_.Get(
		j,
		"deviceDefinitionVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeviceDefinitionVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeviceDefinitionVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

