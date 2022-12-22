package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for enabling scaling based on network utilization.
//
// Example:
//   var autoScalingGroup autoScalingGroup
//
//
//   autoScalingGroup.scaleOnIncomingBytes(jsii.String("LimitIngressPerInstance"), &networkUtilizationScalingProps{
//   	targetBytesPerSecond: jsii.Number(10 * 1024 * 1024),
//   })
//   autoScalingGroup.scaleOnOutgoingBytes(jsii.String("LimitEgressPerInstance"), &networkUtilizationScalingProps{
//   	targetBytesPerSecond: jsii.Number(10 * 1024 * 1024),
//   })
//
// Experimental.
type NetworkUtilizationScalingProps struct {
	// Period after a scaling completes before another scaling activity can start.
	// Experimental.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	// Experimental.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	// Experimental.
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// Target average bytes/seconds on each instance.
	// Experimental.
	TargetBytesPerSecond *float64 `field:"required" json:"targetBytesPerSecond" yaml:"targetBytesPerSecond"`
}

