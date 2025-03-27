package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An abstract target for EventRules.
type IRuleTarget interface {
	// Returns the rule target specification.
	//
	// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
	Bind(rule IRule, id *string) *RuleTargetConfig
}

// The jsii proxy for IRuleTarget
type jsiiProxy_IRuleTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IRuleTarget) Bind(rule IRule, id *string) *RuleTargetConfig {
	if err := i.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *RuleTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

