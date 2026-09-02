package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnJobProps := &CfnJobProps{
//   	JobId: jsii.String("jobId"),
//   	Targets: []*string{
//   		jsii.String("targets"),
//   	},
//
//   	// the properties below are optional
//   	AbortConfig: &AbortConfigProperty{
//   		CriteriaList: []interface{}{
//   			&AbortCriteriaProperty{
//   				Action: jsii.String("action"),
//   				FailureType: jsii.String("failureType"),
//   				MinNumberOfExecutedThings: jsii.Number(123),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DestinationPackageVersions: []*string{
//   		jsii.String("destinationPackageVersions"),
//   	},
//   	Document: jsii.String("document"),
//   	DocumentParameters: map[string]*string{
//   		"documentParametersKey": jsii.String("documentParameters"),
//   	},
//   	DocumentSource: jsii.String("documentSource"),
//   	JobExecutionsRetryConfig: &JobExecutionsRetryConfigProperty{
//   		CriteriaList: []interface{}{
//   			&RetryCriteriaProperty{
//   				FailureType: jsii.String("failureType"),
//   				NumberOfRetries: jsii.Number(123),
//   			},
//   		},
//   	},
//   	JobExecutionsRolloutConfig: &JobExecutionsRolloutConfigProperty{
//   		ExponentialRate: &ExponentialRolloutRateProperty{
//   			BaseRatePerMinute: jsii.Number(123),
//   			IncrementFactor: jsii.Number(123),
//   			RateIncreaseCriteria: &RateIncreaseCriteriaProperty{
//   				NumberOfNotifiedThings: jsii.Number(123),
//   				NumberOfSucceededThings: jsii.Number(123),
//   			},
//   		},
//   		MaximumPerMinute: jsii.Number(123),
//   	},
//   	JobTemplateArn: jsii.String("jobTemplateArn"),
//   	PresignedUrlConfig: &PresignedUrlConfigProperty{
//   		ExpiresInSec: jsii.Number(123),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SchedulingConfig: &SchedulingConfigProperty{
//   		EndBehavior: jsii.String("endBehavior"),
//   		EndTime: jsii.String("endTime"),
//   		MaintenanceWindows: []interface{}{
//   			&MaintenanceWindowProperty{
//   				DurationInMinutes: jsii.Number(123),
//   				StartTime: jsii.String("startTime"),
//   			},
//   		},
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetSelection: jsii.String("targetSelection"),
//   	TimeoutConfig: &TimeoutConfigProperty{
//   		InProgressTimeoutInMinutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html
//
type CfnJobProps struct {
	// A job identifier which must be unique for your AWS account.
	//
	// We recommend using a UUID. Alpha-numeric characters, '-' and '_' are valid for use here.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-jobid
	//
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// A list of things and thing groups to which the job should be sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-targets
	//
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// The criteria that determine when and how a job abort takes place.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-abortconfig
	//
	AbortConfig interface{} `field:"optional" json:"abortConfig" yaml:"abortConfig"`
	// A short text description of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The package version Amazon Resource Names (ARNs) that are installed on the device when the job successfully completes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-destinationpackageversions
	//
	DestinationPackageVersions *[]*string `field:"optional" json:"destinationPackageVersions" yaml:"destinationPackageVersions"`
	// The job document.
	//
	// Required if you don't specify a value for documentSource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-document
	//
	Document *string `field:"optional" json:"document" yaml:"document"`
	// Parameters of an Amazon Web Services managed template that you can specify to create the job document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-documentparameters
	//
	DocumentParameters interface{} `field:"optional" json:"documentParameters" yaml:"documentParameters"`
	// An S3 link, or S3 object URL, to the job document.
	//
	// The link is an Amazon S3 object URL and is required if you don't specify a value for document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-documentsource
	//
	DocumentSource *string `field:"optional" json:"documentSource" yaml:"documentSource"`
	// The configuration that determines how many retries are allowed for each failure type for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-jobexecutionsretryconfig
	//
	JobExecutionsRetryConfig interface{} `field:"optional" json:"jobExecutionsRetryConfig" yaml:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-jobexecutionsrolloutconfig
	//
	JobExecutionsRolloutConfig interface{} `field:"optional" json:"jobExecutionsRolloutConfig" yaml:"jobExecutionsRolloutConfig"`
	// The ARN of the job template used to create the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-jobtemplatearn
	//
	JobTemplateArn *string `field:"optional" json:"jobTemplateArn" yaml:"jobTemplateArn"`
	// Configuration for pre-signed S3 URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-presignedurlconfig
	//
	PresignedUrlConfig interface{} `field:"optional" json:"presignedUrlConfig" yaml:"presignedUrlConfig"`
	// Specifies the date and time that a job will begin the rollout of the job document to all devices in the target group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-schedulingconfig
	//
	SchedulingConfig interface{} `field:"optional" json:"schedulingConfig" yaml:"schedulingConfig"`
	// Metadata which can be used to manage the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether the job will continue to run (CONTINUOUS), or will be complete after all those things specified as targets have completed the job (SNAPSHOT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-targetselection
	//
	TargetSelection *string `field:"optional" json:"targetSelection" yaml:"targetSelection"`
	// Specifies the amount of time each device has to finish its execution of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html#cfn-iot-job-timeoutconfig
	//
	TimeoutConfig interface{} `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

