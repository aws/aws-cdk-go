package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Schedule expression for CloudWatch alarm mute rule.
//
// You can choose from three schedule types when configuring your schedule: cron-based and one-time schedules.
// Cron-based schedule is recurring schedule.
//
// Example:
//   var alarm1 Alarm
//   var alarm2 Alarm
//
//
//   alarmMuteRule := cloudwatch.NewAlarmMuteRule(this, jsii.String("AlarmMuteRule"), &AlarmMuteRuleProps{
//   	Alarms: []IAlarmRef{
//   		alarm1,
//   	},
//   	// Defines the mute period begins at 0:00 everyday in UTC
//   	Schedule: cloudwatch.ScheduleExpression_Cron(&CronOptions{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("0"),
//   	}),
//   	// Specifies the mute rule lasts 1 hour.
//   	Duration: awscdk.Duration_Hours(jsii.Number(1)),
//   })
//
//   // The mute target can be added after construction.
//   alarmMuteRule.AddAlarm(alarm2)
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/alarm-mute-rules.html#defining-alarm-mute-rules
//
type ScheduleExpression interface {
	// Retrieve the expression for this schedule.
	ExpressionString() *string
	// Retrieve the expression for this schedule.
	TimeZone() awscdk.TimeZone
}

// The jsii proxy struct for ScheduleExpression
type jsiiProxy_ScheduleExpression struct {
	_ byte // padding
}

func (j *jsiiProxy_ScheduleExpression) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScheduleExpression) TimeZone() awscdk.TimeZone {
	var returns awscdk.TimeZone
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}


func NewScheduleExpression_Override(s ScheduleExpression) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.ScheduleExpression",
		nil, // no parameters
		s,
	)
}

// Construct a one-time schedule from a date.
func ScheduleExpression_At(date *CalendarDateTime, timeZone awscdk.TimeZone) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_AtParameters(date); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.ScheduleExpression",
		"at",
		[]interface{}{date, timeZone},
		&returns,
	)

	return returns
}

// Create a recurring schedule from a set of cron fields and time zone.
func ScheduleExpression_Cron(options *CronOptions) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_CronParameters(options); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.ScheduleExpression",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func ScheduleExpression_Expression(expression *string, timeZone awscdk.TimeZone) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.ScheduleExpression",
		"expression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

