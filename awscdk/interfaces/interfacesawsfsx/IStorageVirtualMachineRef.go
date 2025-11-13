package interfacesawsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageVirtualMachine.
// Experimental.
type IStorageVirtualMachineRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StorageVirtualMachine resource.
	// Experimental.
	StorageVirtualMachineRef() *StorageVirtualMachineReference
}

// The jsii proxy for IStorageVirtualMachineRef
type jsiiProxy_IStorageVirtualMachineRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStorageVirtualMachineRef) StorageVirtualMachineRef() *StorageVirtualMachineReference {
	var returns *StorageVirtualMachineReference
	_jsii_.Get(
		j,
		"storageVirtualMachineRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageVirtualMachineRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStorageVirtualMachineRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

