package awsapplicationautoscaling


// `ScheduledAction` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html) resource that specifies a scheduled action for a scalable target.
//
// For more information, see [PutScheduledAction](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScheduledAction.html) in the *Application Auto Scaling API Reference* . For more information about scheduled scaling, including the format for cron expressions, see [Scheduled scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledActionProperty := &scheduledActionProperty{
//   	schedule: jsii.String("schedule"),
//   	scheduledActionName: jsii.String("scheduledActionName"),
//
//   	// the properties below are optional
//   	endTime: NewDate(),
//   	scalableTargetAction: &scalableTargetActionProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	startTime: NewDate(),
//   	timezone: jsii.String("timezone"),
//   }
//
type CfnScalableTarget_ScheduledActionProperty struct {
	// The schedule for this action. The following formats are supported:.
	//
	// - At expressions - " `at( *yyyy* - *mm* - *dd* T *hh* : *mm* : *ss* )` "
	// - Rate expressions - " `rate( *value* *unit* )` "
	// - Cron expressions - " `cron( *fields* )` "
	//
	// At expressions are useful for one-time schedules. Cron expressions are useful for scheduled actions that run periodically at a specified date and time, and rate expressions are useful for scheduled actions that run at a regular interval.
	//
	// At and cron expressions use Universal Coordinated Time (UTC) by default.
	//
	// The cron format consists of six fields separated by white spaces: [Minutes] [Hours] [Day_of_Month] [Month] [Day_of_Week] [Year].
	//
	// For rate expressions, *value* is a positive integer and *unit* is `minute` | `minutes` | `hour` | `hours` | `day` | `days` .
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// The name of the scheduled action.
	//
	// This name must be unique among all other scheduled actions on the specified scalable target.
	ScheduledActionName *string `field:"required" json:"scheduledActionName" yaml:"scheduledActionName"`
	// The date and time that the action is scheduled to end, in UTC.
	EndTime interface{} `field:"optional" json:"endTime" yaml:"endTime"`
	// The new minimum and maximum capacity.
	//
	// You can set both values or just one. At the scheduled time, if the current capacity is below the minimum capacity, Application Auto Scaling scales out to the minimum capacity. If the current capacity is above the maximum capacity, Application Auto Scaling scales in to the maximum capacity.
	ScalableTargetAction interface{} `field:"optional" json:"scalableTargetAction" yaml:"scalableTargetAction"`
	// The date and time that the action is scheduled to begin, in UTC.
	StartTime interface{} `field:"optional" json:"startTime" yaml:"startTime"`
	// The time zone used when referring to the date and time of a scheduled action, when the scheduled action uses an at or cron expression.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

