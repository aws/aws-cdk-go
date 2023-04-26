package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnJobTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var abortConfig interface{}
//   var jobExecutionsRolloutConfig interface{}
//   var presignedUrlConfig interface{}
//   var timeoutConfig interface{}
//
//   cfnJobTemplateProps := &CfnJobTemplateProps{
//   	Description: jsii.String("description"),
//   	JobTemplateId: jsii.String("jobTemplateId"),
//
//   	// the properties below are optional
//   	AbortConfig: abortConfig,
//   	Document: jsii.String("document"),
//   	DocumentSource: jsii.String("documentSource"),
//   	JobArn: jsii.String("jobArn"),
//   	JobExecutionsRetryConfig: &JobExecutionsRetryConfigProperty{
//   		RetryCriteriaList: []interface{}{
//   			&RetryCriteriaProperty{
//   				FailureType: jsii.String("failureType"),
//   				NumberOfRetries: jsii.Number(123),
//   			},
//   		},
//   	},
//   	JobExecutionsRolloutConfig: jobExecutionsRolloutConfig,
//   	MaintenanceWindows: []interface{}{
//   		&MaintenanceWindowProperty{
//   			DurationInMinutes: jsii.Number(123),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	PresignedUrlConfig: presignedUrlConfig,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutConfig: timeoutConfig,
//   }
//
type CfnJobTemplateProps struct {
	// A description of the job template.
	Description *string `field:"required" json:"description" yaml:"description"`
	// A unique identifier for the job template.
	//
	// We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
	JobTemplateId *string `field:"required" json:"jobTemplateId" yaml:"jobTemplateId"`
	// The criteria that determine when and how a job abort takes place.
	AbortConfig interface{} `field:"optional" json:"abortConfig" yaml:"abortConfig"`
	// The job document.
	//
	// Required if you don't specify a value for `documentSource` .
	Document *string `field:"optional" json:"document" yaml:"document"`
	// An S3 link to the job document to use in the template.
	//
	// Required if you don't specify a value for `document` .
	//
	// > If the job document resides in an S3 bucket, you must use a placeholder link when specifying the document.
	// >
	// > The placeholder link is of the following form:
	// >
	// > `${aws:iot:s3-presigned-url:https://s3.amazonaws.com/ *bucket* / *key* }`
	// >
	// > where *bucket* is your bucket name and *key* is the object in the bucket to which you are linking.
	DocumentSource *string `field:"optional" json:"documentSource" yaml:"documentSource"`
	// The ARN of the job to use as the basis for the job template.
	JobArn *string `field:"optional" json:"jobArn" yaml:"jobArn"`
	// Allows you to create the criteria to retry a job.
	JobExecutionsRetryConfig interface{} `field:"optional" json:"jobExecutionsRetryConfig" yaml:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig interface{} `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// An optional configuration within the SchedulingConfig to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
	MaintenanceWindows interface{} `field:"optional" json:"maintenanceWindows" yaml:"maintenanceWindows"`
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig interface{} `field:"optional" json:"presignedUrlConfig" yaml:"presignedUrlConfig"`
	// Metadata that can be used to manage the job template.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	//
	// A timer is started when the job execution status is set to `IN_PROGRESS` . If the job execution status is not set to another terminal state before the timer expires, it will be automatically set to `TIMED_OUT` .
	TimeoutConfig interface{} `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

