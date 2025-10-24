package awsscheduler

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// Example:
//   var target LambdaInvoke
//
//
//   rateBasedSchedule := awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
//   	Target: Target,
//   	Description: jsii.String("This is a test rate-based schedule"),
//   })
//
//   cronBasedSchedule := awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Cron(&CronOptionsWithTimezone{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("23"),
//   		Day: jsii.String("20"),
//   		Month: jsii.String("11"),
//   		TimeZone: awscdk.TimeZone_AMERICA_NEW_YORK(),
//   	}),
//   	Target: Target,
//   	Description: jsii.String("This is a test cron-based schedule that will run at 11:00 PM, on day 20 of the month, only in November in New York timezone"),
//   })
//
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html#cron-expressions
//
type CronOptionsWithTimezone struct {
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
	// The day of the week to run this rule at.
	// Default: - Any day of the week.
	//
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	// Default: - Every year.
	//
	Year *string `field:"optional" json:"year" yaml:"year"`
	// The timezone to run the schedule in.
	// Default: - TimeZone.ETC_UTC
	//
	TimeZone awscdk.TimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
}

