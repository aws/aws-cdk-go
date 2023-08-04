package awsautoscaling

import (
	"time"
)

// Properties for a scheduled action on an AutoScalingGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var schedule schedule
//
//   scheduledActionProps := &ScheduledActionProps{
//   	AutoScalingGroup: autoScalingGroup,
//   	Schedule: schedule,
//
//   	// the properties below are optional
//   	DesiredCapacity: jsii.Number(123),
//   	EndTime: NewDate(),
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   	StartTime: NewDate(),
//   	TimeZone: jsii.String("timeZone"),
//   }
//
type ScheduledActionProps struct {
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
	// The AutoScalingGroup to apply the scheduled actions to.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

