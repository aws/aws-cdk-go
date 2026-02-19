package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskTemplate.
// Experimental.
type ITaskTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TaskTemplate resource.
	// Experimental.
	TaskTemplateRef() *TaskTemplateReference
}

// The jsii proxy for ITaskTemplateRef
type jsiiProxy_ITaskTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITaskTemplateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITaskTemplateRef) TaskTemplateRef() *TaskTemplateReference {
	var returns *TaskTemplateReference
	_jsii_.Get(
		j,
		"taskTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

