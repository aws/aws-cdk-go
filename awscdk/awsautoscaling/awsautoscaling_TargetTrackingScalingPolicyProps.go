package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
)

// Properties for a concrete TargetTrackingPolicy.
//
// Adds the scalingTarget.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var duration duration
//   var metric metric
//
//   targetTrackingScalingPolicyProps := &targetTrackingScalingPolicyProps{
//   	autoScalingGroup: autoScalingGroup,
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	cooldown: duration,
//   	customMetric: metric,
//   	disableScaleIn: jsii.Boolean(false),
//   	estimatedInstanceWarmup: duration,
//   	predefinedMetric: awscdk.Aws_autoscaling.predefinedMetric_ASG_AVERAGE_CPU_UTILIZATION,
//   	resourceLabel: jsii.String("resourceLabel"),
//   }
//
// Experimental.
type TargetTrackingScalingPolicyProps struct {
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
	// The target value for the metric.
	// Experimental.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// A custom metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	// Experimental.
	CustomMetric awscloudwatch.IMetric `field:"optional" json:"customMetric" yaml:"customMetric"`
	// A predefined metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	// Experimental.
	PredefinedMetric PredefinedMetric `field:"optional" json:"predefinedMetric" yaml:"predefinedMetric"`
	// The resource label associated with the predefined metric.
	//
	// Should be supplied if the predefined metric is ALBRequestCountPerTarget, and the
	// format should be:
	//
	// app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>.
	// Experimental.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
	// Experimental.
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

