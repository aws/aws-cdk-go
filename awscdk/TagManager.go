package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// TagManager facilitates a common implementation of tagging for Constructs.
//
// Normally, you do not need to use this class, as the CloudFormation specification
// will indicate which resources are taggable. However, sometimes you will need this
// to make custom resources taggable. Used `tagManager.renderedTags` to obtain a
// value that will resolve to the tags at synthesis time.
//
// Example:
//   type myConstruct struct {
//   	resource
//   	tags
//   }
//
//   func newMyConstruct(scope construct, id *string) *myConstruct {
//   	this := &myConstruct{}
//   	newResource_Override(this, scope, id)
//
//   	awscdk.NewCfnResource(this, jsii.String("Resource"), &cfnResourceProps{
//   		Type: jsii.String("Whatever::The::Type"),
//   		Properties: map[string]interface{}{
//   			// ...
//   			"Tags": this.tags.renderedTags,
//   		},
//   	})
//   	return this
//   }
//
type TagManager interface {
	// A lazy value that represents the rendered tags at synthesis time.
	//
	// If you need to make a custom construct taggable, use the value of this
	// property to pass to the `tags` property of the underlying construct.
	RenderedTags() IResolvable
	// The property name for tag values.
	//
	// Normally this is `tags` but some resources choose a different name. Cognito
	// UserPool uses UserPoolTags.
	TagPropertyName() *string
	// Determine if the aspect applies here.
	//
	// Looks at the include and exclude resourceTypeName arrays to determine if
	// the aspect applies here.
	ApplyTagAspectHere(include *[]*string, exclude *[]*string) *bool
	// Returns true if there are any tags defined.
	HasTags() *bool
	// Removes the specified tag from the array if it exists.
	RemoveTag(key *string, priority *float64)
	// Renders tags into the proper format based on TagType.
	//
	// This method will eagerly render the tags currently applied. In
	// most cases, you should be using `tagManager.renderedTags` instead,
	// which will return a `Lazy` value that will resolve to the correct
	// tags at synthesis time.
	RenderTags(combineWithTags interface{}) interface{}
	// Adds the specified tag to the array of tags.
	SetTag(key *string, value *string, priority *float64, applyToLaunchedInstances *bool)
	// Render the tags in a readable format.
	TagValues() *map[string]*string
}

// The jsii proxy struct for TagManager
type jsiiProxy_TagManager struct {
	_ byte // padding
}

func (j *jsiiProxy_TagManager) RenderedTags() IResolvable {
	var returns IResolvable
	_jsii_.Get(
		j,
		"renderedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagManager) TagPropertyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPropertyName",
		&returns,
	)
	return returns
}


func NewTagManager(tagType TagType, resourceTypeName *string, initialTags interface{}, options *TagManagerOptions) TagManager {
	_init_.Initialize()

	if err := validateNewTagManagerParameters(tagType, resourceTypeName, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TagManager{}

	_jsii_.Create(
		"aws-cdk-lib.TagManager",
		[]interface{}{tagType, resourceTypeName, initialTags, options},
		&j,
	)

	return &j
}

func NewTagManager_Override(t TagManager, tagType TagType, resourceTypeName *string, initialTags interface{}, options *TagManagerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.TagManager",
		[]interface{}{tagType, resourceTypeName, initialTags, options},
		t,
	)
}

// Check whether the given construct is Taggable.
func TagManager_IsTaggable(construct interface{}) *bool {
	_init_.Initialize()

	if err := validateTagManager_IsTaggableParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.TagManager",
		"isTaggable",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is ITaggableV2.
func TagManager_IsTaggableV2(construct interface{}) *bool {
	_init_.Initialize()

	if err := validateTagManager_IsTaggableV2Parameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.TagManager",
		"isTaggableV2",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the TagManager associated with the given construct, if any.
func TagManager_Of(construct interface{}) TagManager {
	_init_.Initialize()

	if err := validateTagManager_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns TagManager

	_jsii_.StaticInvoke(
		"aws-cdk-lib.TagManager",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagManager) ApplyTagAspectHere(include *[]*string, exclude *[]*string) *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"applyTagAspectHere",
		[]interface{}{include, exclude},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagManager) HasTags() *bool {
	var returns *bool

	_jsii_.Invoke(
		t,
		"hasTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagManager) RemoveTag(key *string, priority *float64) {
	if err := t.validateRemoveTagParameters(key, priority); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"removeTag",
		[]interface{}{key, priority},
	)
}

func (t *jsiiProxy_TagManager) RenderTags(combineWithTags interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderTags",
		[]interface{}{combineWithTags},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagManager) SetTag(key *string, value *string, priority *float64, applyToLaunchedInstances *bool) {
	if err := t.validateSetTagParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"setTag",
		[]interface{}{key, value, priority, applyToLaunchedInstances},
	)
}

func (t *jsiiProxy_TagManager) TagValues() *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"tagValues",
		nil, // no parameters
		&returns,
	)

	return returns
}

