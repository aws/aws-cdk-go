package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTaskProps := &cfnTaskProps{
//   	destinationLocationArn: jsii.String("destinationLocationArn"),
//   	sourceLocationArn: jsii.String("sourceLocationArn"),
//
//   	// the properties below are optional
//   	cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	excludes: []interface{}{
//   		&filterRuleProperty{
//   			filterType: jsii.String("filterType"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	includes: []interface{}{
//   		&filterRuleProperty{
//   			filterType: jsii.String("filterType"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	options: &optionsProperty{
//   		atime: jsii.String("atime"),
//   		bytesPerSecond: jsii.Number(123),
//   		gid: jsii.String("gid"),
//   		logLevel: jsii.String("logLevel"),
//   		mtime: jsii.String("mtime"),
//   		objectTags: jsii.String("objectTags"),
//   		overwriteMode: jsii.String("overwriteMode"),
//   		posixPermissions: jsii.String("posixPermissions"),
//   		preserveDeletedFiles: jsii.String("preserveDeletedFiles"),
//   		preserveDevices: jsii.String("preserveDevices"),
//   		securityDescriptorCopyFlags: jsii.String("securityDescriptorCopyFlags"),
//   		taskQueueing: jsii.String("taskQueueing"),
//   		transferMode: jsii.String("transferMode"),
//   		uid: jsii.String("uid"),
//   		verifyMode: jsii.String("verifyMode"),
//   	},
//   	schedule: &taskScheduleProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTaskProps struct {
	// The Amazon Resource Name (ARN) of an AWS storage resource's location.
	DestinationLocationArn *string `field:"required" json:"destinationLocationArn" yaml:"destinationLocationArn"`
	// The Amazon Resource Name (ARN) of the source location for the task.
	SourceLocationArn *string `field:"required" json:"sourceLocationArn" yaml:"sourceLocationArn"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that is used to monitor and log events in the task.
	//
	// For more information about how to use CloudWatch Logs with DataSync, see [Monitoring Your Task](https://docs.aws.amazon.com/datasync/latest/userguide/monitor-datasync.html#cloudwatchlogs) in the *AWS DataSync User Guide.*
	//
	// For more information about these groups, see [Working with Log Groups and Log Streams](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html) in the *Amazon CloudWatch Logs User Guide* .
	CloudWatchLogGroupArn *string `field:"optional" json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// A list of filter rules that determines which files to exclude from a task.
	//
	// The list should contain a single filter string that consists of the patterns to exclude. The patterns are delimited by "|" (that is, a pipe), for example, `"/folder1|/folder2"` .
	Excludes interface{} `field:"optional" json:"excludes" yaml:"excludes"`
	// A list of filter rules that determines which files to include when running a task.
	//
	// The pattern contains a single filter string that consists of the patterns to include. The patterns are delimited by "|" (that is, a pipe), for example, `"/folder1|/folder2"` .
	Includes interface{} `field:"optional" json:"includes" yaml:"includes"`
	// The name of a task.
	//
	// This value is a text reference that is used to identify the task in the console.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The set of configuration options that control the behavior of a single execution of the task that occurs when you call `StartTaskExecution` .
	//
	// You can configure these options to preserve metadata such as user ID (UID) and group ID (GID), file permissions, data integrity verification, and so on.
	//
	// For each individual task execution, you can override these options by specifying the `OverrideOptions` before starting the task execution. For more information, see the [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) operation.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Specifies a schedule used to periodically transfer files from a source to a destination location.
	//
	// The schedule should be specified in UTC time. For more information, see [Scheduling your task](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html) .
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The key-value pair that represents the tag that you want to add to the resource.
	//
	// The value can be an empty string.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

