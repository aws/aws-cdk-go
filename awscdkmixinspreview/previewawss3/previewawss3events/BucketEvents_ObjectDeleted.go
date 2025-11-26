package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectDeleted event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectDeleted := #error#.NewObjectDeleted()
//
// Experimental.
type BucketEvents_ObjectDeleted interface {
}

// The jsii proxy struct for BucketEvents_ObjectDeleted
type jsiiProxy_BucketEvents_ObjectDeleted struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectDeleted() BucketEvents_ObjectDeleted {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectDeleted_Override(b BucketEvents_ObjectDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectDeleted",
		nil, // no parameters
		b,
	)
}

