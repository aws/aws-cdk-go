package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectTagsDeleted event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectTagsDeleted := #error#.NewObjectTagsDeleted()
//
// Experimental.
type BucketEvents_ObjectTagsDeleted interface {
}

// The jsii proxy struct for BucketEvents_ObjectTagsDeleted
type jsiiProxy_BucketEvents_ObjectTagsDeleted struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectTagsDeleted() BucketEvents_ObjectTagsDeleted {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectTagsDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectTagsDeleted_Override(b BucketEvents_ObjectTagsDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsDeleted",
		nil, // no parameters
		b,
	)
}

