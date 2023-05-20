package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Interface to represent a Matchmaking RuleSet content.
// Experimental.
type IRuleSetContent interface {
	// Called when the matchmaking ruleSet is initialized to allow this object to bind to the stack and add resources.
	// Experimental.
	Bind(_scope constructs.Construct) *RuleSetBodyConfig
	// RuleSet body content.
	// Experimental.
	Content() IRuleSetBody
}

// The jsii proxy for IRuleSetContent
type jsiiProxy_IRuleSetContent struct {
	_ byte // padding
}

func (i *jsiiProxy_IRuleSetContent) Bind(_scope constructs.Construct) *RuleSetBodyConfig {
	if err := i.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *RuleSetBodyConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRuleSetContent) Content() IRuleSetBody {
	var returns IRuleSetBody
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

