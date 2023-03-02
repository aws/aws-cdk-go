package awssynthetics


// This structure specifies how often a canary is to make runs and the date and time when it should stop making runs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &scheduleProperty{
//   	expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	durationInSeconds: jsii.String("durationInSeconds"),
//   }
//
type CfnCanary_ScheduleProperty struct {
	// A `rate` expression or a `cron` expression that defines how often the canary is to run.
	//
	// For a rate expression, The syntax is `rate( *number unit* )` . *unit* can be `minute` , `minutes` , or `hour` .
	//
	// For example, `rate(1 minute)` runs the canary once a minute, `rate(10 minutes)` runs it once every 10 minutes, and `rate(1 hour)` runs it once every hour. You can specify a frequency between `rate(1 minute)` and `rate(1 hour)` .
	//
	// Specifying `rate(0 minute)` or `rate(0 hour)` is a special value that causes the canary to run only once when it is started.
	//
	// Use `cron( *expression* )` to specify a cron expression. You can't schedule a canary to wait for more than a year before running. For information about the syntax for cron expressions, see [Scheduling canary runs using cron](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_cron.html) .
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// How long, in seconds, for the canary to continue making regular runs according to the schedule in the `Expression` value.
	//
	// If you specify 0, the canary continues making runs until you stop it. If you omit this field, the default of 0 is used.
	DurationInSeconds *string `field:"optional" json:"durationInSeconds" yaml:"durationInSeconds"`
}

