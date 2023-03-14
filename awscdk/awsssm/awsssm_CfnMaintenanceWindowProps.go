package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnMaintenanceWindow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMaintenanceWindowProps := &cfnMaintenanceWindowProps{
//   	allowUnassociatedTargets: jsii.Boolean(false),
//   	cutoff: jsii.Number(123),
//   	duration: jsii.Number(123),
//   	name: jsii.String("name"),
//   	schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	endDate: jsii.String("endDate"),
//   	scheduleOffset: jsii.Number(123),
//   	scheduleTimezone: jsii.String("scheduleTimezone"),
//   	startDate: jsii.String("startDate"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMaintenanceWindowProps struct {
	// Enables a maintenance window task to run on managed instances, even if you have not registered those instances as targets.
	//
	// If enabled, then you must specify the unregistered instances (by instance ID) when you register a task with the maintenance window.
	AllowUnassociatedTargets interface{} `field:"required" json:"allowUnassociatedTargets" yaml:"allowUnassociatedTargets"`
	// The number of hours before the end of the maintenance window that AWS Systems Manager stops scheduling new tasks for execution.
	Cutoff *float64 `field:"required" json:"cutoff" yaml:"cutoff"`
	// The duration of the maintenance window in hours.
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// The name of the maintenance window.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// A description of the maintenance window.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The number of days to wait to run a maintenance window after the scheduled cron expression date and time.
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format.
	ScheduleTimezone *string `field:"optional" json:"scheduleTimezone" yaml:"scheduleTimezone"`
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active.
	//
	// StartDate allows you to delay activation of the Maintenance Window until the specified future date.
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a maintenance window to identify the type of tasks it will run, the types of targets, and the environment it will run in.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

