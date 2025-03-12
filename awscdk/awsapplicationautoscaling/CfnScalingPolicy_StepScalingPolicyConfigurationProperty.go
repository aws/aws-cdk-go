package awsapplicationautoscaling


// `StepScalingPolicyConfiguration` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a step scaling policy configuration for Application Auto Scaling.
//
// For more information, see [Step scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepScalingPolicyConfigurationProperty := &StepScalingPolicyConfigurationProperty{
//   	AdjustmentType: jsii.String("adjustmentType"),
//   	Cooldown: jsii.Number(123),
//   	MetricAggregationType: jsii.String("metricAggregationType"),
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   	StepAdjustments: []interface{}{
//   		&StepAdjustmentProperty{
//   			ScalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			MetricIntervalLowerBound: jsii.Number(123),
//   			MetricIntervalUpperBound: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html
//
type CfnScalingPolicy_StepScalingPolicyConfigurationProperty struct {
	// Specifies whether the `ScalingAdjustment` value in the `StepAdjustment` property is an absolute number or a percentage of the current capacity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-adjustmenttype
	//
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, to wait for a previous scaling activity to take effect.
	//
	// If not specified, the default value is 300. For more information, see [Cooldown period](https://docs.aws.amazon.com/autoscaling/application/userguide/step-scaling-policy-overview.html#step-scaling-cooldown) in the *Application Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-cooldown
	//
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The aggregation type for the CloudWatch metrics.
	//
	// Valid values are `Minimum` , `Maximum` , and `Average` . If the aggregation type is null, the value is treated as `Average` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-metricaggregationtype
	//
	MetricAggregationType *string `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is `PercentChangeInCapacity` .
	//
	// For example, suppose that you create a step scaling policy to scale out an Amazon ECS service by 25 percent and you specify a `MinAdjustmentMagnitude` of 2. If the service has 4 tasks and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a `MinAdjustmentMagnitude` of 2, Application Auto Scaling scales out the service by 2 tasks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-minadjustmentmagnitude
	//
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// At least one step adjustment is required if you are adding a new step scaling policy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration-stepadjustments
	//
	StepAdjustments interface{} `field:"optional" json:"stepAdjustments" yaml:"stepAdjustments"`
}

