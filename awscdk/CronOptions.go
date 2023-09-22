package awscdk


// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
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
//   	Schedule: appscaling.Schedule_Cron(&CronOptions{
//   		Hour: jsii.String("8"),
//   		Minute: jsii.String("0"),
//   	}),
//   	MinCapacity: jsii.Number(20),
//   })
//
//   capacity.scaleOnSchedule(jsii.String("AllowDownscalingAtNight"), &scalingSchedule{
//   	Schedule: appscaling.Schedule_*Cron(&CronOptions{
//   		Hour: jsii.String("20"),
//   		Minute: jsii.String("0"),
//   	}),
//   	MinCapacity: jsii.Number(1),
//   })
//
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html#cron-expressions
//
type CronOptions struct {
	// The day of the month to run this rule at.
	// Default: - Every day of the month.
	//
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The hour to run this rule at.
	// Default: - Every hour.
	//
	Hour *string `field:"optional" json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	// Default: - Every minute.
	//
	Minute *string `field:"optional" json:"minute" yaml:"minute"`
	// The month to run this rule at.
	// Default: - Every month.
	//
	Month *string `field:"optional" json:"month" yaml:"month"`
	// Retrieve the expression for this schedule.
	// Default: TimeZone.ETC_UTC
	//
	TimeZone TimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The day of the week to run this rule at.
	// Default: - Any day of the week.
	//
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	// Default: - Every year.
	//
	Year *string `field:"optional" json:"year" yaml:"year"`
}

