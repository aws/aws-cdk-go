package awscdkscheduleralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   var target lambdaInvoke
//
//
//   group := awscdkscheduleralpha.NewGroup(this, jsii.String("Group"), &GroupProps{
//   	GroupName: jsii.String("MyGroup"),
//   })
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
//   	Target: Target,
//   	Group: Group,
//   })
//
// Experimental.
type GroupProps struct {
	// The name of the schedule group.
	//
	// Up to 64 letters (uppercase and lowercase), numbers, hyphens, underscores and dots are allowed.
	// Default: - A unique name will be generated.
	//
	// Experimental.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The removal policy for the group.
	//
	// If the group is removed also all schedules are removed.
	// Default: RemovalPolicy.RETAIN
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

