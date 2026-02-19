package interfacesawsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskSet.
// Experimental.
type ITaskSetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TaskSet resource.
	// Experimental.
	TaskSetRef() *TaskSetReference
}

// The jsii proxy for ITaskSetRef
type jsiiProxy_ITaskSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITaskSetRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITaskSetRef) TaskSetRef() *TaskSetReference {
	var returns *TaskSetReference
	_jsii_.Get(
		j,
		"taskSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskSetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

