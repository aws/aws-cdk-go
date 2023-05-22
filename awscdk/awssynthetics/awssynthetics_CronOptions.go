package awssynthetics


// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// Example:
//   schedule := synthetics.Schedule_Cron(&CronOptions{
//   	Hour: jsii.String("0,8,16"),
//   })
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_cron.html
//
// Experimental.
type CronOptions struct {
	// The day of the month to run this rule at.
	// Experimental.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The hour to run this rule at.
	// Experimental.
	Hour *string `field:"optional" json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	// Experimental.
	Minute *string `field:"optional" json:"minute" yaml:"minute"`
	// The month to run this rule at.
	// Experimental.
	Month *string `field:"optional" json:"month" yaml:"month"`
	// The day of the week to run this rule at.
	// Experimental.
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
}

