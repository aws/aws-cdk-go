package awsappflow


// Specifies the configuration details of a schedule-triggered flow as defined by the user.
//
// Currently, these settings only apply to the `Scheduled` trigger type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledTriggerPropertiesProperty := &scheduledTriggerPropertiesProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//
//   	// the properties below are optional
//   	dataPullMode: jsii.String("dataPullMode"),
//   	firstExecutionFrom: jsii.Number(123),
//   	flowErrorDeactivationThreshold: jsii.Number(123),
//   	scheduleEndTime: jsii.Number(123),
//   	scheduleOffset: jsii.Number(123),
//   	scheduleStartTime: jsii.Number(123),
//   	timeZone: jsii.String("timeZone"),
//   }
//
type CfnFlow_ScheduledTriggerPropertiesProperty struct {
	// The scheduling expression that determines the rate at which the schedule will run, for example `rate(5minutes)` .
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Specifies whether a scheduled flow has an incremental data transfer or a complete data transfer for each flow run.
	DataPullMode *string `field:"optional" json:"dataPullMode" yaml:"dataPullMode"`
	// `CfnFlow.ScheduledTriggerPropertiesProperty.FirstExecutionFrom`.
	FirstExecutionFrom *float64 `field:"optional" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// `CfnFlow.ScheduledTriggerPropertiesProperty.FlowErrorDeactivationThreshold`.
	FlowErrorDeactivationThreshold *float64 `field:"optional" json:"flowErrorDeactivationThreshold" yaml:"flowErrorDeactivationThreshold"`
	// The time at which the scheduled flow ends.
	//
	// The time is formatted as a timestamp that follows the ISO 8601 standard, such as `2022-04-27T13:00:00-07:00` .
	ScheduleEndTime *float64 `field:"optional" json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// Specifies the optional offset that is added to the time interval for a schedule-triggered flow.
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// The time at which the scheduled flow starts.
	//
	// The time is formatted as a timestamp that follows the ISO 8601 standard, such as `2022-04-26T13:00:00-07:00` .
	ScheduleStartTime *float64 `field:"optional" json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// Specifies the time zone used when referring to the dates and times of a scheduled flow, such as `America/New_York` .
	//
	// This time zone is only a descriptive label. It doesn't affect how Amazon AppFlow interprets the timestamps that you specify to schedule the flow.
	//
	// If you want to schedule a flow by using times in a particular time zone, indicate the time zone as a UTC offset in your timestamps. For example, the UTC offsets for the `America/New_York` timezone are `-04:00` EDT and `-05:00 EST` .
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

