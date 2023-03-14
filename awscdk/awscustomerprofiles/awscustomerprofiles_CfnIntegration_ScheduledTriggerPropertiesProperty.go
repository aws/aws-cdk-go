package awscustomerprofiles


// Specifies the configuration details of a scheduled-trigger flow that you define.
//
// Currently, these settings only apply to the scheduled-trigger type.
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
//   	scheduleEndTime: jsii.Number(123),
//   	scheduleOffset: jsii.Number(123),
//   	scheduleStartTime: jsii.Number(123),
//   	timezone: jsii.String("timezone"),
//   }
//
type CfnIntegration_ScheduledTriggerPropertiesProperty struct {
	// The scheduling expression that determines the rate at which the schedule will run, for example rate (5 minutes).
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Specifies whether a scheduled flow has an incremental data transfer or a complete data transfer for each flow run.
	DataPullMode *string `field:"optional" json:"dataPullMode" yaml:"dataPullMode"`
	// Specifies the date range for the records to import from the connector in the first flow run.
	FirstExecutionFrom *float64 `field:"optional" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// Specifies the scheduled end time for a scheduled-trigger flow.
	ScheduleEndTime *float64 `field:"optional" json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// Specifies the optional offset that is added to the time interval for a schedule-triggered flow.
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// Specifies the scheduled start time for a scheduled-trigger flow.
	ScheduleStartTime *float64 `field:"optional" json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// Specifies the time zone used when referring to the date and time of a scheduled-triggered flow, such as America/New_York.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

