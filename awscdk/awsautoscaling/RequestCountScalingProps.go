package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for enabling scaling based on request/second.
//
// Example:
//   var autoScalingGroup autoScalingGroup
//
//
//   autoScalingGroup.scaleOnRequestCount(jsii.String("LimitRPS"), &RequestCountScalingProps{
//   	TargetRequestsPerSecond: jsii.Number(1000),
//   })
//
type RequestCountScalingProps struct {
	// Period after a scaling completes before another scaling activity can start.
	// Default: - The default cooldown configured on the AutoScalingGroup.
	//
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the autoscaling group. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// group.
	// Default: false.
	//
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	// Default: - Same as the cooldown.
	//
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// Target average requests/minute on each instance.
	// Default: - Specify exactly one of 'targetRequestsPerMinute' and 'targetRequestsPerSecond'.
	//
	TargetRequestsPerMinute *float64 `field:"optional" json:"targetRequestsPerMinute" yaml:"targetRequestsPerMinute"`
}

