package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectRestoreCompleted event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectRestoreCompleted := #error#.NewObjectRestoreCompleted()
//
// Experimental.
type BucketEvents_ObjectRestoreCompleted interface {
}

// The jsii proxy struct for BucketEvents_ObjectRestoreCompleted
type jsiiProxy_BucketEvents_ObjectRestoreCompleted struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectRestoreCompleted() BucketEvents_ObjectRestoreCompleted {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectRestoreCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectRestoreCompleted_Override(b BucketEvents_ObjectRestoreCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreCompleted",
		nil, // no parameters
		b,
	)
}

