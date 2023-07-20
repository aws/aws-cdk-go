package awsautoscalingplans


// `TargetTrackingConfiguration` is a subproperty of [ScalingInstruction](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html) that specifies a target tracking configuration to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingConfigurationProperty := &TargetTrackingConfigurationProperty{
//   	TargetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	CustomizedScalingMetricSpecification: &CustomizedScalingMetricSpecificationProperty{
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
//   	EstimatedInstanceWarmup: jsii.Number(123),
//   	PredefinedScalingMetricSpecification: &PredefinedScalingMetricSpecificationProperty{
//   		PredefinedScalingMetricType: jsii.String("predefinedScalingMetricType"),
//
//   		// the properties below are optional
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	ScaleInCooldown: jsii.Number(123),
//   	ScaleOutCooldown: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html
//
type CfnScalingPlan_TargetTrackingConfigurationProperty struct {
	// The target value for the metric.
	//
	// Although this property accepts numbers of type Double, it won't accept values that are either too small or too large. Values must be in the range of -2^360 to 2^360.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html#cfn-autoscalingplans-scalingplan-targettrackingconfiguration-targetvalue
	//
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// A customized metric.
	//
	// You can specify either a predefined metric or a customized metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html#cfn-autoscalingplans-scalingplan-targettrackingconfiguration-customizedscalingmetricspecification
	//
	CustomizedScalingMetricSpecification interface{} `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// Indicates whether scale in by the target tracking scaling policy is disabled.
	//
	// If the value is `true` , scale in is disabled and the target tracking scaling policy doesn't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking scaling policy can remove capacity from the scalable resource.
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html#cfn-autoscalingplans-scalingplan-targettrackingconfiguration-disablescalein
	//
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	//
	// This value is used only if the resource is an Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html#cfn-autoscalingplans-scalingplan-targettrackingconfiguration-estimatedinstancewarmup
	//
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// A predefined metric.
	//
	// You can specify either a predefined metric or a customized metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html#cfn-autoscalingplans-scalingplan-targettrackingconfiguration-predefinedscalingmetricspecification
	//
	PredefinedScalingMetricSpecification interface{} `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
	// The amount of time, in seconds, after a scale-in activity completes before another scale in activity can start.
	//
	// This value is not used if the scalable resource is an Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html#cfn-autoscalingplans-scalingplan-targettrackingconfiguration-scaleincooldown
	//
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// The amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	//
	// This value is not used if the scalable resource is an Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html#cfn-autoscalingplans-scalingplan-targettrackingconfiguration-scaleoutcooldown
	//
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

