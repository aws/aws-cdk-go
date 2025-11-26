package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectStorageClassChanged event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectStorageClassChanged := #error#.NewObjectStorageClassChanged()
//
// Experimental.
type BucketEvents_ObjectStorageClassChanged interface {
}

// The jsii proxy struct for BucketEvents_ObjectStorageClassChanged
type jsiiProxy_BucketEvents_ObjectStorageClassChanged struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectStorageClassChanged() BucketEvents_ObjectStorageClassChanged {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectStorageClassChanged{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectStorageClassChanged",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectStorageClassChanged_Override(b BucketEvents_ObjectStorageClassChanged) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectStorageClassChanged",
		nil, // no parameters
		b,
	)
}

