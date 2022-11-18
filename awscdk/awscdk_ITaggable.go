// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface to implement tags.
type ITaggable interface {
	// TagManager to set, remove and format tags.
	Tags() TagManager
}

// The jsii proxy for ITaggable
type jsiiProxy_ITaggable struct {
	_ byte // padding
}

func (j *jsiiProxy_ITaggable) Tags() TagManager {
	var returns TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

