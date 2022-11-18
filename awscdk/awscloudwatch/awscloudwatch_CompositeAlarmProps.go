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
//   alarmRule := cloudwatch.alarmRule.anyOf(cloudwatch.alarmRule.allOf(cloudwatch.alarmRule.anyOf(alarm1, cloudwatch.alarmRule.fromAlarm(alarm2, cloudwatch.alarmState_OK), alarm3), cloudwatch.alarmRule.not(cloudwatch.alarmRule.fromAlarm(alarm4, cloudwatch.alarmState_INSUFFICIENT_DATA))), cloudwatch.alarmRule.fromBoolean(jsii.Boolean(false)))
//
//   cloudwatch.NewCompositeAlarm(this, jsii.String("MyAwesomeCompositeAlarm"), &compositeAlarmProps{
//   	alarmRule: alarmRule,
//   })
//
type CompositeAlarmProps struct {
	// Expression that specifies which other alarms are to be evaluated to determine this composite alarm's state.
	AlarmRule IAlarmRule `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// Whether the actions for this alarm are enabled.
	ActionsEnabled *bool `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Actions will be suppressed if the suppressor alarm is in the ALARM state.
	ActionsSuppressor IAlarm `field:"optional" json:"actionsSuppressor" yaml:"actionsSuppressor"`
	// The maximum duration that the composite alarm waits after suppressor alarm goes out of the ALARM state.
	//
	// After this time, the composite alarm performs its actions.
	ActionsSuppressorExtensionPeriod awscdk.Duration `field:"optional" json:"actionsSuppressorExtensionPeriod" yaml:"actionsSuppressorExtensionPeriod"`
	// The maximum duration that the composite alarm waits for the suppressor alarm to go into the ALARM state.
	//
	// After this time, the composite alarm performs its actions.
	ActionsSuppressorWaitPeriod awscdk.Duration `field:"optional" json:"actionsSuppressorWaitPeriod" yaml:"actionsSuppressorWaitPeriod"`
	// Description for the alarm.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Name of the alarm.
	CompositeAlarmName *string `field:"optional" json:"compositeAlarmName" yaml:"compositeAlarmName"`
}

