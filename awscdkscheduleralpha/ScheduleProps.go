package awscdkscheduleralpha


// Construction properties for `Schedule`.
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
// Experimental.
type ScheduleProps struct {
	// The expression that defines when the schedule runs.
	//
	// Can be either a `at`, `rate`
	// or `cron` expression.
	// Experimental.
	Schedule ScheduleExpression `field:"required" json:"schedule" yaml:"schedule"`
	// The schedule's target details.
	// Experimental.
	Target IScheduleTarget `field:"required" json:"target" yaml:"target"`
	// The description you specify for the schedule.
	// Default: - no value.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The schedule's group.
	// Experimental.
	Group IGroup `field:"optional" json:"group" yaml:"group"`
	// The name of the schedule.
	//
	// Up to 64 letters (uppercase and lowercase), numbers, hyphens, underscores and dots are allowed.
	// Default: - A unique name will be generated.
	//
	// Experimental.
	ScheduleName *string `field:"optional" json:"scheduleName" yaml:"scheduleName"`
}

