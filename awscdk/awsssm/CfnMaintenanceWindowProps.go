package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMaintenanceWindow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMaintenanceWindowProps := &CfnMaintenanceWindowProps{
//   	AllowUnassociatedTargets: jsii.Boolean(false),
//   	Cutoff: jsii.Number(123),
//   	Duration: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EndDate: jsii.String("endDate"),
//   	ScheduleOffset: jsii.Number(123),
//   	ScheduleTimezone: jsii.String("scheduleTimezone"),
//   	StartDate: jsii.String("startDate"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html
//
type CfnMaintenanceWindowProps struct {
	// Enables a maintenance window task to run on managed instances, even if you have not registered those instances as targets.
	//
	// If enabled, then you must specify the unregistered instances (by instance ID) when you register a task with the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-allowunassociatedtargets
	//
	AllowUnassociatedTargets interface{} `field:"required" json:"allowUnassociatedTargets" yaml:"allowUnassociatedTargets"`
	// The number of hours before the end of the maintenance window that AWS Systems Manager stops scheduling new tasks for execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-cutoff
	//
	Cutoff *float64 `field:"required" json:"cutoff" yaml:"cutoff"`
	// The duration of the maintenance window in hours.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-duration
	//
	Duration *float64 `field:"required" json:"duration" yaml:"duration"`
	// The name of the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schedule of the maintenance window in the form of a cron or rate expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-schedule
	//
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// A description of the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-enddate
	//
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// The number of days to wait to run a maintenance window after the scheduled cron expression date and time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-scheduleoffset
	//
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-scheduletimezone
	//
	ScheduleTimezone *string `field:"optional" json:"scheduleTimezone" yaml:"scheduleTimezone"`
	// The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active.
	//
	// `StartDate` allows you to delay activation of the maintenance window until the specified future date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-startdate
	//
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a maintenance window to identify the type of tasks it will run, the types of targets, and the environment it will run in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html#cfn-ssm-maintenancewindow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

