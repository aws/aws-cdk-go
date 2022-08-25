package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for enabling Lambda utilization tracking.
//
// Example:
//   import autoscaling "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   alias := fn.addAlias(jsii.String("prod"))
//
//   // Create AutoScaling target
//   as := alias.addAutoScaling(&autoScalingOptions{
//   	maxCapacity: jsii.Number(50),
//   })
//
//   // Configure Target Tracking
//   as.scaleOnUtilization(&utilizationScalingOptions{
//   	utilizationTarget: jsii.Number(0.5),
//   })
//
//   // Configure Scheduled Scaling
//   as.scaleOnSchedule(jsii.String("ScaleUpInTheMorning"), &scalingSchedule{
//   	schedule: autoscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(20),
//   })
//
type UtilizationScalingOptions struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// Utilization target for the attribute.
	//
	// For example, .5 indicates that 50 percent of allocated provisioned concurrency is in use.
	UtilizationTarget *float64 `field:"required" json:"utilizationTarget" yaml:"utilizationTarget"`
}

