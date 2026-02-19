package interfacesawsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComputeEnvironment.
// Experimental.
type IComputeEnvironmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ComputeEnvironment resource.
	// Experimental.
	ComputeEnvironmentRef() *ComputeEnvironmentReference
}

// The jsii proxy for IComputeEnvironmentRef
type jsiiProxy_IComputeEnvironmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IComputeEnvironmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IComputeEnvironmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

