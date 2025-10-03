package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualMFADevice.
// Experimental.
type IVirtualMFADeviceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVirtualMFADeviceRef
type jsiiProxy_IVirtualMFADeviceRef struct {
	internal.Type__constructsIConstruct
}

