package awsapplicationautoscaling

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
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var timeZone timeZone
//
//   cronOptions := &CronOptions{
//   	Day: jsii.String("day"),
//   	Hour: jsii.String("hour"),
//   	Minute: jsii.String("minute"),
//   	Month: jsii.String("month"),
//   	TimeZone: timeZone,
//   	WeekDay: jsii.String("weekDay"),
//   	Year: jsii.String("year"),
//   }
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
//
// Deprecated: use core.CronOptions instead
type CronOptions struct {
	// The day of the month to run this rule at.
	// Default: - Every day of the month.
	//
	// Deprecated: use core.CronOptions instead
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The hour to run this rule at.
	// Default: - Every hour.
	//
	// Deprecated: use core.CronOptions instead
	Hour *string `field:"optional" json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	// Default: - Every minute.
	//
	// Deprecated: use core.CronOptions instead
	Minute *string `field:"optional" json:"minute" yaml:"minute"`
	// The month to run this rule at.
	// Default: - Every month.
	//
	// Deprecated: use core.CronOptions instead
	Month *string `field:"optional" json:"month" yaml:"month"`
	// Retrieve the expression for this schedule.
	// Default: TimeZone.ETC_UTC
	//
	// Deprecated: use core.CronOptions instead
	TimeZone awscdk.TimeZone `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The day of the week to run this rule at.
	// Default: - Any day of the week.
	//
	// Deprecated: use core.CronOptions instead
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	// Default: - Every year.
	//
	// Deprecated: use core.CronOptions instead
	Year *string `field:"optional" json:"year" yaml:"year"`
}

