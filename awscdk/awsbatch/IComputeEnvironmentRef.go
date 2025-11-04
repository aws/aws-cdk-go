package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputeEnvironment.
// Experimental.
type IComputeEnvironmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ComputeEnvironment resource.
	// Experimental.
	ComputeEnvironmentRef() *ComputeEnvironmentReference
}

// The jsii proxy for IComputeEnvironmentRef
type jsiiProxy_IComputeEnvironmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IComputeEnvironmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

