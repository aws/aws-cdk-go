package awscloudwatch


// Options for the AT_LEAST AlarmRule wrapper function.
//
// Example:
//   var alarm1 Alarm
//   var alarm2 Alarm
//   var alarm3 Alarm
//   var alarm4 Alarm
//
//
//   alarmRule := cloudwatch.AlarmRule_AnyOf(cloudwatch.AlarmRule_AllOf(cloudwatch.AlarmRule_AnyOf(alarm1, cloudwatch.AlarmRule_FromAlarm(alarm2, cloudwatch.AlarmState_OK), alarm3), cloudwatch.AlarmRule_Not(cloudwatch.AlarmRule_FromAlarm(alarm4, cloudwatch.AlarmState_INSUFFICIENT_DATA)), cloudwatch.AlarmRule_AtLeast(cloudwatch.AlarmState_ALARM, &AtLeastOptions{
//   	Operands: []IAlarm{
//   		alarm1,
//   		alarm2,
//   		alarm3,
//   	},
//   	Threshold: cloudwatch.AtLeastThreshold_Count(jsii.Number(2)),
//   }), cloudwatch.AlarmRule_AtLeastNot(cloudwatch.AlarmState_OK, &AtLeastOptions{
//   	Operands: []IAlarm{
//   		alarm1,
//   		alarm2,
//   		alarm3,
//   	},
//   	Threshold: cloudwatch.AtLeastThreshold_Percentage(jsii.Number(60)),
//   })), cloudwatch.AlarmRule_FromBoolean(jsii.Boolean(false)))
//
//   cloudwatch.NewCompositeAlarm(this, jsii.String("MyAwesomeCompositeAlarm"), &CompositeAlarmProps{
//   	AlarmRule: AlarmRule,
//   })
//
type AtLeastOptions struct {
	// Alarms to evaluate in the AT_LEAST expression.
	//
	// Must contain at least one alarm.
	Operands *[]IAlarm `field:"required" json:"operands" yaml:"operands"`
	// Threshold for the AT_LEAST expression.
	//
	// Use `AtLeastThreshold.count()` for an absolute number
	// or `AtLeastThreshold.percentage()` for a percentage.
	Threshold AtLeastThreshold `field:"required" json:"threshold" yaml:"threshold"`
}

