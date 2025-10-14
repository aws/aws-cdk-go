package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeviceDefinitionVersion.
// Experimental.
type IDeviceDefinitionVersionRef interface {
	constructs.IConstruct
	// A reference to a DeviceDefinitionVersion resource.
	// Experimental.
	DeviceDefinitionVersionRef() *DeviceDefinitionVersionReference
}

// The jsii proxy for IDeviceDefinitionVersionRef
type jsiiProxy_IDeviceDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
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

