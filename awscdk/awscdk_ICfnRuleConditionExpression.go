// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface to specify certain functions as Service Catalog rule-specifc.
//
// These functions can only be used in ``Rules`` section of template.
// Experimental.
type ICfnRuleConditionExpression interface {
	ICfnConditionExpression
	// This field is only needed to defeat TypeScript's structural typing.
	//
	// It is never used.
	// Experimental.
	Disambiguator() *bool
}

// The jsii proxy for ICfnRuleConditionExpression
type jsiiProxy_ICfnRuleConditionExpression struct {
	jsiiProxy_ICfnConditionExpression
}

func (j *jsiiProxy_ICfnRuleConditionExpression) Disambiguator() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disambiguator",
		&returns,
	)
	return returns
}

