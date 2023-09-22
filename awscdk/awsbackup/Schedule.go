package awsbackup

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
)

// Example:
//   var plan backupPlan
//
//   plan.AddRule(backup.NewBackupPlanRule(&BackupPlanRuleProps{
//   	CompletionWindow: awscdk.Duration_Hours(jsii.Number(2)),
//   	StartWindow: awscdk.Duration_*Hours(jsii.Number(1)),
//   	Schedule: backup.Schedule_Cron(&CronOptions{
//   		 // Only cron expressions are supported
//   		Day: jsii.String("15"),
//   		Hour: jsii.String("3"),
//   		Minute: jsii.String("30"),
//   	}),
//   	MoveToColdStorageAfter: awscdk.Duration_Days(jsii.Number(30)),
//   }))
//
type Schedule interface {
	awscdk.Schedule
	// Retrieve the expression for this schedule.
	ExpressionString() *string
	// The timezone of the expression, if applicable.
	TimeZone() awscdk.TimeZone
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	internal.Type__awscdkSchedule
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

func (j *jsiiProxy_Schedule) TimeZone() awscdk.TimeZone {
	var returns awscdk.TimeZone
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
		"aws-cdk-lib.aws_backup.Schedule",
		nil, // no parameters
		s,
	)
}

// Construct a schedule from cron options.
func Schedule_Cron(options *awscdk.CronOptions) Schedule {
	_init_.Initialize()

	if err := validateSchedule_CronParameters(options); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_Expression(expression *string, timeZone awscdk.TimeZone) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.Schedule",
		"expression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

// Construct a one-time schedule from a date.
func Schedule_ProtectedAt(date *time.Time, timeZone awscdk.TimeZone) awscdk.Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedAtParameters(date); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.Schedule",
		"protectedAt",
		[]interface{}{date, timeZone},
		&returns,
	)

	return returns
}

// Create a schedule from a set of cron fields.
func Schedule_ProtectedCron(options *awscdk.CronOptions, module *string) awscdk.Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedCronParameters(options); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.Schedule",
		"protectedCron",
		[]interface{}{options, module},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_ProtectedExpression(expression *string, timeZone awscdk.TimeZone) awscdk.Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.Schedule",
		"protectedExpression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
//
// Rates may be defined with any unit of time, but when converted into minutes, the duration must be a positive whole number of minutes.
func Schedule_ProtectedRate(duration awscdk.Duration) awscdk.Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedRateParameters(duration); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.Schedule",
		"protectedRate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

