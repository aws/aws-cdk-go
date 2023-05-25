package awsapplicationautoscaling


// `TargetTrackingScalingPolicyConfiguration` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a target tracking scaling policy configuration for Application Auto Scaling. Use a target tracking scaling policy to adjust the capacity of the specified scalable target in response to actual workloads, so that resource utilization remains at or near the target utilization value.
//
// For more information, see [Target tracking scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingScalingPolicyConfigurationProperty := &TargetTrackingScalingPolicyConfigurationProperty{
//   	TargetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	CustomizedMetricSpecification: &CustomizedMetricSpecificationProperty{
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   		Statistic: jsii.String("statistic"),
//
//   		// the properties below are optional
//   		Dimensions: []interface{}{
//   			&MetricDimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Unit: jsii.String("unit"),
//   	},
//   	DisableScaleIn: jsii.Boolean(false),
//   	PredefinedMetricSpecification: &PredefinedMetricSpecificationProperty{
//   		PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	ScaleInCooldown: jsii.Number(123),
//   	ScaleOutCooldown: jsii.Number(123),
//   }
//
type CfnScalingPolicy_TargetTrackingScalingPolicyConfigurationProperty struct {
	// The target value for the metric.
	//
	// Although this property accepts numbers of type Double, it won't accept values that are either too small or too large. Values must be in the range of -2^360 to 2^360. The value must be a valid number based on the choice of metric. For example, if the metric is CPU utilization, then the target value is a percent value that represents how much of the CPU can be used before scaling out.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// A customized metric.
	//
	// You can specify either a predefined metric or a customized metric.
	CustomizedMetricSpecification interface{} `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Indicates whether scale in by the target tracking scaling policy is disabled.
	//
	// If the value is `true` , scale in is disabled and the target tracking scaling policy won't remove capacity from the scalable target. Otherwise, scale in is enabled and the target tracking scaling policy can remove capacity from the scalable target. The default value is `false` .
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A predefined metric.
	//
	// You can specify either a predefined metric or a customized metric.
	PredefinedMetricSpecification interface{} `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
	// The amount of time, in seconds, after a scale-in activity completes before another scale-in activity can start.
	//
	// For more information and for default values, see [Define cooldown periods](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html#target-tracking-cooldown) in the *Application Auto Scaling User Guide* .
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// The amount of time, in seconds, to wait for a previous scale-out activity to take effect.
	//
	// For more information and for default values, see [Define cooldown periods](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html#target-tracking-cooldown) in the *Application Auto Scaling User Guide* .
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

