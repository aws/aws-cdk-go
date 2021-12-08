package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Capture values while matching templates.
//
// Using an instance of this class within a Matcher will capture the matching value.
// The `as*()` APIs on the instance can be used to get the captured value.
//
// TODO: EXAMPLE
//
type Capture interface {
	Matcher
	Name() *string
	AsArray() *[]interface{}
	AsBoolean() *bool
	AsNumber() *float64
	AsObject() *map[string]interface{}
	AsString() *string
	Next() *bool
	Test(actual interface{}) MatchResult
}

// The jsii proxy struct for Capture
type jsiiProxy_Capture struct {
	jsiiProxy_Matcher
}

func (j *jsiiProxy_Capture) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Initialize a new capture.
func NewCapture(pattern interface{}) Capture {
	_init_.Initialize()

	j := jsiiProxy_Capture{}

	_jsii_.Create(
		"aws-cdk-lib.assertions.Capture",
		[]interface{}{pattern},
		&j,
	)

	return &j
}

// Initialize a new capture.
func NewCapture_Override(c Capture, pattern interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.assertions.Capture",
		[]interface{}{pattern},
		c,
	)
}

// Check whether the provided object is a subtype of the `IMatcher`.
func Capture_IsMatcher(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Capture",
		"isMatcher",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Retrieve the captured value as an array.
//
// An error is generated if no value is captured or if the value is not an array.
func (c *jsiiProxy_Capture) AsArray() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"asArray",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Retrieve the captured value as a boolean.
//
// An error is generated if no value is captured or if the value is not a boolean.
func (c *jsiiProxy_Capture) AsBoolean() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"asBoolean",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Retrieve the captured value as a number.
//
// An error is generated if no value is captured or if the value is not a number.
func (c *jsiiProxy_Capture) AsNumber() *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"asNumber",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Retrieve the captured value as a JSON object.
//
// An error is generated if no value is captured or if the value is not an object.
func (c *jsiiProxy_Capture) AsObject() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"asObject",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Retrieve the captured value as a string.
//
// An error is generated if no value is captured or if the value is not a string.
func (c *jsiiProxy_Capture) AsString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"asString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// When multiple results are captured, move the iterator to the next result.
//
// Returns: true if another capture is present, false otherwise
func (c *jsiiProxy_Capture) Next() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"next",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Test whether a target matches the provided pattern.
//
// Every Matcher must implement this method.
// This method will be invoked by the assertions framework. Do not call this method directly.
func (c *jsiiProxy_Capture) Test(actual interface{}) MatchResult {
	var returns MatchResult

	_jsii_.Invoke(
		c,
		"test",
		[]interface{}{actual},
		&returns,
	)

	return returns
}

// Partial and special matching during template assertions.
type Match interface {
}

// The jsii proxy struct for Match
type jsiiProxy_Match struct {
	_ byte // padding
}

func NewMatch_Override(m Match) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.assertions.Match",
		nil, // no parameters
		m,
	)
}

// Use this matcher in the place of a field's value, if the field must not be present.
func Match_Absent() Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"absent",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches any non-null value at the target.
func Match_AnyValue() Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"anyValue",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must match exactly and in order.
func Match_ArrayEquals(pattern *[]interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"arrayEquals",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must be in the same order as would be found.
func Match_ArrayWith(pattern *[]interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"arrayWith",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Deep exact matching of the specified pattern to the target.
func Match_Exact(pattern interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"exact",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches any target which does NOT follow the specified pattern.
func Match_Not(pattern interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"not",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must match exactly with the target.
func Match_ObjectEquals(pattern *map[string]interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"objectEquals",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must be present in the target but the target can be a superset.
func Match_ObjectLike(pattern *map[string]interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"objectLike",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches any string-encoded JSON and applies the specified pattern after parsing it.
func Match_SerializedJson(pattern interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Match",
		"serializedJson",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Information about a value captured during match.
//
// TODO: EXAMPLE
//
type MatchCapture struct {
	// The instance of Capture class to which this capture is associated with.
	Capture Capture `json:"capture"`
	// The value that was captured.
	Value interface{} `json:"value"`
}

// Match failure details.
//
// TODO: EXAMPLE
//
type MatchFailure struct {
	// The matcher that had the failure.
	Matcher Matcher `json:"matcher"`
	// Failure message.
	Message *string `json:"message"`
	// The relative path in the target where the failure occurred.
	//
	// If the failure occurred at root of the match tree, set the path to an empty list.
	// If it occurs in the 5th index of an array nested within the 'foo' key of an object,
	// set the path as `['/foo', '[5]']`.
	Path *[]*string `json:"path"`
}

// The result of `Match.test()`.
//
// TODO: EXAMPLE
//
type MatchResult interface {
	FailCount() *float64
	Target() interface{}
	Compose(id *string, inner MatchResult) MatchResult
	Finished() MatchResult
	HasFailed() *bool
	Push(matcher Matcher, path *[]*string, message *string) MatchResult
	RecordCapture(options *MatchCapture)
	RecordFailure(failure *MatchFailure) MatchResult
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


func NewMatchResult(target interface{}) MatchResult {
	_init_.Initialize()

	j := jsiiProxy_MatchResult{}

	_jsii_.Create(
		"aws-cdk-lib.assertions.MatchResult",
		[]interface{}{target},
		&j,
	)

	return &j
}

func NewMatchResult_Override(m MatchResult, target interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.assertions.MatchResult",
		[]interface{}{target},
		m,
	)
}

// Compose the results of a previous match as a subtree.
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

// Prepare the result to be analyzed.
//
// This API *must* be called prior to analyzing these results.
func (m *jsiiProxy_MatchResult) Finished() MatchResult {
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"finished",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Does the result contain any failures.
//
// If not, the result is a success
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

// DEPRECATED.
// Deprecated: use recordFailure()
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

// Record a capture against in this match result.
func (m *jsiiProxy_MatchResult) RecordCapture(options *MatchCapture) {
	_jsii_.InvokeVoid(
		m,
		"recordCapture",
		[]interface{}{options},
	)
}

// Record a new failure into this result at a specific path.
func (m *jsiiProxy_MatchResult) RecordFailure(failure *MatchFailure) MatchResult {
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"recordFailure",
		[]interface{}{failure},
		&returns,
	)

	return returns
}

// Get the list of failures as human readable strings.
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
//
// TODO: EXAMPLE
//
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


func NewMatcher_Override(m Matcher) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.assertions.Matcher",
		nil, // no parameters
		m,
	)
}

// Check whether the provided object is a subtype of the `IMatcher`.
func Matcher_IsMatcher(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Matcher",
		"isMatcher",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether a target matches the provided pattern.
//
// Every Matcher must implement this method.
// This method will be invoked by the assertions framework. Do not call this method directly.
//
// Returns: the list of match failures. An empty array denotes a successful match.
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
//
// TODO: EXAMPLE
//
type Template interface {
	FindMappings(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	FindOutputs(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	FindResources(type_ *string, props interface{}) *map[string]*map[string]interface{}
	HasMapping(logicalId *string, props interface{})
	HasOutput(logicalId *string, props interface{})
	HasResource(type_ *string, props interface{})
	HasResourceProperties(type_ *string, props interface{})
	ResourceCountIs(type_ *string, count *float64)
	TemplateMatches(expected interface{})
	ToJSON() *map[string]interface{}
}

// The jsii proxy struct for Template
type jsiiProxy_Template struct {
	_ byte // padding
}

// Base your assertions from an existing CloudFormation template formatted as an in-memory JSON object.
func Template_FromJSON(template *map[string]interface{}) Template {
	_init_.Initialize()

	var returns Template

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Template",
		"fromJSON",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Base your assertions on the CloudFormation template synthesized by a CDK `Stack`.
func Template_FromStack(stack awscdk.Stack) Template {
	_init_.Initialize()

	var returns Template

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Template",
		"fromStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

// Base your assertions from an existing CloudFormation template formatted as a JSON string.
func Template_FromString(template *string) Template {
	_init_.Initialize()

	var returns Template

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Template",
		"fromString",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Get the set of matching Mappings that match the given properties in the CloudFormation template.
func (t *jsiiProxy_Template) FindMappings(logicalId *string, props interface{}) *map[string]*map[string]interface{} {
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findMappings",
		[]interface{}{logicalId, props},
		&returns,
	)

	return returns
}

// Get the set of matching Outputs that match the given properties in the CloudFormation template.
func (t *jsiiProxy_Template) FindOutputs(logicalId *string, props interface{}) *map[string]*map[string]interface{} {
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findOutputs",
		[]interface{}{logicalId, props},
		&returns,
	)

	return returns
}

// Get the set of matching resources of a given type and properties in the CloudFormation template.
func (t *jsiiProxy_Template) FindResources(type_ *string, props interface{}) *map[string]*map[string]interface{} {
	var returns *map[string]*map[string]interface{}

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
func (t *jsiiProxy_Template) HasMapping(logicalId *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasMapping",
		[]interface{}{logicalId, props},
	)
}

// Assert that an Output with the given properties exists in the CloudFormation template.
//
// By default, performs partial matching on the resource, via the `Match.objectLike()`.
// To configure different behavour, use other matchers in the `Match` class.
func (t *jsiiProxy_Template) HasOutput(logicalId *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasOutput",
		[]interface{}{logicalId, props},
	)
}

// Assert that a resource of the given type and given definition exists in the CloudFormation template.
//
// By default, performs partial matching on the resource, via the `Match.objectLike()`.
// To configure different behavour, use other matchers in the `Match` class.
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
func (t *jsiiProxy_Template) HasResourceProperties(type_ *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasResourceProperties",
		[]interface{}{type_, props},
	)
}

// Assert that the given number of resources of the given type exist in the template.
func (t *jsiiProxy_Template) ResourceCountIs(type_ *string, count *float64) {
	_jsii_.InvokeVoid(
		t,
		"resourceCountIs",
		[]interface{}{type_, count},
	)
}

// Assert that the CloudFormation template matches the given value.
func (t *jsiiProxy_Template) TemplateMatches(expected interface{}) {
	_jsii_.InvokeVoid(
		t,
		"templateMatches",
		[]interface{}{expected},
	)
}

// The CloudFormation template deserialized into an object.
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

