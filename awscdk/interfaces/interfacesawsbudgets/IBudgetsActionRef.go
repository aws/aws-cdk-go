package interfacesawsbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbudgets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BudgetsAction.
// Experimental.
type IBudgetsActionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BudgetsAction resource.
	// Experimental.
	BudgetsActionRef() *BudgetsActionReference
}

// The jsii proxy for IBudgetsActionRef
type jsiiProxy_IBudgetsActionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IBudgetsActionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBudgetsActionRef) BudgetsActionRef() *BudgetsActionReference {
	var returns *BudgetsActionReference
	_jsii_.Get(
		j,
		"budgetsActionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBudgetsActionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBudgetsActionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

