package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StorageVirtualMachine.
// Experimental.
type IStorageVirtualMachineRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStorageVirtualMachineRef
type jsiiProxy_IStorageVirtualMachineRef struct {
	internal.Type__constructsIConstruct
}

