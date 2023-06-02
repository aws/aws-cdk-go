package awscdkbatchalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an UnmanagedComputeEnvironment.
//
// Batch will not provision instances on your behalf
// in this ComputeEvironment.
// Experimental.
type IUnmanagedComputeEnvironment interface {
	IComputeEnvironment
	// The vCPUs this Compute Environment provides. Used only by the scheduler to schedule jobs in `Queue`s that use `FairshareSchedulingPolicy`s.
	//
	// **If this parameter is not provided on a fairshare queue, no capacity is reserved**;
	// that is, the `FairshareSchedulingPolicy` is ignored.
	// Experimental.
	UnmanagedvCPUs() *float64
}

// The jsii proxy for IUnmanagedComputeEnvironment
type jsiiProxy_IUnmanagedComputeEnvironment struct {
	jsiiProxy_IComputeEnvironment
}

func (j *jsiiProxy_IUnmanagedComputeEnvironment) UnmanagedvCPUs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unmanagedvCPUs",
		&returns,
	)
	return returns
}

