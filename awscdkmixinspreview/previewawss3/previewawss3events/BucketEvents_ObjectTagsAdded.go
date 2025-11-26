package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectTagsAdded event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectTagsAdded := #error#.NewObjectTagsAdded()
//
// Experimental.
type BucketEvents_ObjectTagsAdded interface {
}

// The jsii proxy struct for BucketEvents_ObjectTagsAdded
type jsiiProxy_BucketEvents_ObjectTagsAdded struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectTagsAdded() BucketEvents_ObjectTagsAdded {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectTagsAdded{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsAdded",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectTagsAdded_Override(b BucketEvents_ObjectTagsAdded) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsAdded",
		nil, // no parameters
		b,
	)
}

