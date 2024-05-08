package awsdatasync


// Configures your AWS DataSync task to run on a [schedule](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html) (at a minimum interval of 1 hour).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskScheduleProperty := &TaskScheduleProperty{
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskschedule.html
//
type CfnTask_TaskScheduleProperty struct {
	// Specifies your task schedule by using a cron expression in UTC time.
	//
	// For information about cron expression syntax, see the [*Amazon EventBridge User Guide*](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-cron-expressions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskschedule.html#cfn-datasync-task-taskschedule-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Specifies whether to enable or disable your task schedule.
	//
	// Your schedule is enabled by default, but there can be situations where you need to disable it. For example, you might need to perform maintenance on a storage system before you can begin a recurring DataSync transfer.
	//
	// DataSync might disable your schedule automatically if your task fails repeatedly with the same error. For more information, see the [*DataSync User Guide*](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html#pause-task-schedule) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-taskschedule.html#cfn-datasync-task-taskschedule-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

