package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// The properties for enabling scaling based on Application Load Balancer (ALB) request counts.
//
// Example:
//   var target applicationTargetGroup
//   var service baseService
//
//   scaling := service.AutoScaleTaskCount(&EnableScalingProps{
//   	MaxCapacity: jsii.Number(10),
//   })
//   scaling.ScaleOnCpuUtilization(jsii.String("CpuScaling"), &CpuUtilizationScalingProps{
//   	TargetUtilizationPercent: jsii.Number(50),
//   })
//
//   scaling.ScaleOnRequestCount(jsii.String("RequestScaling"), &RequestCountScalingProps{
//   	RequestsPerTarget: jsii.Number(10000),
//   	TargetGroup: target,
//   })
//
// Experimental.
type RequestCountScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	// Experimental.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Experimental.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Experimental.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The number of ALB requests per target.
	// Experimental.
	RequestsPerTarget *float64 `field:"required" json:"requestsPerTarget" yaml:"requestsPerTarget"`
	// The ALB target group name.
	// Experimental.
	TargetGroup awselasticloadbalancingv2.ApplicationTargetGroup `field:"required" json:"targetGroup" yaml:"targetGroup"`
}

