package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DevicePool.
// Experimental.
type IDevicePoolRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDevicePoolRef
type jsiiProxy_IDevicePoolRef struct {
	internal.Type__constructsIConstruct
}

