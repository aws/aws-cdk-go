package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectCreated event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectCreated := #error#.NewObjectCreated()
//
// Experimental.
type BucketEvents_ObjectCreated interface {
}

// The jsii proxy struct for BucketEvents_ObjectCreated
type jsiiProxy_BucketEvents_ObjectCreated struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectCreated() BucketEvents_ObjectCreated {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectCreated_Override(b BucketEvents_ObjectCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectCreated",
		nil, // no parameters
		b,
	)
}

