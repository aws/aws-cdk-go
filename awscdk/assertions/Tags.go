package assertions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Allows assertions on the tags associated with a synthesized CDK stack's manifest.
//
// Stack tags are not part of the synthesized template, so can only be
// checked from the manifest in this manner.
//
// Example:
//   tags := awscdk.Tags_FromStack(stack)
//
//   // using a default 'objectLike' Matcher
//   tags.HasValues(map[string]*string{
//   	"tag-name": jsii.String("tag-value"),
//   })
//
//   // ... with Matchers embedded
//   tags.HasValues(map[string]matcher{
//   	"tag-name": awscdk.Match_stringLikeRegexp(jsii.String("value")),
//   })
//
//   // or another object Matcher at the top level
//   tags.HasValues(awscdk.Match_ObjectEquals(map[string]interface{}{
//   	"tag-name": awscdk.Match_anyValue(),
//   }))
//
type Tags interface {
	// Get the tags associated with the manifest.
	//
	// This will be an empty object if
	// no tags were supplied.
	//
	// Returns: The tags associated with the stack's synthesized manifest.
	All() *map[string]*string
	// Assert that the there are no tags associated with the synthesized CDK Stack's manifest.
	//
	// This is a convenience method over `hasValues(Match.exact({}))`, and is
	// present because the more obvious method of detecting no tags
	// (`Match.absent()`) will not work. Manifests default the tag set to an empty
	// object.
	HasNone()
	// Assert that the given Matcher or object matches the tags associated with the synthesized CDK Stack's manifest.
	HasValues(tags interface{})
}

// The jsii proxy struct for Tags
type jsiiProxy_Tags struct {
	_ byte // padding
}

// Find tags associated with a synthesized CDK `Stack`.
func Tags_FromStack(stack awscdk.Stack) Tags {
	_init_.Initialize()

	if err := validateTags_FromStackParameters(stack); err != nil {
		panic(err)
	}
	var returns Tags

	_jsii_.StaticInvoke(
		"aws-cdk-lib.assertions.Tags",
		"fromStack",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Tags) All() *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Tags) HasNone() {
	_jsii_.InvokeVoid(
		t,
		"hasNone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Tags) HasValues(tags interface{}) {
	if err := t.validateHasValuesParameters(tags); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"hasValues",
		[]interface{}{tags},
	)
}

