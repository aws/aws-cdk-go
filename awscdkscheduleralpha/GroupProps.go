package awscdkscheduleralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   group := awscdkscheduleralpha.NewGroup(this, jsii.String("Group"), &GroupProps{
//   	GroupName: jsii.String("MyGroup"),
//   })
//
//   cloudwatch.NewAlarm(this, jsii.String("MyGroupErrorAlarm"), &AlarmProps{
//   	Metric: group.metricTargetErrors(),
//   	EvaluationPeriods: jsii.Number(1),
//   	Threshold: jsii.Number(0),
//   })
//
//   // Or use default group
//   defaultGroup := awscdkscheduleralpha.Group_FromDefaultGroup(this, jsii.String("DefaultGroup"))
//   cloudwatch.NewAlarm(this, jsii.String("DefaultGroupErrorAlarm"), &AlarmProps{
//   	Metric: defaultGroup.MetricTargetErrors(),
//   	EvaluationPeriods: jsii.Number(1),
//   	Threshold: jsii.Number(0),
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

