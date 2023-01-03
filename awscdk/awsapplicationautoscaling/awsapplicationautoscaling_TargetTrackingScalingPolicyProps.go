package awsapplicationautoscaling

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
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metric metric
//   var scalableTarget scalableTarget
//
//   targetTrackingScalingPolicyProps := &targetTrackingScalingPolicyProps{
//   	scalingTarget: scalableTarget,
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	customMetric: metric,
//   	disableScaleIn: jsii.Boolean(false),
//   	policyName: jsii.String("policyName"),
//   	predefinedMetric: awscdk.Aws_applicationautoscaling.predefinedMetric_APPSTREAM_AVERAGE_CAPACITY_UTILIZATION,
//   	resourceLabel: jsii.String("resourceLabel"),
//   	scaleInCooldown: cdk.duration.minutes(jsii.Number(30)),
//   	scaleOutCooldown: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
type TargetTrackingScalingPolicyProps struct {
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
	// The target value for the metric.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// A custom metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	CustomMetric awscloudwatch.IMetric `field:"optional" json:"customMetric" yaml:"customMetric"`
	// A predefined metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	PredefinedMetric PredefinedMetric `field:"optional" json:"predefinedMetric" yaml:"predefinedMetric"`
	// Identify the resource associated with the metric type.
	//
	// Only used for predefined metric ALBRequestCountPerTarget.
	//
	// Example value: `app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>`.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
	ScalingTarget IScalableTarget `field:"required" json:"scalingTarget" yaml:"scalingTarget"`
}

