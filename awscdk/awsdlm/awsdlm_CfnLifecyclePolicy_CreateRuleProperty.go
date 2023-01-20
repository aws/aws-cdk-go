package awsdlm


// Specifies when to create snapshots of EBS volumes.
//
// You must specify either a Cron expression or an interval, interval unit, and start time. You cannot specify both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createRuleProperty := &createRuleProperty{
//   	cronExpression: jsii.String("cronExpression"),
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   	location: jsii.String("location"),
//   	times: []*string{
//   		jsii.String("times"),
//   	},
//   }
//
type CfnLifecyclePolicy_CreateRuleProperty struct {
	// The schedule, as a Cron expression.
	//
	// The schedule interval must be between 1 hour and 1 year. For more information, see [Cron expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions) in the *Amazon CloudWatch User Guide* .
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// The interval between snapshots.
	//
	// The supported values are 1, 2, 3, 4, 6, 8, 12, and 24.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The interval unit.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Specifies the destination for snapshots created by the policy.
	//
	// To create snapshots in the same Region as the source resource, specify `CLOUD` . To create snapshots on the same Outpost as the source resource, specify `OUTPOST_LOCAL` . If you omit this parameter, `CLOUD` is used by default.
	//
	// If the policy targets resources in an AWS Region , then you must create snapshots in the same Region as the source resource.
	//
	// If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost as the source resource, or in the Region of that Outpost.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The time, in UTC, to start the operation. The supported format is hh:mm.
	//
	// The operation occurs within a one-hour window following the specified time. If you do not specify a time, Amazon DLM selects a time within the next 24 hours.
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

