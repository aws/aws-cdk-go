package previewawsiotmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnJobTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var abortConfig interface{}
//   var jobExecutionsRolloutConfig interface{}
//   var presignedUrlConfig interface{}
//   var timeoutConfig interface{}
//
//   cfnJobTemplateMixinProps := &CfnJobTemplateMixinProps{
//   	AbortConfig: abortConfig,
//   	Description: jsii.String("description"),
//   	DestinationPackageVersions: []*string{
//   		jsii.String("destinationPackageVersions"),
//   	},
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
//   	JobTemplateId: jsii.String("jobTemplateId"),
//   	MaintenanceWindows: []interface{}{
//   		&MaintenanceWindowProperty{
//   			DurationInMinutes: jsii.Number(123),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	PresignedUrlConfig: presignedUrlConfig,
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutConfig: timeoutConfig,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html
//
type CfnJobTemplateMixinProps struct {
	// The criteria that determine when and how a job abort takes place.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-abortconfig
	//
	AbortConfig interface{} `field:"optional" json:"abortConfig" yaml:"abortConfig"`
	// A description of the job template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The package version Amazon Resource Names (ARNs) that are installed on the device’s reserved named shadow ( `$package` ) when the job successfully completes.
	//
	// *Note:* Up to 25 package version ARNS are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-destinationpackageversions
	//
	DestinationPackageVersions *[]*string `field:"optional" json:"destinationPackageVersions" yaml:"destinationPackageVersions"`
	// The job document.
	//
	// Required if you don't specify a value for `documentSource` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-document
	//
	Document *string `field:"optional" json:"document" yaml:"document"`
	// An S3 link, or S3 object URL, to the job document.
	//
	// The link is an Amazon S3 object URL and is required if you don't specify a value for `document` .
	//
	// For example, `--document-source https://s3. *region-code* .amazonaws.com/example-firmware/device-firmware.1.0`
	//
	// For more information, see [Methods for accessing a bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-bucket-intro.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-documentsource
	//
	DocumentSource *string `field:"optional" json:"documentSource" yaml:"documentSource"`
	// The ARN of the job to use as the basis for the job template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-jobarn
	//
	JobArn *string `field:"optional" json:"jobArn" yaml:"jobArn"`
	// Allows you to create the criteria to retry a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-jobexecutionsretryconfig
	//
	JobExecutionsRetryConfig interface{} `field:"optional" json:"jobExecutionsRetryConfig" yaml:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-jobexecutionsrolloutconfig
	//
	JobExecutionsRolloutConfig interface{} `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// A unique identifier for the job template.
	//
	// We recommend using a UUID. Alpha-numeric characters, "-", and "_" are valid for use here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-jobtemplateid
	//
	JobTemplateId *string `field:"optional" json:"jobTemplateId" yaml:"jobTemplateId"`
	// An optional configuration within the SchedulingConfig to setup a recurring maintenance window with a predetermined start time and duration for the rollout of a job document to all devices in a target group for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-maintenancewindows
	//
	MaintenanceWindows interface{} `field:"optional" json:"maintenanceWindows" yaml:"maintenanceWindows"`
	// Configuration for pre-signed S3 URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-presignedurlconfig
	//
	PresignedUrlConfig interface{} `field:"optional" json:"presignedUrlConfig" yaml:"presignedUrlConfig"`
	// Metadata that can be used to manage the job template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	//
	// A timer is started when the job execution status is set to `IN_PROGRESS` . If the job execution status is not set to another terminal state before the timer expires, it will be automatically set to `TIMED_OUT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-jobtemplate.html#cfn-iot-jobtemplate-timeoutconfig
	//
	TimeoutConfig interface{} `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

