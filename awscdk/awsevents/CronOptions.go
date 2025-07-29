package awsevents


// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   rule := events.NewRule(this, jsii.String("Schedule Rule"), &RuleProps{
//   	Schedule: events.Schedule_Cron(&CronOptions{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("4"),
//   	}),
//   })
//   rule.AddTarget(targets.NewLambdaFunction(fn))
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
	// The day of the week to run this rule at.
	// Default: - Any day of the week.
	//
	WeekDay *string `field:"optional" json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	// Default: - Every year.
	//
	Year *string `field:"optional" json:"year" yaml:"year"`
}

