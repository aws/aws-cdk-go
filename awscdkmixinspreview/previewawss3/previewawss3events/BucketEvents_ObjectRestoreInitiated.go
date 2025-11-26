package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectRestoreInitiated event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectRestoreInitiated := #error#.NewObjectRestoreInitiated()
//
// Experimental.
type BucketEvents_ObjectRestoreInitiated interface {
}

// The jsii proxy struct for BucketEvents_ObjectRestoreInitiated
type jsiiProxy_BucketEvents_ObjectRestoreInitiated struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectRestoreInitiated() BucketEvents_ObjectRestoreInitiated {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectRestoreInitiated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreInitiated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectRestoreInitiated_Override(b BucketEvents_ObjectRestoreInitiated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreInitiated",
		nil, // no parameters
		b,
	)
}

