package awscdkscheduleralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import scheduler_alpha "github.com/aws/aws-cdk-go/awscdkscheduleralpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var timeZone timeZone
//
//   cronOptionsWithTimezone := &CronOptionsWithTimezone{
//   	Day: jsii.String("day"),
//   	Hour: jsii.String("hour"),
//   	Minute: jsii.String("minute"),
//   	Month: jsii.String("month"),
//   	TimeZone: timeZone,
//   	WeekDay: jsii.String("weekDay"),
//   	Year: jsii.String("year"),
//   }
//
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html#cron-expressions
//
// Experimental.
type CronOptionsWithTimezone struct {
	// The day of the month to run this rule at.
	// Default: - Every day of the month.
	//
	// Experimental.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The hour to run this rule at.
	// Default: - Every hour.
	//
	// Experimental.
	Hour *string `field:"optional" json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	// Default: - Every minute.
	//
	// Experimental.
	Minute *string `field:"optional" json:"minute" yaml:"minute"`
	// The month to run this rule at.
	// Default: - Every month.
	//
	// Experimental.
	Month *string `field:"optional" json:"month" yaml:"month"`
	// The day of the week to run this rule at.
	// Default: - Any day of the week.
	//
	// Experimental.
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	// Default: - Every year.
	//
	// Experimental.
	Year *string `field:"optional" json:"year" yaml:"year"`
	// The timezone to run the schedule in.
	// Default: - TimeZone.ETC_UTC
	//
	// Experimental.
	TimeZone awscdk.TimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
}

