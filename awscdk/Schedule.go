package awscdk

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Core Schedule.
//
// This construct is not meant to be used as is or exposed to consumers in other modules.
// It is meant to be extended by other modules that require some sort of schedule implementation. All
// methods in `core.Schedule` are protected, so that construct authors can decide which APIs to expose.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html
//
type Schedule interface {
	// Retrieve the expression for this schedule.
	ExpressionString() *string
	// The timezone of the expression, if applicable.
	TimeZone() TimeZone
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	_ byte // padding
}

func (j *jsiiProxy_Schedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Schedule) TimeZone() TimeZone {
	var returns TimeZone
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}


func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Schedule",
		nil, // no parameters
		s,
	)
}

// Construct a one-time schedule from a date.
func Schedule_ProtectedAt(date *time.Time, timeZone TimeZone) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedAtParameters(date); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Schedule",
		"protectedAt",
		[]interface{}{date, timeZone},
		&returns,
	)

	return returns
}

// Create a schedule from a set of cron fields.
func Schedule_ProtectedCron(options *CronOptions, module *string) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedCronParameters(options); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Schedule",
		"protectedCron",
		[]interface{}{options, module},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_ProtectedExpression(expression *string, timeZone TimeZone) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Schedule",
		"protectedExpression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
//
// Rates may be defined with any unit of time, but when converted into minutes, the duration must be a positive whole number of minutes.
func Schedule_ProtectedRate(duration Duration) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedRateParameters(duration); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Schedule",
		"protectedRate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

