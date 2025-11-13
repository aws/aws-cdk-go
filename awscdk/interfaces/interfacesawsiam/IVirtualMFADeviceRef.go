package interfacesawsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualMFADevice.
// Experimental.
type IVirtualMFADeviceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a VirtualMFADevice resource.
	// Experimental.
	VirtualMfaDeviceRef() *VirtualMFADeviceReference
}

// The jsii proxy for IVirtualMFADeviceRef
type jsiiProxy_IVirtualMFADeviceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IVirtualMFADeviceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVirtualMFADeviceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

