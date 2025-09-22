package awsautoscaling

import (
	"time"
)

// Properties for a scheduled scaling action.
//
// Example:
//   var autoScalingGroup autoScalingGroup
//
//
//   autoScalingGroup.scaleOnSchedule(jsii.String("PrescaleInTheMorning"), &BasicScheduledActionProps{
//   	Schedule: autoscaling.Schedule_Cron(&CronOptions{
//   		Hour: jsii.String("8"),
//   		Minute: jsii.String("0"),
//   	}),
//   	MinCapacity: jsii.Number(20),
//   })
//
//   autoScalingGroup.scaleOnSchedule(jsii.String("AllowDownscalingAtNight"), &BasicScheduledActionProps{
//   	Schedule: autoscaling.Schedule_*Cron(&CronOptions{
//   		Hour: jsii.String("20"),
//   		Minute: jsii.String("0"),
//   	}),
//   	MinCapacity: jsii.Number(1),
//   })
//
type BasicScheduledActionProps struct {
	// When to perform this action.
	//
	// Supports cron expressions.
	//
	// For more information about cron expressions, see https://en.wikipedia.org/wiki/Cron.
	Schedule Schedule `field:"required" json:"schedule" yaml:"schedule"`
	// The new desired capacity.
	//
	// At the scheduled time, set the desired capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	// Default: - No new desired capacity.
	//
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// When this scheduled action expires.
	// Default: - The rule never expires.
	//
	EndTime *time.Time `field:"optional" json:"endTime" yaml:"endTime"`
	// The new maximum capacity.
	//
	// At the scheduled time, set the maximum capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	// Default: - No new maximum capacity.
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The new minimum capacity.
	//
	// At the scheduled time, set the minimum capacity to the given capacity.
	//
	// At least one of maxCapacity, minCapacity, or desiredCapacity must be supplied.
	// Default: - No new minimum capacity.
	//
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// When this scheduled action becomes active.
	// Default: - The rule is activate immediately.
	//
	StartTime *time.Time `field:"optional" json:"startTime" yaml:"startTime"`
	// Specifies the time zone for a cron expression.
	//
	// If a time zone is not provided, UTC is used by default.
	//
	// Valid values are the canonical names of the IANA time zones, derived from the IANA Time Zone Database (such as Etc/GMT+9 or Pacific/Tahiti).
	//
	// For more information, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
	// Default: - UTC.
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

