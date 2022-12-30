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
//   cfnJobTemplateProps := &cfnJobTemplateProps{
//   	description: jsii.String("description"),
//   	jobTemplateId: jsii.String("jobTemplateId"),
//
//   	// the properties below are optional
//   	abortConfig: abortConfig,
//   	document: jsii.String("document"),
//   	documentSource: jsii.String("documentSource"),
//   	jobArn: jsii.String("jobArn"),
//   	jobExecutionsRolloutConfig: jobExecutionsRolloutConfig,
//   	presignedUrlConfig: presignedUrlConfig,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeoutConfig: timeoutConfig,
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
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig interface{} `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig interface{} `field:"optional" json:"presignedUrlConfig" yaml:"presignedUrlConfig"`
	// Metadata that can be used to manage the job template.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	//
	// A timer is started when the job execution status is set to `IN_PROGRESS` . If the job execution status is not set to another terminal state before the timer expires, it will be automatically set to `TIMED_OUT` .
	TimeoutConfig interface{} `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

