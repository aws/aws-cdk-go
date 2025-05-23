package awsscheduler

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// ScheduleExpression for EventBridge Schedule.
//
// You can choose from three schedule types when configuring your schedule: rate-based, cron-based, and one-time schedules.
// Both rate-based and cron-based schedules are recurring schedules.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   var deliveryStream iDeliveryStream
//
//
//   payload := map[string]*string{
//   	"Data": jsii.String("record"),
//   }
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewFirehosePutRecord(deliveryStream, &ScheduleTargetBaseProps{
//   		Input: awscdk.ScheduleTargetInput_FromObject(payload),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html
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
		"aws-cdk-lib.aws_scheduler.ScheduleExpression",
		nil, // no parameters
		s,
	)
}

// Construct a one-time schedule from a date.
func ScheduleExpression_At(date *time.Time, timeZone awscdk.TimeZone) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_AtParameters(date); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_scheduler.ScheduleExpression",
		"at",
		[]interface{}{date, timeZone},
		&returns,
	)

	return returns
}

// Create a recurring schedule from a set of cron fields and time zone.
func ScheduleExpression_Cron(options *CronOptionsWithTimezone) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_CronParameters(options); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_scheduler.ScheduleExpression",
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
		"aws-cdk-lib.aws_scheduler.ScheduleExpression",
		"expression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

// Construct a recurring schedule from an interval and a time unit.
//
// Rates may be defined with any unit of time, but when converted into minutes, the duration must be a positive whole number of minutes.
func ScheduleExpression_Rate(duration awscdk.Duration) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_RateParameters(duration); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_scheduler.ScheduleExpression",
		"rate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

