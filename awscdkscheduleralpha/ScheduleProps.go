package awscdkscheduleralpha

import (
	"time"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Construction properties for `Schedule`.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   var deliveryStream iDeliveryStream
//
//
//   payload := map[string]*string{
//   	"Data": jsii.String("record"),
//   }
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewFirehosePutRecord(deliveryStream, &ScheduleTargetBaseProps{
//   		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
//   	}),
//   })
//
// Deprecated.
type ScheduleProps struct {
	// The expression that defines when the schedule runs.
	//
	// Can be either a `at`, `rate`
	// or `cron` expression.
	// Deprecated.
	Schedule ScheduleExpression `field:"required" json:"schedule" yaml:"schedule"`
	// The schedule's target details.
	// Deprecated.
	Target IScheduleTarget `field:"required" json:"target" yaml:"target"`
	// The description you specify for the schedule.
	// Default: - no value.
	//
	// Deprecated.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the schedule is enabled.
	// Default: true.
	//
	// Deprecated.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The date, in UTC, before which the schedule can invoke its target.
	//
	// EventBridge Scheduler ignores end for one-time schedules.
	// Default: - no value.
	//
	// Deprecated.
	End *time.Time `field:"optional" json:"end" yaml:"end"`
	// The schedule's group.
	// Default: - By default a schedule will be associated with the `default` group.
	//
	// Deprecated: Use `scheduleGroup` instead. `group` will be removed when this module is stabilized.
	Group IGroup `field:"optional" json:"group" yaml:"group"`
	// The customer managed KMS key that EventBridge Scheduler will use to encrypt and decrypt your data.
	// Default: - All events in Scheduler are encrypted with a key that AWS owns and manages.
	//
	// Deprecated.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// The schedule's group.
	// Default: - By default a schedule will be associated with the `default` group.
	//
	// Deprecated.
	ScheduleGroup IScheduleGroup `field:"optional" json:"scheduleGroup" yaml:"scheduleGroup"`
	// The name of the schedule.
	//
	// Up to 64 letters (uppercase and lowercase), numbers, hyphens, underscores and dots are allowed.
	// Default: - A unique name will be generated.
	//
	// Deprecated.
	ScheduleName *string `field:"optional" json:"scheduleName" yaml:"scheduleName"`
	// The date, in UTC, after which the schedule can begin invoking its target.
	//
	// EventBridge Scheduler ignores start for one-time schedules.
	// Default: - no value.
	//
	// Deprecated.
	Start *time.Time `field:"optional" json:"start" yaml:"start"`
	// A time window during which EventBridge Scheduler invokes the schedule.
	// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-schedule-flexible-time-windows.html
	//
	// Default: TimeWindow.off()
	//
	// Deprecated.
	TimeWindow TimeWindow `field:"optional" json:"timeWindow" yaml:"timeWindow"`
}

