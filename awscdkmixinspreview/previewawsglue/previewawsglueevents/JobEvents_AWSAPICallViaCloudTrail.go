package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.glue@AWSAPICallViaCloudTrail event types for Job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := #error#.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type JobEvents_AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for JobEvents_AWSAPICallViaCloudTrail
type jsiiProxy_JobEvents_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewJobEvents_AWSAPICallViaCloudTrail() JobEvents_AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_JobEvents_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewJobEvents_AWSAPICallViaCloudTrail_Override(j JobEvents_AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		j,
	)
}

