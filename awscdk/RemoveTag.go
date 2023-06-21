package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The RemoveTag Aspect will handle removing tags from this node and children.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   removeTag := cdk.NewRemoveTag(jsii.String("key"), &TagProps{
//   	ApplyToLaunchedInstances: jsii.Boolean(false),
//   	ExcludeResourceTypes: []*string{
//   		jsii.String("excludeResourceTypes"),
//   	},
//   	IncludeResourceTypes: []*string{
//   		jsii.String("includeResourceTypes"),
//   	},
//   	Priority: jsii.Number(123),
//   })
//
type RemoveTag interface {
	IAspect
	// The string key for the tag.
	Key() *string
	Props() *TagProps
	ApplyTag(resource ITaggable)
	ApplyTagV2(resource ITaggableV2)
	// All aspects can visit an IConstruct.
	Visit(construct constructs.IConstruct)
}

// The jsii proxy struct for RemoveTag
type jsiiProxy_RemoveTag struct {
	jsiiProxy_IAspect
}

func (j *jsiiProxy_RemoveTag) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RemoveTag) Props() *TagProps {
	var returns *TagProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewRemoveTag(key *string, props *TagProps) RemoveTag {
	_init_.Initialize()

	if err := validateNewRemoveTagParameters(key, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RemoveTag{}

	_jsii_.Create(
		"aws-cdk-lib.RemoveTag",
		[]interface{}{key, props},
		&j,
	)

	return &j
}

func NewRemoveTag_Override(r RemoveTag, key *string, props *TagProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.RemoveTag",
		[]interface{}{key, props},
		r,
	)
}

func (r *jsiiProxy_RemoveTag) ApplyTag(resource ITaggable) {
	if err := r.validateApplyTagParameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyTag",
		[]interface{}{resource},
	)
}

func (r *jsiiProxy_RemoveTag) ApplyTagV2(resource ITaggableV2) {
	if err := r.validateApplyTagV2Parameters(resource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyTagV2",
		[]interface{}{resource},
	)
}

func (r *jsiiProxy_RemoveTag) Visit(construct constructs.IConstruct) {
	if err := r.validateVisitParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"visit",
		[]interface{}{construct},
	)
}

