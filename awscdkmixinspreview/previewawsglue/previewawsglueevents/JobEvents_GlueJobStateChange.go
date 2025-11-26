package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.glue@GlueJobStateChange event types for Job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueJobStateChange := #error#.NewGlueJobStateChange()
//
// Experimental.
type JobEvents_GlueJobStateChange interface {
}

// The jsii proxy struct for JobEvents_GlueJobStateChange
type jsiiProxy_JobEvents_GlueJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewJobEvents_GlueJobStateChange() JobEvents_GlueJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_JobEvents_GlueJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewJobEvents_GlueJobStateChange_Override(j JobEvents_GlueJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobStateChange",
		nil, // no parameters
		j,
	)
}

