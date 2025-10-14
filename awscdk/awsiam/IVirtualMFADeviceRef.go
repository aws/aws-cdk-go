package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualMFADevice.
// Experimental.
type IVirtualMFADeviceRef interface {
	constructs.IConstruct
	// A reference to a VirtualMFADevice resource.
	// Experimental.
	VirtualMfaDeviceRef() *VirtualMFADeviceReference
}

// The jsii proxy for IVirtualMFADeviceRef
type jsiiProxy_IVirtualMFADeviceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVirtualMFADeviceRef) VirtualMfaDeviceRef() *VirtualMFADeviceReference {
	var returns *VirtualMFADeviceReference
	_jsii_.Get(
		j,
		"virtualMfaDeviceRef",
		&returns,
	)
	return returns
}

