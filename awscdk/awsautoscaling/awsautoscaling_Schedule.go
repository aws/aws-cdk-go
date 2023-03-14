package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Schedule for scheduled scaling actions.
//
// Example:
//   var autoScalingGroup autoScalingGroup
//
//
//   autoScalingGroup.scaleOnSchedule(jsii.String("PrescaleInTheMorning"), &basicScheduledActionProps{
//   	schedule: autoscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(20),
//   })
//
//   autoScalingGroup.scaleOnSchedule(jsii.String("AllowDownscalingAtNight"), &basicScheduledActionProps{
//   	schedule: autoscaling.*schedule.cron(&cronOptions{
//   		hour: jsii.String("20"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(1),
//   })
//
// Experimental.
type Schedule interface {
	// Retrieve the expression for this schedule.
	// Experimental.
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


// Experimental.
func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling.Schedule",
		nil, // no parameters
		s,
	)
}

// Create a schedule from a set of cron fields.
// Experimental.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	if err := validateSchedule_CronParameters(options); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
// See: http://crontab.org/
//
// Experimental.
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

