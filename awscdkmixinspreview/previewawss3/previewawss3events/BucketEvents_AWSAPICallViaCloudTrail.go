package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.s3@AWSAPICallViaCloudTrail event types for Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := #error#.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type BucketEvents_AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for BucketEvents_AWSAPICallViaCloudTrail
type jsiiProxy_BucketEvents_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewBucketEvents_AWSAPICallViaCloudTrail() BucketEvents_AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_BucketEvents_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBucketEvents_AWSAPICallViaCloudTrail_Override(b BucketEvents_AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		b,
	)
}

