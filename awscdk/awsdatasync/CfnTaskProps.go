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
//   	TaskMode: jsii.String("taskMode"),
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
	// Specifies the ARN of your transfer's source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-sourcelocationarn
	//
	SourceLocationArn *string `field:"required" json:"sourceLocationArn" yaml:"sourceLocationArn"`
	// Specifies the Amazon Resource Name (ARN) of an Amazon CloudWatch log group for monitoring your task.
	//
	// For Enhanced mode tasks, you don't need to specify anything. DataSync automatically sends logs to a CloudWatch log group named `/aws/datasync` .
	//
	// For more information, see [Monitoring data transfers with CloudWatch Logs](https://docs.aws.amazon.com/datasync/latest/userguide/configure-logging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-cloudwatchloggrouparn
	//
	CloudWatchLogGroupArn *string `field:"optional" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// Specifies exclude filters that define the files, objects, and folders in your source location that you don't want DataSync to transfer.
	//
	// For more information and examples, see [Specifying what DataSync transfers by using filters](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-excludes
	//
	Excludes interface{} `field:"optional" json:"excludes" yaml:"excludes"`
	// Specifies include filters that define the files, objects, and folders in your source location that you want DataSync to transfer.
	//
	// For more information and examples, see [Specifying what DataSync transfers by using filters](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-includes
	//
	Includes interface{} `field:"optional" json:"includes" yaml:"includes"`
	// The configuration of the manifest that lists the files or objects that you want DataSync to transfer.
	//
	// For more information, see [Specifying what DataSync transfers by using a manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-manifestconfig
	//
	ManifestConfig interface{} `field:"optional" json:"manifestConfig" yaml:"manifestConfig"`
	// Specifies the name of your task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies your task's settings, such as preserving file metadata, verifying data integrity, among other options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Specifies a schedule for when you want your task to run.
	//
	// For more information, see [Scheduling your task](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Specifies the tags that you want to apply to your task.
	//
	// *Tags* are key-value pairs that help you manage, filter, and search for your DataSync resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies one of the following task modes for your data transfer:.
	//
	// - `ENHANCED` - Transfer virtually unlimited numbers of objects with higher performance than Basic mode. Enhanced mode tasks optimize the data transfer process by listing, preparing, transferring, and verifying data in parallel. Enhanced mode is currently available for transfers between Amazon S3 locations.
	//
	// > To create an Enhanced mode task, the IAM role that you use to call the `CreateTask` operation must have the `iam:CreateServiceLinkedRole` permission.
	// - `BASIC` (default) - Transfer files or objects between AWS storage and all other supported DataSync locations. Basic mode tasks are subject to [quotas](https://docs.aws.amazon.com/datasync/latest/userguide/datasync-limits.html) on the number of files, objects, and directories in a dataset. Basic mode sequentially prepares, transfers, and verifies data, making it slower than Enhanced mode for most workloads.
	//
	// For more information, see [Understanding task mode differences](https://docs.aws.amazon.com/datasync/latest/userguide/choosing-task-mode.html#task-mode-differences) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-taskmode
	//
	TaskMode *string `field:"optional" json:"taskMode" yaml:"taskMode"`
	// Specifies how you want to configure a task report, which provides detailed information about your DataSync transfer.
	//
	// For more information, see [Monitoring your DataSync transfers with task reports](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html) .
	//
	// When using this parameter, your caller identity (the role that you're using DataSync with) must have the `iam:PassRole` permission. The [AWSDataSyncFullAccess](https://docs.aws.amazon.com/datasync/latest/userguide/security-iam-awsmanpol.html#security-iam-awsmanpol-awsdatasyncfullaccess) policy includes this permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-taskreportconfig
	//
	TaskReportConfig interface{} `field:"optional" json:"taskReportConfig" yaml:"taskReportConfig"`
}

