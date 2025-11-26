package previewawsglueevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Job aws.glue@GlueJobRunStatus event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   glueJobRunStatusProps := &GlueJobRunStatusProps{
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	JobName: []*string{
//   		jsii.String("jobName"),
//   	},
//   	JobRunId: []*string{
//   		jsii.String("jobRunId"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   	NotificationCondition: &NotificationCondition{
//   		NotifyDelayAfter: []*string{
//   			jsii.String("notifyDelayAfter"),
//   		},
//   	},
//   	Severity: []*string{
//   		jsii.String("severity"),
//   	},
//   	StartedOn: []*string{
//   		jsii.String("startedOn"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   }
//
// Experimental.
type JobEvents_GlueJobRunStatus_GlueJobRunStatusProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// jobName property.
	//
	// Specify an array of string values to match this event if the actual value of jobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Job reference.
	//
	// Experimental.
	JobName *[]*string `field:"optional" json:"jobName" yaml:"jobName"`
	// jobRunId property.
	//
	// Specify an array of string values to match this event if the actual value of jobRunId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobRunId *[]*string `field:"optional" json:"jobRunId" yaml:"jobRunId"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
	// notificationCondition property.
	//
	// Specify an array of string values to match this event if the actual value of notificationCondition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotificationCondition *JobEvents_GlueJobRunStatus_NotificationCondition `field:"optional" json:"notificationCondition" yaml:"notificationCondition"`
	// severity property.
	//
	// Specify an array of string values to match this event if the actual value of severity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Severity *[]*string `field:"optional" json:"severity" yaml:"severity"`
	// startedOn property.
	//
	// Specify an array of string values to match this event if the actual value of startedOn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedOn *[]*string `field:"optional" json:"startedOn" yaml:"startedOn"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
}

