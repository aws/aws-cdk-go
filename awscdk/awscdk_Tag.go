// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The Tag Aspect will handle adding a tag to this node and cascading tags to children.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   tag := cdk.NewTag(jsii.String("key"), jsii.String("value"), &tagProps{
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
type Tag interface {
	IAspect
	// The string key for the tag.
	Key() *string
	Props() *TagProps
	// The string value of the tag.
	Value() *string
	ApplyTag(resource ITaggable)
	// All aspects can visit an IConstruct.
	Visit(construct constructs.IConstruct)
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


func NewTag(key *string, value *string, props *TagProps) Tag {
	_init_.Initialize()

	if err := validateNewTagParameters(key, value, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Tag{}

	_jsii_.Create(
		"aws-cdk-lib.Tag",
		[]interface{}{key, value, props},
		&j,
	)

	return &j
}

func NewTag_Override(t Tag, key *string, value *string, props *TagProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Tag",
		[]interface{}{key, value, props},
		t,
	)
}

func (t *jsiiProxy_Tag) ApplyTag(resource ITaggable) {
	if err := t.validateApplyTagParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyTag",
		[]interface{}{resource},
	)
}

func (t *jsiiProxy_Tag) Visit(construct constructs.IConstruct) {
	if err := t.validateVisitParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"visit",
		[]interface{}{construct},
	)
}

