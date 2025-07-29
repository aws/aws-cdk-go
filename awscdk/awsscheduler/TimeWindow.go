package awsscheduler

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A time window during which EventBridge Scheduler invokes the schedule.
//
// Example:
//   var target lambdaInvoke
//
//
//   schedule := awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(12))),
//   	Target: Target,
//   	TimeWindow: awscdk.TimeWindow_Flexible(awscdk.Duration_*Hours(jsii.Number(10))),
//   })
//
type TimeWindow interface {
	// The maximum time window during which the schedule can be invoked.
	//
	// Must be between 1 to 1440 minutes.
	// Default: - no value.
	//
	MaxWindow() awscdk.Duration
	// Determines whether the schedule is invoked within a flexible time window.
	Mode() *string
}

// The jsii proxy struct for TimeWindow
type jsiiProxy_TimeWindow struct {
	_ byte // padding
}

func (j *jsiiProxy_TimeWindow) MaxWindow() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"maxWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TimeWindow) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}


// TimeWindow is enabled.
func TimeWindow_Flexible(maxWindow awscdk.Duration) TimeWindow {
	_init_.Initialize()

	if err := validateTimeWindow_FlexibleParameters(maxWindow); err != nil {
		panic(err)
	}
	var returns TimeWindow

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_scheduler.TimeWindow",
		"flexible",
		[]interface{}{maxWindow},
		&returns,
	)

	return returns
}

// TimeWindow is disabled.
func TimeWindow_Off() TimeWindow {
	_init_.Initialize()

	var returns TimeWindow

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_scheduler.TimeWindow",
		"off",
		nil, // no parameters
		&returns,
	)

	return returns
}

