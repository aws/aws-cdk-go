package awsautoscalingplans


// `TargetTrackingConfiguration` is a subproperty of [ScalingInstruction](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html) that specifies a target tracking configuration to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingConfigurationProperty := &targetTrackingConfigurationProperty{
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	customizedScalingMetricSpecification: &customizedScalingMetricSpecificationProperty{
//   		metricName: jsii.String("metricName"),
//   		namespace: jsii.String("namespace"),
//   		statistic: jsii.String("statistic"),
//
//   		// the properties below are optional
//   		dimensions: []interface{}{
//   			&metricDimensionProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		unit: jsii.String("unit"),
//   	},
//   	disableScaleIn: jsii.Boolean(false),
//   	estimatedInstanceWarmup: jsii.Number(123),
//   	predefinedScalingMetricSpecification: &predefinedScalingMetricSpecificationProperty{
//   		predefinedScalingMetricType: jsii.String("predefinedScalingMetricType"),
//
//   		// the properties below are optional
//   		resourceLabel: jsii.String("resourceLabel"),
//   	},
//   	scaleInCooldown: jsii.Number(123),
//   	scaleOutCooldown: jsii.Number(123),
//   }
//
type CfnScalingPlan_TargetTrackingConfigurationProperty struct {
	// The target value for the metric.
	//
	// Although this property accepts numbers of type Double, it won't accept values that are either too small or too large. Values must be in the range of -2^360 to 2^360.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// A customized metric.
	//
	// You can specify either a predefined metric or a customized metric.
	CustomizedScalingMetricSpecification interface{} `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// Indicates whether scale in by the target tracking scaling policy is disabled.
	//
	// If the value is `true` , scale in is disabled and the target tracking scaling policy doesn't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking scaling policy can remove capacity from the scalable resource.
	//
	// The default value is `false` .
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	//
	// This value is used only if the resource is an Auto Scaling group.
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// A predefined metric.
	//
	// You can specify either a predefined metric or a customized metric.
	PredefinedScalingMetricSpecification interface{} `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
	// The amount of time, in seconds, after a scale-in activity completes before another scale in activity can start.
	//
	// This value is not used if the scalable resource is an Auto Scaling group.
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// The amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	//
	// This value is not used if the scalable resource is an Auto Scaling group.
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

