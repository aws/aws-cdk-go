package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@ObjectACLUpdated event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectACLUpdated := #error#.NewObjectACLUpdated()
//
// Experimental.
type BucketEvents_ObjectACLUpdated interface {
}

// The jsii proxy struct for BucketEvents_ObjectACLUpdated
type jsiiProxy_BucketEvents_ObjectACLUpdated struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_ObjectACLUpdated() BucketEvents_ObjectACLUpdated {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_ObjectACLUpdated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectACLUpdated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_ObjectACLUpdated_Override(b BucketEvents_ObjectACLUpdated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectACLUpdated",
		nil, // no parameters
		b,
	)
}

