package awssam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateMachine.
// Experimental.
type IStateMachineRef interface {
	constructs.IConstruct
	// A reference to a StateMachine resource.
	// Experimental.
	StateMachineRef() *StateMachineReference
}

// The jsii proxy for IStateMachineRef
type jsiiProxy_IStateMachineRef struct {
	internal.Type__constructsIConstruct
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

