package awscdkscheduleralpha

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2/internal"
)

// ScheduleExpression for EventBridge Schedule.
//
// You can choose from three schedule types when configuring your schedule: rate-based, cron-based, and one-time schedules.
// Both rate-based and cron-based schedules are recurring schedules.
//
// Example:
//   var fn function
//
//
//   target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
//   	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
//   		"payload": jsii.String("useful"),
//   	}),
//   })
//
//   schedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
//   	Target: Target,
//   	Description: jsii.String("This is a test schedule that invokes lambda function every 10 minutes."),
//   })
//
// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html
//
// Experimental.
type ScheduleExpression interface {
	awscdk.Schedule
	// Retrieve the expression for this schedule.
	// Experimental.
	ExpressionString() *string
	// The timezone of the expression, if applicable.
	// Experimental.
	TimeZone() awscdk.TimeZone
}

// The jsii proxy struct for ScheduleExpression
type jsiiProxy_ScheduleExpression struct {
	internal.Type__awscdkSchedule
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


// Experimental.
func NewScheduleExpression_Override(s ScheduleExpression) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		nil, // no parameters
		s,
	)
}

// Construct a one-time schedule from a date.
// Experimental.
func ScheduleExpression_At(date *time.Time, timeZone awscdk.TimeZone) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_AtParameters(date); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		"at",
		[]interface{}{date, timeZone},
		&returns,
	)

	return returns
}

// Create a recurring schedule from a set of cron fields and time zone.
// Experimental.
func ScheduleExpression_Cron(options *awscdk.CronOptions) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_CronParameters(options); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
// Experimental.
func ScheduleExpression_Expression(expression *string, timeZone awscdk.TimeZone) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		"expression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

// Construct a one-time schedule from a date.
// Experimental.
func ScheduleExpression_ProtectedAt(date *time.Time, timeZone awscdk.TimeZone) awscdk.Schedule {
	_init_.Initialize()

	if err := validateScheduleExpression_ProtectedAtParameters(date); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		"protectedAt",
		[]interface{}{date, timeZone},
		&returns,
	)

	return returns
}

// Create a schedule from a set of cron fields.
// Experimental.
func ScheduleExpression_ProtectedCron(options *awscdk.CronOptions, module *string) awscdk.Schedule {
	_init_.Initialize()

	if err := validateScheduleExpression_ProtectedCronParameters(options); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		"protectedCron",
		[]interface{}{options, module},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
// Experimental.
func ScheduleExpression_ProtectedExpression(expression *string, timeZone awscdk.TimeZone) awscdk.Schedule {
	_init_.Initialize()

	if err := validateScheduleExpression_ProtectedExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		"protectedExpression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
//
// Rates may be defined with any unit of time, but when converted into minutes, the duration must be a positive whole number of minutes.
// Experimental.
func ScheduleExpression_ProtectedRate(duration awscdk.Duration) awscdk.Schedule {
	_init_.Initialize()

	if err := validateScheduleExpression_ProtectedRateParameters(duration); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		"protectedRate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Construct a recurring schedule from an interval and a time unit.
//
// Rates may be defined with any unit of time, but when converted into minutes, the duration must be a positive whole number of minutes.
// Experimental.
func ScheduleExpression_Rate(duration awscdk.Duration) ScheduleExpression {
	_init_.Initialize()

	if err := validateScheduleExpression_RateParameters(duration); err != nil {
		panic(err)
	}
	var returns ScheduleExpression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ScheduleExpression",
		"rate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

