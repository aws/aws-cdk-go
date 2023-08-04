package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Properties for a concrete TargetTrackingPolicy.
//
// Adds the scalingTarget.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var metric metric
//
//   targetTrackingScalingPolicyProps := &TargetTrackingScalingPolicyProps{
//   	AutoScalingGroup: autoScalingGroup,
//   	TargetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	Cooldown: cdk.Duration_Minutes(jsii.Number(30)),
//   	CustomMetric: metric,
//   	DisableScaleIn: jsii.Boolean(false),
//   	EstimatedInstanceWarmup: cdk.Duration_*Minutes(jsii.Number(30)),
//   	PredefinedMetric: awscdk.Aws_autoscaling.PredefinedMetric_ASG_AVERAGE_CPU_UTILIZATION,
//   	ResourceLabel: jsii.String("resourceLabel"),
//   }
//
type TargetTrackingScalingPolicyProps struct {
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
	// The target value for the metric.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// A custom metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	// Default: - No custom metric.
	//
	CustomMetric awscloudwatch.IMetric `field:"optional" json:"customMetric" yaml:"customMetric"`
	// A predefined metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	// Default: - No predefined metric.
	//
	PredefinedMetric PredefinedMetric `field:"optional" json:"predefinedMetric" yaml:"predefinedMetric"`
	// The resource label associated with the predefined metric.
	//
	// Should be supplied if the predefined metric is ALBRequestCountPerTarget, and the
	// format should be:
	//
	// app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>.
	// Default: - No resource label.
	//
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
	AutoScalingGroup IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

