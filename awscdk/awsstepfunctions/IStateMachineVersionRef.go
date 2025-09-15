package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateMachineVersion.
// Experimental.
type IStateMachineVersionRef interface {
	constructs.IConstruct
	// A reference to a StateMachineVersion resource.
	// Experimental.
	StateMachineVersionRef() *StateMachineVersionReference
}

// The jsii proxy for IStateMachineVersionRef
type jsiiProxy_IStateMachineVersionRef struct {
	internal.Type__constructsIConstruct
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

