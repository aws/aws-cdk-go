package triggers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/triggers/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for triggers.
type ITrigger interface {
	constructs.IConstruct
	// Adds trigger dependencies.
	//
	// Execute this trigger only after these construct
	// scopes have been provisioned.
	ExecuteAfter(scopes ...constructs.Construct)
	// Adds this trigger as a dependency on other constructs.
	//
	// This means that this
	// trigger will get executed *before* the given construct(s).
	ExecuteBefore(scopes ...constructs.Construct)
}

// The jsii proxy for ITrigger
type jsiiProxy_ITrigger struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_ITrigger) ExecuteAfter(scopes ...constructs.Construct) {
	args := []interface{}{}
	for _, a := range scopes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"executeAfter",
		args,
	)
}

func (i *jsiiProxy_ITrigger) ExecuteBefore(scopes ...constructs.Construct) {
	args := []interface{}{}
	for _, a := range scopes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"executeBefore",
		args,
	)
}

