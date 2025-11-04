package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageVirtualMachine.
// Experimental.
type IStorageVirtualMachineRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StorageVirtualMachine resource.
	// Experimental.
	StorageVirtualMachineRef() *StorageVirtualMachineReference
}

// The jsii proxy for IStorageVirtualMachineRef
type jsiiProxy_IStorageVirtualMachineRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IStorageVirtualMachineRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

