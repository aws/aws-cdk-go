// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The Tag Aspect will handle adding a tag to this node and cascading tags to children.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   tag := monocdk.NewTag(jsii.String("key"), jsii.String("value"), &tagProps{
//   	applyToLaunchedInstances: jsii.Boolean(false),
//   	excludeResourceTypes: []*string{
//   		jsii.String("excludeResourceTypes"),
//   	},
//   	includeResourceTypes: []*string{
//   		jsii.String("includeResourceTypes"),
//   	},
//   	priority: jsii.Number(123),
//   })
//
// Experimental.
type Tag interface {
	IAspect
	// The string key for the tag.
	// Experimental.
	Key() *string
	// Experimental.
	Props() *TagProps
	// The string value of the tag.
	// Experimental.
	Value() *string
	// Experimental.
	ApplyTag(resource ITaggable)
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(construct IConstruct)
}

// The jsii proxy struct for Tag
type jsiiProxy_Tag struct {
	jsiiProxy_IAspect
}

func (j *jsiiProxy_Tag) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tag) Props() *TagProps {
	var returns *TagProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Tag) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewTag(key *string, value *string, props *TagProps) Tag {
	_init_.Initialize()

	j := jsiiProxy_Tag{}

	_jsii_.Create(
		"monocdk.Tag",
		[]interface{}{key, value, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTag_Override(t Tag, key *string, value *string, props *TagProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.Tag",
		[]interface{}{key, value, props},
		t,
	)
}

// DEPRECATED: add tags to the node of a construct and all its the taggable children.
// Deprecated: use `Tags.of(scope).add()`
func Tag_Add(scope Construct, key *string, value *string, props *TagProps) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.Tag",
		"add",
		[]interface{}{scope, key, value, props},
	)
}

// DEPRECATED: remove tags to the node of a construct and all its the taggable children.
// Deprecated: use `Tags.of(scope).remove()`
func Tag_Remove(scope Construct, key *string, props *TagProps) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.Tag",
		"remove",
		[]interface{}{scope, key, props},
	)
}

func (t *jsiiProxy_Tag) ApplyTag(resource ITaggable) {
	_jsii_.InvokeVoid(
		t,
		"applyTag",
		[]interface{}{resource},
	)
}

func (t *jsiiProxy_Tag) Visit(construct IConstruct) {
	_jsii_.InvokeVoid(
		t,
		"visit",
		[]interface{}{construct},
	)
}

