package awscdkscheduleralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Construction properties for `Schedule`.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   var deliveryStream cfnDeliveryStream
//
//
//   payload := map[string]*string{
//   	"Data": jsii.String("record"),
//   }
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewKinesisDataFirehosePutRecord(deliveryStream, &ScheduleTargetBaseProps{
//   		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
//   	}),
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
	// Indicates whether the schedule is enabled.
	// Default: true.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The schedule's group.
	// Default: - By default a schedule will be associated with the `default` group.
	//
	// Experimental.
	Group IGroup `field:"optional" json:"group" yaml:"group"`
	// The customer managed KMS key that EventBridge Scheduler will use to encrypt and decrypt your data.
	// Default: - All events in Scheduler are encrypted with a key that AWS owns and manages.
	//
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// The name of the schedule.
	//
	// Up to 64 letters (uppercase and lowercase), numbers, hyphens, underscores and dots are allowed.
	// Default: - A unique name will be generated.
	//
	// Experimental.
	ScheduleName *string `field:"optional" json:"scheduleName" yaml:"scheduleName"`
	// Allows to override target properties when creating a new schedule.
	// Experimental.
	TargetOverrides *ScheduleTargetProps `field:"optional" json:"targetOverrides" yaml:"targetOverrides"`
}

