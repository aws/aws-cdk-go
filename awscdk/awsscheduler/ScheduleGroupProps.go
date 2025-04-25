package awsscheduler

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a Schedule Group.
//
// Example:
//   var target lambdaInvoke
//
//
//   scheduleGroup := awscdk.NewScheduleGroup(this, jsii.String("ScheduleGroup"), &ScheduleGroupProps{
//   	ScheduleGroupName: jsii.String("MyScheduleGroup"),
//   })
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
//   	Target: Target,
//   	ScheduleGroup: ScheduleGroup,
//   })
//
type ScheduleGroupProps struct {
	// The removal policy for the group.
	//
	// If the group is removed also all schedules are removed.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The name of the schedule group.
	//
	// Up to 64 letters (uppercase and lowercase), numbers, hyphens, underscores and dots are allowed.
	// Default: - A unique name will be generated.
	//
	ScheduleGroupName *string `field:"optional" json:"scheduleGroupName" yaml:"scheduleGroupName"`
}

