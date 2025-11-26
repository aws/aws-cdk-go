package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectAccessTierChanged event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectAccessTierChanged := #error#.NewObjectAccessTierChanged()
//
// Experimental.
type BucketEvents_ObjectAccessTierChanged interface {
}

// The jsii proxy struct for BucketEvents_ObjectAccessTierChanged
type jsiiProxy_BucketEvents_ObjectAccessTierChanged struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectAccessTierChanged() BucketEvents_ObjectAccessTierChanged {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectAccessTierChanged{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectAccessTierChanged",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectAccessTierChanged_Override(b BucketEvents_ObjectAccessTierChanged) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectAccessTierChanged",
		nil, // no parameters
		b,
	)
}

