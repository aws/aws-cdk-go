package interfacesawscontrolcatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscontrolcatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Objective.
// Experimental.
type IObjectiveRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Objective resource.
	// Experimental.
	ObjectiveRef() *ObjectiveReference
}

// The jsii proxy for IObjectiveRef
type jsiiProxy_IObjectiveRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IObjectiveRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IObjectiveRef) ObjectiveRef() *ObjectiveReference {
	var returns *ObjectiveReference
	_jsii_.Get(
		j,
		"objectiveRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IObjectiveRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IObjectiveRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

