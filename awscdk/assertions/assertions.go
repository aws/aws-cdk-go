package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Partial and special matching during template assertions.
// Experimental.
type Match interface {
}

// The jsii proxy struct for Match
type jsiiProxy_Match struct {
	_ byte // padding
}

// Experimental.
func NewMatch() Match {
	_init_.Initialize()

	j := jsiiProxy_Match{}

	_jsii_.Create(
		"monocdk.assertions.Match",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMatch_Override(m Match) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.assertions.Match",
		nil, // no parameters
		m,
	)
}

// Use this matcher in the place of a field's value, if the field must not be present.
// Experimental.
func Match_AbsentProperty() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"absentProperty",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Suite of assertions that can be run on a CDK stack.
//
// Typically used, as part of unit tests, to validate that the rendered
// CloudFormation template has expected resources and properties.
// Experimental.
type TemplateAssertions interface {
	HasResourceDefinition(type_ *string, props interface{})
	HasResourceProperties(type_ *string, props interface{})
	ResourceCountIs(type_ *string, count *float64)
	TemplateMatches(expected *map[string]interface{})
}

// The jsii proxy struct for TemplateAssertions
type jsiiProxy_TemplateAssertions struct {
	_ byte // padding
}

// Base your assertions on the CloudFormation template synthesized by a CDK `Stack`.
// Experimental.
func TemplateAssertions_FromStack(stack awscdk.Stack) TemplateAssertions {
	_init_.Initialize()

	var returns TemplateAssertions

	_jsii_.StaticInvoke(
		"monocdk.assertions.TemplateAssertions",
		"fromStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

// Base your assertions from an existing CloudFormation template formatted as a string.
// Experimental.
func TemplateAssertions_FromString(template *string) TemplateAssertions {
	_init_.Initialize()

	var returns TemplateAssertions

	_jsii_.StaticInvoke(
		"monocdk.assertions.TemplateAssertions",
		"fromString",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Base your assertions from an existing CloudFormation template formatted as a nested set of records.
// Experimental.
func TemplateAssertions_FromTemplate(template *map[string]interface{}) TemplateAssertions {
	_init_.Initialize()

	var returns TemplateAssertions

	_jsii_.StaticInvoke(
		"monocdk.assertions.TemplateAssertions",
		"fromTemplate",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Assert that a resource of the given type and given definition exists in the CloudFormation template.
// Experimental.
func (t *jsiiProxy_TemplateAssertions) HasResourceDefinition(type_ *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasResourceDefinition",
		[]interface{}{type_, props},
	)
}

// Assert that a resource of the given type and properties exists in the CloudFormation template.
// Experimental.
func (t *jsiiProxy_TemplateAssertions) HasResourceProperties(type_ *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasResourceProperties",
		[]interface{}{type_, props},
	)
}

// Assert that the given number of resources of the given type exist in the template.
// Experimental.
func (t *jsiiProxy_TemplateAssertions) ResourceCountIs(type_ *string, count *float64) {
	_jsii_.InvokeVoid(
		t,
		"resourceCountIs",
		[]interface{}{type_, count},
	)
}

// Assert that the CloudFormation template matches the given value.
// Experimental.
func (t *jsiiProxy_TemplateAssertions) TemplateMatches(expected *map[string]interface{}) {
	_jsii_.InvokeVoid(
		t,
		"templateMatches",
		[]interface{}{expected},
	)
}

