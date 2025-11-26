package previewawsglueevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

// EventBridge event patterns for Job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jobRef IJobRef
//
//   jobEvents := awscdkmixinspreview.Events.JobEvents_FromJob(jobRef)
//
// Experimental.
type JobEvents interface {
	// EventBridge event pattern for Job AWS API Call via CloudTrail.
	// Experimental.
	AwsAPICallViaCloudTrailPattern(options *JobEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
	// EventBridge event pattern for Job Glue Job Run Status.
	// Experimental.
	GlueJobRunStatusPattern(options *JobEvents_GlueJobRunStatus_GlueJobRunStatusProps) *awsevents.EventPattern
	// EventBridge event pattern for Job Glue Job State Change.
	// Experimental.
	GlueJobStateChangePattern(options *JobEvents_GlueJobStateChange_GlueJobStateChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for JobEvents
type jsiiProxy_JobEvents struct {
	_ byte // padding
}

// Create JobEvents from a Job reference.
// Experimental.
func JobEvents_FromJob(jobRef interfacesawsglue.IJobRef) JobEvents {
	_init_.Initialize()

	if err := validateJobEvents_FromJobParameters(jobRef); err != nil {
		panic(err)
	}
	var returns JobEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents",
		"fromJob",
		[]interface{}{jobRef},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobEvents) AwsAPICallViaCloudTrailPattern(options *JobEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	if err := j.validateAwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		j,
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobEvents) GlueJobRunStatusPattern(options *JobEvents_GlueJobRunStatus_GlueJobRunStatusProps) *awsevents.EventPattern {
	if err := j.validateGlueJobRunStatusPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		j,
		"glueJobRunStatusPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobEvents) GlueJobStateChangePattern(options *JobEvents_GlueJobStateChange_GlueJobStateChangeProps) *awsevents.EventPattern {
	if err := j.validateGlueJobStateChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		j,
		"glueJobStateChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

