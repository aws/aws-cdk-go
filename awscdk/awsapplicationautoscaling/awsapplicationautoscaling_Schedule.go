package awsapplicationautoscaling

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Schedule for scheduled scaling actions.
//
// Example:
//   var resource someScalableResource
//
//
//   capacity := resource.autoScaleCapacity(&caps{
//   	minCapacity: jsii.Number(1),
//   	maxCapacity: jsii.Number(50),
//   })
//
//   capacity.scaleOnSchedule(jsii.String("PrescaleInTheMorning"), &scalingSchedule{
//   	schedule: appscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(20),
//   })
//
//   capacity.scaleOnSchedule(jsii.String("AllowDownscalingAtNight"), &scalingSchedule{
//   	schedule: appscaling.*schedule.cron(&cronOptions{
//   		hour: jsii.String("20"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(1),
//   })
//
type Schedule interface {
	// Retrieve the expression for this schedule.
	ExpressionString() *string
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


func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		nil, // no parameters
		s,
	)
}

// Construct a Schedule from a moment in time.
func Schedule_At(moment *time.Time) Schedule {
	_init_.Initialize()

	if err := validateSchedule_AtParameters(moment); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"at",
		[]interface{}{moment},
		&returns,
	)

	return returns
}

// Create a schedule from a set of cron fields.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	if err := validateSchedule_CronParameters(options); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
func Schedule_Rate(duration awscdk.Duration) Schedule {
	_init_.Initialize()

	if err := validateSchedule_RateParameters(duration); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"rate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

