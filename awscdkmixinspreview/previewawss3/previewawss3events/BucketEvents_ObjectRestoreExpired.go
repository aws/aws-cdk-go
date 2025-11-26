package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectRestoreExpired event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectRestoreExpired := #error#.NewObjectRestoreExpired()
//
// Experimental.
type BucketEvents_ObjectRestoreExpired interface {
}

// The jsii proxy struct for BucketEvents_ObjectRestoreExpired
type jsiiProxy_BucketEvents_ObjectRestoreExpired struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectRestoreExpired() BucketEvents_ObjectRestoreExpired {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectRestoreExpired{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreExpired",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectRestoreExpired_Override(b BucketEvents_ObjectRestoreExpired) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreExpired",
		nil, // no parameters
		b,
	)
}

