package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateMachine.
// Experimental.
type IStateMachineRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StateMachine resource.
	// Experimental.
	StateMachineRef() *StateMachineReference
}

// The jsii proxy for IStateMachineRef
type jsiiProxy_IStateMachineRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStateMachineRef) StateMachineRef() *StateMachineReference {
	var returns *StateMachineReference
	_jsii_.Get(
		j,
		"stateMachineRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateMachineRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateMachineRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

