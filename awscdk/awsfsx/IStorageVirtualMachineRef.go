package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageVirtualMachine.
// Experimental.
type IStorageVirtualMachineRef interface {
	constructs.IConstruct
	// A reference to a StorageVirtualMachine resource.
	// Experimental.
	StorageVirtualMachineRef() *StorageVirtualMachineReference
}

// The jsii proxy for IStorageVirtualMachineRef
type jsiiProxy_IStorageVirtualMachineRef struct {
	internal.Type__constructsIConstruct
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

