package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputeEnvironment.
// Experimental.
type IComputeEnvironmentRef interface {
	constructs.IConstruct
	// A reference to a ComputeEnvironment resource.
	// Experimental.
	ComputeEnvironmentRef() *ComputeEnvironmentReference
}

// The jsii proxy for IComputeEnvironmentRef
type jsiiProxy_IComputeEnvironmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IComputeEnvironmentRef) ComputeEnvironmentRef() *ComputeEnvironmentReference {
	var returns *ComputeEnvironmentReference
	_jsii_.Get(
		j,
		"computeEnvironmentRef",
		&returns,
	)
	return returns
}

