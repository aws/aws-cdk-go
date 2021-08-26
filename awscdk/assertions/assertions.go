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

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must match exactly and in order.
// Experimental.
func Match_ArrayEquals(pattern *[]interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"arrayEquals",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must be in the same order as would be found.
// Experimental.
func Match_ArrayWith(pattern *[]interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"arrayWith",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Deep exact matching of the specified pattern to the target.
// Experimental.
func Match_Exact(pattern interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"exact",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must match exactly with the target.
// Experimental.
func Match_ObjectEquals(pattern *map[string]interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"objectEquals",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must be present in the target but the target can be a superset.
// Experimental.
func Match_ObjectLike(pattern *map[string]interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"objectLike",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// The result of `Match.test()`.
// Experimental.
type MatchResult interface {
	FailCount() *float64
	Target() interface{}
	Compose(id *string, inner MatchResult) MatchResult
	HasFailed() *bool
	Push(matcher Matcher, path *[]*string, message *string) MatchResult
	ToHumanStrings() *[]*string
}

// The jsii proxy struct for MatchResult
type jsiiProxy_MatchResult struct {
	_ byte // padding
}

func (j *jsiiProxy_MatchResult) FailCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchResult) Target() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}


// Experimental.
func NewMatchResult(target interface{}) MatchResult {
	_init_.Initialize()

	j := jsiiProxy_MatchResult{}

	_jsii_.Create(
		"monocdk.assertions.MatchResult",
		[]interface{}{target},
		&j,
	)

	return &j
}

// Experimental.
func NewMatchResult_Override(m MatchResult, target interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.assertions.MatchResult",
		[]interface{}{target},
		m,
	)
}

// Compose the results of a previous match as a subtree.
// Experimental.
func (m *jsiiProxy_MatchResult) Compose(id *string, inner MatchResult) MatchResult {
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"compose",
		[]interface{}{id, inner},
		&returns,
	)

	return returns
}

// Does the result contain any failures.
//
// If not, the result is a success
// Experimental.
func (m *jsiiProxy_MatchResult) HasFailed() *bool {
	var returns *bool

	_jsii_.Invoke(
		m,
		"hasFailed",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Push a new failure into this result at a specific path.
//
// If the failure occurred at root of the match tree, set the path to an empty list.
// If it occurs in the 5th index of an array nested within the 'foo' key of an object,
// set the path as `['/foo', '[5]']`.
// Experimental.
func (m *jsiiProxy_MatchResult) Push(matcher Matcher, path *[]*string, message *string) MatchResult {
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"push",
		[]interface{}{matcher, path, message},
		&returns,
	)

	return returns
}

// Get the list of failures as human readable strings.
// Experimental.
func (m *jsiiProxy_MatchResult) ToHumanStrings() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"toHumanStrings",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a matcher that can perform special data matching capabilities between a given pattern and a target.
// Experimental.
type Matcher interface {
	Name() *string
	Test(actual interface{}) MatchResult
}

// The jsii proxy struct for Matcher
type jsiiProxy_Matcher struct {
	_ byte // padding
}

func (j *jsiiProxy_Matcher) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewMatcher_Override(m Matcher) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.assertions.Matcher",
		nil, // no parameters
		m,
	)
}

// Check whether the provided object is a subtype of the `IMatcher`.
// Experimental.
func Matcher_IsMatcher(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.assertions.Matcher",
		"isMatcher",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether a target matches the provided pattern.
//
// Returns: the list of match failures. An empty array denotes a successful match.
// Experimental.
func (m *jsiiProxy_Matcher) Test(actual interface{}) MatchResult {
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"test",
		[]interface{}{actual},
		&returns,
	)

	return returns
}

// Suite of assertions that can be run on a CDK stack.
//
// Typically used, as part of unit tests, to validate that the rendered
// CloudFormation template has expected resources and properties.
// Experimental.
type Template interface {
	FindMappings(props interface{}) *[]*map[string]interface{}
	FindOutputs(props interface{}) *[]*map[string]interface{}
	FindResources(type_ *string, props interface{}) *[]*map[string]interface{}
	HasMapping(props interface{})
	HasOutput(props interface{})
	HasResource(type_ *string, props interface{})
	HasResourceProperties(type_ *string, props interface{})
	ResourceCountIs(type_ *string, count *float64)
	TemplateMatches(expected *map[string]interface{})
	ToJSON() *map[string]interface{}
}

// The jsii proxy struct for Template
type jsiiProxy_Template struct {
	_ byte // padding
}

// Base your assertions from an existing CloudFormation template formatted as an in-memory JSON object.
// Experimental.
func Template_FromJSON(template *map[string]interface{}) Template {
	_init_.Initialize()

	var returns Template

	_jsii_.StaticInvoke(
		"monocdk.assertions.Template",
		"fromJSON",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Base your assertions on the CloudFormation template synthesized by a CDK `Stack`.
// Experimental.
func Template_FromStack(stack awscdk.Stack) Template {
	_init_.Initialize()

	var returns Template

	_jsii_.StaticInvoke(
		"monocdk.assertions.Template",
		"fromStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

// Base your assertions from an existing CloudFormation template formatted as a JSON string.
// Experimental.
func Template_FromString(template *string) Template {
	_init_.Initialize()

	var returns Template

	_jsii_.StaticInvoke(
		"monocdk.assertions.Template",
		"fromString",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Get the set of matching Mappings that match the given properties in the CloudFormation template.
// Experimental.
func (t *jsiiProxy_Template) FindMappings(props interface{}) *[]*map[string]interface{} {
	var returns *[]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findMappings",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Get the set of matching Outputs that match the given properties in the CloudFormation template.
// Experimental.
func (t *jsiiProxy_Template) FindOutputs(props interface{}) *[]*map[string]interface{} {
	var returns *[]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findOutputs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Get the set of matching resources of a given type and properties in the CloudFormation template.
// Experimental.
func (t *jsiiProxy_Template) FindResources(type_ *string, props interface{}) *[]*map[string]interface{} {
	var returns *[]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findResources",
		[]interface{}{type_, props},
		&returns,
	)

	return returns
}

// Assert that a Mapping with the given properties exists in the CloudFormation template.
//
// By default, performs partial matching on the resource, via the `Match.objectLike()`.
// To configure different behavour, use other matchers in the `Match` class.
// Experimental.
func (t *jsiiProxy_Template) HasMapping(props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasMapping",
		[]interface{}{props},
	)
}

// Assert that an Output with the given properties exists in the CloudFormation template.
//
// By default, performs partial matching on the resource, via the `Match.objectLike()`.
// To configure different behavour, use other matchers in the `Match` class.
// Experimental.
func (t *jsiiProxy_Template) HasOutput(props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasOutput",
		[]interface{}{props},
	)
}

// Assert that a resource of the given type and given definition exists in the CloudFormation template.
//
// By default, performs partial matching on the resource, via the `Match.objectLike()`.
// To configure different behavour, use other matchers in the `Match` class.
// Experimental.
func (t *jsiiProxy_Template) HasResource(type_ *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasResource",
		[]interface{}{type_, props},
	)
}

// Assert that a resource of the given type and properties exists in the CloudFormation template.
//
// By default, performs partial matching on the `Properties` key of the resource, via the
// `Match.objectLike()`. To configure different behavour, use other matchers in the `Match` class.
// Experimental.
func (t *jsiiProxy_Template) HasResourceProperties(type_ *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasResourceProperties",
		[]interface{}{type_, props},
	)
}

// Assert that the given number of resources of the given type exist in the template.
// Experimental.
func (t *jsiiProxy_Template) ResourceCountIs(type_ *string, count *float64) {
	_jsii_.InvokeVoid(
		t,
		"resourceCountIs",
		[]interface{}{type_, count},
	)
}

// Assert that the CloudFormation template matches the given value.
// Experimental.
func (t *jsiiProxy_Template) TemplateMatches(expected *map[string]interface{}) {
	_jsii_.InvokeVoid(
		t,
		"templateMatches",
		[]interface{}{expected},
	)
}

// The CloudFormation template deserialized into an object.
// Experimental.
func (t *jsiiProxy_Template) ToJSON() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

