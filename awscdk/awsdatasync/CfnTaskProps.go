package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskProps := &CfnTaskProps{
//   	DestinationLocationArn: jsii.String("destinationLocationArn"),
//   	SourceLocationArn: jsii.String("sourceLocationArn"),
//
//   	// the properties below are optional
//   	CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	Excludes: []interface{}{
//   		&FilterRuleProperty{
//   			FilterType: jsii.String("filterType"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Includes: []interface{}{
//   		&FilterRuleProperty{
//   			FilterType: jsii.String("filterType"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ManifestConfig: &ManifestConfigProperty{
//   		Source: &SourceProperty{
//   			S3: &ManifestConfigSourceS3Property{
//   				BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   				ManifestObjectPath: jsii.String("manifestObjectPath"),
//   				ManifestObjectVersionId: jsii.String("manifestObjectVersionId"),
//   				S3BucketArn: jsii.String("s3BucketArn"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Action: jsii.String("action"),
//   		Format: jsii.String("format"),
//   	},
//   	Name: jsii.String("name"),
//   	Options: &OptionsProperty{
//   		Atime: jsii.String("atime"),
//   		BytesPerSecond: jsii.Number(123),
//   		Gid: jsii.String("gid"),
//   		LogLevel: jsii.String("logLevel"),
//   		Mtime: jsii.String("mtime"),
//   		ObjectTags: jsii.String("objectTags"),
//   		OverwriteMode: jsii.String("overwriteMode"),
//   		PosixPermissions: jsii.String("posixPermissions"),
//   		PreserveDeletedFiles: jsii.String("preserveDeletedFiles"),
//   		PreserveDevices: jsii.String("preserveDevices"),
//   		SecurityDescriptorCopyFlags: jsii.String("securityDescriptorCopyFlags"),
//   		TaskQueueing: jsii.String("taskQueueing"),
//   		TransferMode: jsii.String("transferMode"),
//   		Uid: jsii.String("uid"),
//   		VerifyMode: jsii.String("verifyMode"),
//   	},
//   	Schedule: &TaskScheduleProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   		Status: jsii.String("status"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskReportConfig: &TaskReportConfigProperty{
//   		Destination: &DestinationProperty{
//   			S3: &S3Property{
//   				BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   				S3BucketArn: jsii.String("s3BucketArn"),
//   				Subdirectory: jsii.String("subdirectory"),
//   			},
//   		},
//   		OutputType: jsii.String("outputType"),
//
//   		// the properties below are optional
//   		ObjectVersionIds: jsii.String("objectVersionIds"),
//   		Overrides: &OverridesProperty{
//   			Deleted: &DeletedProperty{
//   				ReportLevel: jsii.String("reportLevel"),
//   			},
//   			Skipped: &SkippedProperty{
//   				ReportLevel: jsii.String("reportLevel"),
//   			},
//   			Transferred: &TransferredProperty{
//   				ReportLevel: jsii.String("reportLevel"),
//   			},
//   			Verified: &VerifiedProperty{
//   				ReportLevel: jsii.String("reportLevel"),
//   			},
//   		},
//   		ReportLevel: jsii.String("reportLevel"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html
//
type CfnTaskProps struct {
	// The Amazon Resource Name (ARN) of an AWS storage resource's location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-destinationlocationarn
	//
	DestinationLocationArn *string `field:"required" json:"destinationLocationArn" yaml:"destinationLocationArn"`
	// The Amazon Resource Name (ARN) of the source location for the task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-sourcelocationarn
	//
	SourceLocationArn *string `field:"required" json:"sourceLocationArn" yaml:"sourceLocationArn"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that is used to monitor and log events in the task.
	//
	// For more information about how to use CloudWatch Logs with DataSync, see [Monitoring Your Task](https://docs.aws.amazon.com/datasync/latest/userguide/monitor-datasync.html#cloudwatchlogs) in the *AWS DataSync User Guide.*
	//
	// For more information about these groups, see [Working with Log Groups and Log Streams](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html) in the *Amazon CloudWatch Logs User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-cloudwatchloggrouparn
	//
	CloudWatchLogGroupArn *string `field:"optional" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// Specifies a list of filter rules that exclude specific data during your transfer.
	//
	// For more information and examples, see [Filtering data transferred by DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-excludes
	//
	Excludes interface{} `field:"optional" json:"excludes" yaml:"excludes"`
	// Specifies a list of filter rules that include specific data during your transfer.
	//
	// For more information and examples, see [Filtering data transferred by DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-includes
	//
	Includes interface{} `field:"optional" json:"includes" yaml:"includes"`
	// The configuration of the manifest that lists the files or objects to transfer.
	//
	// For more information, see [Specifying what DataSync transfers by using a manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-manifestconfig
	//
	ManifestConfig interface{} `field:"optional" json:"manifestConfig" yaml:"manifestConfig"`
	// The name of a task.
	//
	// This value is a text reference that is used to identify the task in the console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the configuration options for a task. Some options include preserving file or object metadata and verifying data integrity.
	//
	// You can also override these options before starting an individual run of a task (also known as a *task execution* ). For more information, see [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Specifies a schedule used to periodically transfer files from a source to a destination location.
	//
	// The schedule should be specified in UTC time. For more information, see [Scheduling your task](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Specifies the tags that you want to apply to the Amazon Resource Name (ARN) representing the task.
	//
	// *Tags* are key-value pairs that help you manage, filter, and search for your DataSync resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies how you want to configure a task report, which provides detailed information about your DataSync transfer.
	//
	// For more information, see [Monitoring your DataSync transfers with task reports](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html) .
	//
	// When using this parameter, your caller identity (the role that you're using DataSync with) must have the `iam:PassRole` permission. The [AWSDataSyncFullAccess](https://docs.aws.amazon.com/datasync/latest/userguide/security-iam-awsmanpol.html#security-iam-awsmanpol-awsdatasyncfullaccess) policy includes this permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-taskreportconfig
	//
	TaskReportConfig interface{} `field:"optional" json:"taskReportConfig" yaml:"taskReportConfig"`
}

