package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Threshold configuration for the AT_LEAST composite alarm rule expression.
//
// Use `AtLeastThreshold.count()` for an absolute number or
// `AtLeastThreshold.percentage()` for a percentage-based threshold.
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
type AtLeastThreshold interface {
}

// The jsii proxy struct for AtLeastThreshold
type jsiiProxy_AtLeastThreshold struct {
	_ byte // padding
}

func NewAtLeastThreshold_Override(a AtLeastThreshold) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.AtLeastThreshold",
		nil, // no parameters
		a,
	)
}

// Creates a count-based threshold for the AT_LEAST expression.
//
// The count must be a positive integer between 1 and the number of operands.
func AtLeastThreshold_Count(count *float64) AtLeastThreshold {
	_init_.Initialize()

	if err := validateAtLeastThreshold_CountParameters(count); err != nil {
		panic(err)
	}
	var returns AtLeastThreshold

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.AtLeastThreshold",
		"count",
		[]interface{}{count},
		&returns,
	)

	return returns
}

// Creates a percentage-based threshold for the AT_LEAST expression.
//
// The percentage must be an integer between 1 and 100.
func AtLeastThreshold_Percentage(percentage *float64) AtLeastThreshold {
	_init_.Initialize()

	if err := validateAtLeastThreshold_PercentageParameters(percentage); err != nil {
		panic(err)
	}
	var returns AtLeastThreshold

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.AtLeastThreshold",
		"percentage",
		[]interface{}{percentage},
		&returns,
	)

	return returns
}

