package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.glue@GlueJobRunStatus event types for Job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueJobRunStatus := #error#.NewGlueJobRunStatus()
//
// Experimental.
type JobEvents_GlueJobRunStatus interface {
}

// The jsii proxy struct for JobEvents_GlueJobRunStatus
type jsiiProxy_JobEvents_GlueJobRunStatus struct {
	_ byte // padding
}

// Experimental.
func NewJobEvents_GlueJobRunStatus() JobEvents_GlueJobRunStatus {
	_init_.Initialize()

	j := jsiiProxy_JobEvents_GlueJobRunStatus{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobRunStatus",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewJobEvents_GlueJobRunStatus_Override(j JobEvents_GlueJobRunStatus) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobRunStatus",
		nil, // no parameters
		j,
	)
}

