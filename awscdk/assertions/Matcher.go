package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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
//   template.HasResourceProperties(jsii.String("Foo::Bar"), map[string]matcher{
//   	"Fred": awscdk.Match_arrayWith([]interface{}{
//   		jsii.String("Flob"),
//   	}),
//   })
//
//   // The following will throw an assertion error
//   template.HasResourceProperties(jsii.String("Foo::Bar"), awscdk.Match_ObjectLike(map[string]interface{}{
//   	"Fred": awscdk.Match_arrayWith([]interface{}{
//   		jsii.String("Wobble"),
//   	}),
//   }))
//
type Matcher interface {
	// A name for the matcher.
	//
	// This is collected as part of the result and may be presented to the user.
	Name() *string
	// Test whether a target matches the provided pattern.
	//
	// Every Matcher must implement this method.
	// This method will be invoked by the assertions framework. Do not call this method directly.
	//
	// Returns: the list of match failures. An empty array denotes a successful match.
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

	if err := validateMatcher_IsMatcherParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Matcher",
		"isMatcher",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Matcher) Test(actual interface{}) MatchResult {
	if err := m.validateTestParameters(actual); err != nil {
		panic(err)
	}
	var returns MatchResult

	_jsii_.Invoke(
		m,
		"test",
		[]interface{}{actual},
		&returns,
	)

	return returns
}

