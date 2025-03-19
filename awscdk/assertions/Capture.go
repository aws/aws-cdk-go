package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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
//   template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]capture{
//   	"Fred": fredCapture,
//   })
//
//   fredCapture.AsString() // returns "Flob"
//   fredCapture.Next() // returns true
//   fredCapture.AsString()
//
type Capture interface {
	Matcher
	// A name for the matcher.
	//
	// This is collected as part of the result and may be presented to the user.
	Name() *string
	// Retrieve the captured value as an array.
	//
	// An error is generated if no value is captured or if the value is not an array.
	AsArray() *[]interface{}
	// Retrieve the captured value as a boolean.
	//
	// An error is generated if no value is captured or if the value is not a boolean.
	AsBoolean() *bool
	// Retrieve the captured value as a number.
	//
	// An error is generated if no value is captured or if the value is not a number.
	AsNumber() *float64
	// Retrieve the captured value as a JSON object.
	//
	// An error is generated if no value is captured or if the value is not an object.
	AsObject() *map[string]interface{}
	// Retrieve the captured value as a string.
	//
	// An error is generated if no value is captured or if the value is not a string.
	AsString() *string
	// When multiple results are captured, move the iterator to the next result.
	//
	// Returns: true if another capture is present, false otherwise.
	Next() *bool
	// Test whether a target matches the provided pattern.
	//
	// Every Matcher must implement this method.
	// This method will be invoked by the assertions framework. Do not call this method directly.
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

	if err := validateCapture_IsMatcherParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Capture",
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
	if err := c.validateTestParameters(actual); err != nil {
		panic(err)
	}
	var returns MatchResult

	_jsii_.Invoke(
		c,
		"test",
		[]interface{}{actual},
		&returns,
	)

	return returns
}

