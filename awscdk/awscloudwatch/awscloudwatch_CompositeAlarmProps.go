package awscloudwatch


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
// Experimental.
type CompositeAlarmProps struct {
	// Expression that specifies which other alarms are to be evaluated to determine this composite alarm's state.
	// Experimental.
	AlarmRule IAlarmRule `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// Whether the actions for this alarm are enabled.
	// Experimental.
	ActionsEnabled *bool `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Description for the alarm.
	// Experimental.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Name of the alarm.
	// Experimental.
	CompositeAlarmName *string `field:"optional" json:"compositeAlarmName" yaml:"compositeAlarmName"`
}

