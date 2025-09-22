package awsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DevicePool.
// Experimental.
type IDevicePoolRef interface {
	constructs.IConstruct
	// A reference to a DevicePool resource.
	// Experimental.
	DevicePoolRef() *DevicePoolReference
}

// The jsii proxy for IDevicePoolRef
type jsiiProxy_IDevicePoolRef struct {
	internal.Type__constructsIConstruct
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

