package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
)

// Suite of assertions that can be run on a CDK Stack.
//
// Focused on asserting annotations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//   annotations := awscdk.Assertions.annotations.fromStack(stack)
//
// Experimental.
type Annotations interface {
	// Get the set of matching errors of a given construct path and message.
	// Experimental.
	FindError(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage
	// Get the set of matching infos of a given construct path and message.
	// Experimental.
	FindInfo(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage
	// Get the set of matching warning of a given construct path and message.
	// Experimental.
	FindWarning(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage
	// Assert that an error with the given message exists in the synthesized CDK `Stack`.
	// Experimental.
	HasError(constructPath *string, message interface{})
	// Assert that an info with the given message exists in the synthesized CDK `Stack`.
	// Experimental.
	HasInfo(constructPath *string, message interface{})
	// Assert that an error with the given message does not exist in the synthesized CDK `Stack`.
	// Experimental.
	HasNoError(constructPath *string, message interface{})
	// Assert that an info with the given message does not exist in the synthesized CDK `Stack`.
	// Experimental.
	HasNoInfo(constructPath *string, message interface{})
	// Assert that an warning with the given message does not exist in the synthesized CDK `Stack`.
	// Experimental.
	HasNoWarning(constructPath *string, message interface{})
	// Assert that an warning with the given message exists in the synthesized CDK `Stack`.
	// Experimental.
	HasWarning(constructPath *string, message interface{})
}

// The jsii proxy struct for Annotations
type jsiiProxy_Annotations struct {
	_ byte // padding
}

// Base your assertions on the messages returned by a synthesized CDK `Stack`.
// Experimental.
func Annotations_FromStack(stack awscdk.Stack) Annotations {
	_init_.Initialize()

	var returns Annotations

	_jsii_.StaticInvoke(
		"monocdk.assertions.Annotations",
		"fromStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) FindError(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage {
	var returns *[]*cxapi.SynthesisMessage

	_jsii_.Invoke(
		a,
		"findError",
		[]interface{}{constructPath, message},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) FindInfo(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage {
	var returns *[]*cxapi.SynthesisMessage

	_jsii_.Invoke(
		a,
		"findInfo",
		[]interface{}{constructPath, message},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) FindWarning(constructPath *string, message interface{}) *[]*cxapi.SynthesisMessage {
	var returns *[]*cxapi.SynthesisMessage

	_jsii_.Invoke(
		a,
		"findWarning",
		[]interface{}{constructPath, message},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) HasError(constructPath *string, message interface{}) {
	_jsii_.InvokeVoid(
		a,
		"hasError",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasInfo(constructPath *string, message interface{}) {
	_jsii_.InvokeVoid(
		a,
		"hasInfo",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasNoError(constructPath *string, message interface{}) {
	_jsii_.InvokeVoid(
		a,
		"hasNoError",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasNoInfo(constructPath *string, message interface{}) {
	_jsii_.InvokeVoid(
		a,
		"hasNoInfo",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasNoWarning(constructPath *string, message interface{}) {
	_jsii_.InvokeVoid(
		a,
		"hasNoWarning",
		[]interface{}{constructPath, message},
	)
}

func (a *jsiiProxy_Annotations) HasWarning(constructPath *string, message interface{}) {
	_jsii_.InvokeVoid(
		a,
		"hasWarning",
		[]interface{}{constructPath, message},
	)
}

// Capture values while matching templates.
//
// Using an instance of this class within a Matcher will capture the matching value.
// The `as*()` APIs on the instance can be used to get the captured value.
//
// Example:
//   // Given a template -
//   // {
//   //   "Resources": {
//   //     "MyBar": {
//   //       "Type": "Foo::Bar",
//   //       "Properties": {
//   //         "Fred": "Flob",
//   //       }
//   //     },
//   //     "MyBaz": {
//   //       "Type": "Foo::Bar",
//   //       "Properties": {
//   //         "Fred": "Quib",
//   //       }
//   //     }
//   //   }
//   // }
//
//   fredCapture := awscdk.NewCapture()
//   template.hasResourceProperties(jsii.String("Foo::Bar"), map[string]capture{
//   	"Fred": fredCapture,
//   })
//
//   fredCapture.asString() // returns "Flob"
//   fredCapture.next() // returns true
//   fredCapture.asString()
//
// Experimental.
type Capture interface {
	Matcher
	// A name for the matcher.
	//
	// This is collected as part of the result and may be presented to the user.
	// Experimental.
	Name() *string
	// Retrieve the captured value as an array.
	//
	// An error is generated if no value is captured or if the value is not an array.
	// Experimental.
	AsArray() *[]interface{}
	// Retrieve the captured value as a boolean.
	//
	// An error is generated if no value is captured or if the value is not a boolean.
	// Experimental.
	AsBoolean() *bool
	// Retrieve the captured value as a number.
	//
	// An error is generated if no value is captured or if the value is not a number.
	// Experimental.
	AsNumber() *float64
	// Retrieve the captured value as a JSON object.
	//
	// An error is generated if no value is captured or if the value is not an object.
	// Experimental.
	AsObject() *map[string]interface{}
	// Retrieve the captured value as a string.
	//
	// An error is generated if no value is captured or if the value is not a string.
	// Experimental.
	AsString() *string
	// When multiple results are captured, move the iterator to the next result.
	//
	// Returns: true if another capture is present, false otherwise.
	// Experimental.
	Next() *bool
	// Test whether a target matches the provided pattern.
	//
	// Every Matcher must implement this method.
	// This method will be invoked by the assertions framework. Do not call this method directly.
	// Experimental.
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
// Experimental.
func NewCapture(pattern interface{}) Capture {
	_init_.Initialize()

	j := jsiiProxy_Capture{}

	_jsii_.Create(
		"monocdk.assertions.Capture",
		[]interface{}{pattern},
		&j,
	)

	return &j
}

// Initialize a new capture.
// Experimental.
func NewCapture_Override(c Capture, pattern interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.assertions.Capture",
		[]interface{}{pattern},
		c,
	)
}

// Check whether the provided object is a subtype of the `IMatcher`.
// Experimental.
func Capture_IsMatcher(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.assertions.Capture",
		"isMatcher",
		[]interface{}{x},
		&returns,
	)

	return returns
}

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
func Match_Absent() Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"absent",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches any non-null value at the target.
// Experimental.
func Match_AnyValue() Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"anyValue",
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

// Matches any target which does NOT follow the specified pattern.
// Experimental.
func Match_Not(pattern interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"not",
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

// Matches any string-encoded JSON and applies the specified pattern after parsing it.
// Experimental.
func Match_SerializedJson(pattern interface{}) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"serializedJson",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches targets according to a regular expression.
// Experimental.
func Match_StringLikeRegexp(pattern *string) Matcher {
	_init_.Initialize()

	var returns Matcher

	_jsii_.StaticInvoke(
		"monocdk.assertions.Match",
		"stringLikeRegexp",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Information about a value captured during match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var capture capture
//   var value interface{}
//
//   matchCapture := &matchCapture{
//   	capture: capture,
//   	value: value,
//   }
//
// Experimental.
type MatchCapture struct {
	// The instance of Capture class to which this capture is associated with.
	// Experimental.
	Capture Capture `field:"required" json:"capture" yaml:"capture"`
	// The value that was captured.
	// Experimental.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

// Match failure details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var matcher matcher
//
//   matchFailure := &matchFailure{
//   	matcher: matcher,
//   	message: jsii.String("message"),
//   	path: []*string{
//   		jsii.String("path"),
//   	},
//   }
//
// Experimental.
type MatchFailure struct {
	// The matcher that had the failure.
	// Experimental.
	Matcher Matcher `field:"required" json:"matcher" yaml:"matcher"`
	// Failure message.
	// Experimental.
	Message *string `field:"required" json:"message" yaml:"message"`
	// The relative path in the target where the failure occurred.
	//
	// If the failure occurred at root of the match tree, set the path to an empty list.
	// If it occurs in the 5th index of an array nested within the 'foo' key of an object,
	// set the path as `['/foo', '[5]']`.
	// Experimental.
	Path *[]*string `field:"required" json:"path" yaml:"path"`
}

// The result of `Match.test()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var target interface{}
//
//   matchResult := awscdk.Assertions.NewMatchResult(target)
//
// Experimental.
type MatchResult interface {
	// The number of failures.
	// Experimental.
	FailCount() *float64
	// The target for which this result was generated.
	// Experimental.
	Target() interface{}
	// Compose the results of a previous match as a subtree.
	// Experimental.
	Compose(id *string, inner MatchResult) MatchResult
	// Prepare the result to be analyzed.
	//
	// This API *must* be called prior to analyzing these results.
	// Experimental.
	Finished() MatchResult
	// Does the result contain any failures.
	//
	// If not, the result is a success.
	// Experimental.
	HasFailed() *bool
	// DEPRECATED.
	// Deprecated: use recordFailure().
	Push(matcher Matcher, path *[]*string, message *string) MatchResult
	// Record a capture against in this match result.
	// Experimental.
	RecordCapture(options *MatchCapture)
	// Record a new failure into this result at a specific path.
	// Experimental.
	RecordFailure(failure *MatchFailure) MatchResult
	// Get the list of failures as human readable strings.
	// Experimental.
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

func (m *jsiiProxy_MatchResult) RecordCapture(options *MatchCapture) {
	_jsii_.InvokeVoid(
		m,
		"recordCapture",
		[]interface{}{options},
	)
}

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
// Example:
//   // Given a template -
//   // {
//   //   "Resources": {
//   //     "MyBar": {
//   //       "Type": "Foo::Bar",
//   //       "Properties": {
//   //         "Fred": ["Flob", "Cat"]
//   //       }
//   //     }
//   //   }
//   // }
//
//   // The following will NOT throw an assertion error
//   template.hasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
//   	"Fred": awscdk.Match.arrayWith([]interface{}{
//   		jsii.String("Flob"),
//   	}),
//   })
//
//   // The following will throw an assertion error
//   template.hasResourceProperties(jsii.String("Foo::Bar"), awscdk.Match.objectLike(map[string]interface{}{
//   	"Fred": awscdk.Match.arrayWith([]interface{}{
//   		jsii.String("Wobble"),
//   	}),
//   }))
//
// Experimental.
type Matcher interface {
	// A name for the matcher.
	//
	// This is collected as part of the result and may be presented to the user.
	// Experimental.
	Name() *string
	// Test whether a target matches the provided pattern.
	//
	// Every Matcher must implement this method.
	// This method will be invoked by the assertions framework. Do not call this method directly.
	//
	// Returns: the list of match failures. An empty array denotes a successful match.
	// Experimental.
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
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stack := awscdk.NewStack()
//   // ...
//   template := awscdk.Template.fromStack(stack)
//
// Experimental.
type Template interface {
	// Get the set of matching Conditions that match the given properties in the CloudFormation template.
	// Experimental.
	FindConditions(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	// Get the set of matching Mappings that match the given properties in the CloudFormation template.
	// Experimental.
	FindMappings(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	// Get the set of matching Outputs that match the given properties in the CloudFormation template.
	// Experimental.
	FindOutputs(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	// Get the set of matching Parameters that match the given properties in the CloudFormation template.
	// Experimental.
	FindParameters(logicalId *string, props interface{}) *map[string]*map[string]interface{}
	// Get the set of matching resources of a given type and properties in the CloudFormation template.
	// Experimental.
	FindResources(type_ *string, props interface{}) *map[string]*map[string]interface{}
	// Assert that a Condition with the given properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavour, use other matchers in the `Match` class.
	// Experimental.
	HasCondition(logicalId *string, props interface{})
	// Assert that a Mapping with the given properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavour, use other matchers in the `Match` class.
	// Experimental.
	HasMapping(logicalId *string, props interface{})
	// Assert that an Output with the given properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavour, use other matchers in the `Match` class.
	// Experimental.
	HasOutput(logicalId *string, props interface{})
	// Assert that a Parameter with the given properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the parameter, via the `Match.objectLike()`.
	// To configure different behavior, use other matchers in the `Match` class.
	// Experimental.
	HasParameter(logicalId *string, props interface{})
	// Assert that a resource of the given type and given definition exists in the CloudFormation template.
	//
	// By default, performs partial matching on the resource, via the `Match.objectLike()`.
	// To configure different behavour, use other matchers in the `Match` class.
	// Experimental.
	HasResource(type_ *string, props interface{})
	// Assert that a resource of the given type and properties exists in the CloudFormation template.
	//
	// By default, performs partial matching on the `Properties` key of the resource, via the
	// `Match.objectLike()`. To configure different behavour, use other matchers in the `Match` class.
	// Experimental.
	HasResourceProperties(type_ *string, props interface{})
	// Assert that the given number of resources of the given type exist in the template.
	// Experimental.
	ResourceCountIs(type_ *string, count *float64)
	// Assert that the CloudFormation template matches the given value.
	// Experimental.
	TemplateMatches(expected interface{})
	// The CloudFormation template deserialized into an object.
	// Experimental.
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

func (t *jsiiProxy_Template) FindConditions(logicalId *string, props interface{}) *map[string]*map[string]interface{} {
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findConditions",
		[]interface{}{logicalId, props},
		&returns,
	)

	return returns
}

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

func (t *jsiiProxy_Template) FindParameters(logicalId *string, props interface{}) *map[string]*map[string]interface{} {
	var returns *map[string]*map[string]interface{}

	_jsii_.Invoke(
		t,
		"findParameters",
		[]interface{}{logicalId, props},
		&returns,
	)

	return returns
}

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

func (t *jsiiProxy_Template) HasCondition(logicalId *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasCondition",
		[]interface{}{logicalId, props},
	)
}

func (t *jsiiProxy_Template) HasMapping(logicalId *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasMapping",
		[]interface{}{logicalId, props},
	)
}

func (t *jsiiProxy_Template) HasOutput(logicalId *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasOutput",
		[]interface{}{logicalId, props},
	)
}

func (t *jsiiProxy_Template) HasParameter(logicalId *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasParameter",
		[]interface{}{logicalId, props},
	)
}

func (t *jsiiProxy_Template) HasResource(type_ *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasResource",
		[]interface{}{type_, props},
	)
}

func (t *jsiiProxy_Template) HasResourceProperties(type_ *string, props interface{}) {
	_jsii_.InvokeVoid(
		t,
		"hasResourceProperties",
		[]interface{}{type_, props},
	)
}

func (t *jsiiProxy_Template) ResourceCountIs(type_ *string, count *float64) {
	_jsii_.InvokeVoid(
		t,
		"resourceCountIs",
		[]interface{}{type_, count},
	)
}

func (t *jsiiProxy_Template) TemplateMatches(expected interface{}) {
	_jsii_.InvokeVoid(
		t,
		"templateMatches",
		[]interface{}{expected},
	)
}

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

