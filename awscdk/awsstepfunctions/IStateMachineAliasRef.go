package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateMachineAlias.
// Experimental.
type IStateMachineAliasRef interface {
	constructs.IConstruct
	// A reference to a StateMachineAlias resource.
	// Experimental.
	StateMachineAliasRef() *StateMachineAliasReference
}

// The jsii proxy for IStateMachineAliasRef
type jsiiProxy_IStateMachineAliasRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStateMachineAliasRef) StateMachineAliasRef() *StateMachineAliasReference {
	var returns *StateMachineAliasReference
	_jsii_.Get(
		j,
		"stateMachineAliasRef",
		&returns,
	)
	return returns
}

