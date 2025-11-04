package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceDefinition.
// Experimental.
type IDeviceDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DeviceDefinition resource.
	// Experimental.
	DeviceDefinitionRef() *DeviceDefinitionReference
}

// The jsii proxy for IDeviceDefinitionRef
type jsiiProxy_IDeviceDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDeviceDefinitionRef) DeviceDefinitionRef() *DeviceDefinitionReference {
	var returns *DeviceDefinitionReference
	_jsii_.Get(
		j,
		"deviceDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeviceDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeviceDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

