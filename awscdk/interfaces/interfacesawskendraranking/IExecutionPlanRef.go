package interfacesawskendraranking

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskendraranking/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExecutionPlan.
// Experimental.
type IExecutionPlanRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ExecutionPlan resource.
	// Experimental.
	ExecutionPlanRef() *ExecutionPlanReference
}

// The jsii proxy for IExecutionPlanRef
type jsiiProxy_IExecutionPlanRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IExecutionPlanRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IExecutionPlanRef) ExecutionPlanRef() *ExecutionPlanReference {
	var returns *ExecutionPlanReference
	_jsii_.Get(
		j,
		"executionPlanRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExecutionPlanRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExecutionPlanRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

