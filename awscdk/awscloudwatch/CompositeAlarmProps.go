package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for creating a Composite Alarm.
//
// Example:
//   var alarm1 alarm
//   var alarm2 alarm
//   var alarm3 alarm
//   var alarm4 alarm
//
//
//   alarmRule := cloudwatch.AlarmRule_AnyOf(cloudwatch.AlarmRule_AllOf(cloudwatch.AlarmRule_AnyOf(alarm1, cloudwatch.AlarmRule_FromAlarm(alarm2, cloudwatch.AlarmState_OK), alarm3), cloudwatch.AlarmRule_Not(cloudwatch.AlarmRule_FromAlarm(alarm4, cloudwatch.AlarmState_INSUFFICIENT_DATA))), cloudwatch.AlarmRule_FromBoolean(jsii.Boolean(false)))
//
//   cloudwatch.NewCompositeAlarm(this, jsii.String("MyAwesomeCompositeAlarm"), &CompositeAlarmProps{
//   	AlarmRule: AlarmRule,
//   })
//
type CompositeAlarmProps struct {
	// Expression that specifies which other alarms are to be evaluated to determine this composite alarm's state.
	AlarmRule IAlarmRule `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// Whether the actions for this alarm are enabled.
	// Default: true.
	//
	ActionsEnabled *bool `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Actions will be suppressed if the suppressor alarm is in the ALARM state.
	// Default: - alarm will not be suppressed.
	//
	ActionsSuppressor IAlarm `field:"optional" json:"actionsSuppressor" yaml:"actionsSuppressor"`
	// The maximum duration that the composite alarm waits after suppressor alarm goes out of the ALARM state.
	//
	// After this time, the composite alarm performs its actions.
	// Default: - 1 minute extension period will be set.
	//
	ActionsSuppressorExtensionPeriod awscdk.Duration `field:"optional" json:"actionsSuppressorExtensionPeriod" yaml:"actionsSuppressorExtensionPeriod"`
	// The maximum duration that the composite alarm waits for the suppressor alarm to go into the ALARM state.
	//
	// After this time, the composite alarm performs its actions.
	// Default: - 1 minute wait period will be set.
	//
	ActionsSuppressorWaitPeriod awscdk.Duration `field:"optional" json:"actionsSuppressorWaitPeriod" yaml:"actionsSuppressorWaitPeriod"`
	// Description for the alarm.
	// Default: - No description.
	//
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Name of the alarm.
	// Default: - Automatically generated name.
	//
	CompositeAlarmName *string `field:"optional" json:"compositeAlarmName" yaml:"compositeAlarmName"`
}

