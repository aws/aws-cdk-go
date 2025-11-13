package interfacesawsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateMachineVersion.
// Experimental.
type IStateMachineVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StateMachineVersion resource.
	// Experimental.
	StateMachineVersionRef() *StateMachineVersionReference
}

// The jsii proxy for IStateMachineVersionRef
type jsiiProxy_IStateMachineVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStateMachineVersionRef) StateMachineVersionRef() *StateMachineVersionReference {
	var returns *StateMachineVersionReference
	_jsii_.Get(
		j,
		"stateMachineVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateMachineVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateMachineVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

