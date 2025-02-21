package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Modernized version of ITaggable.
//
// `ITaggable` has a problem: for a number of L1 resources, we failed to generate
// `tags: TagManager`, and generated `tags: CfnSomeResource.TagProperty[]` instead.
//
// To mark these resources as taggable, we need to put the `TagManager` in a new property
// whose name is unlikely to conflict with any existing properties. Hence, a new interface
// for that purpose. All future resources will implement `ITaggableV2`.
type ITaggableV2 interface {
	// TagManager to set, remove and format tags.
	CdkTagManager() TagManager
}

// The jsii proxy for ITaggableV2
type jsiiProxy_ITaggableV2 struct {
	_ byte // padding
}

func (j *jsiiProxy_ITaggableV2) CdkTagManager() TagManager {
	var returns TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

