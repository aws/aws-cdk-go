package awsautoscaling


// Properties for defining a `CfnScalingPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScalingPolicyProps := &CfnScalingPolicyProps{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//
//   	// the properties below are optional
//   	AdjustmentType: jsii.String("adjustmentType"),
//   	Cooldown: jsii.String("cooldown"),
//   	EstimatedInstanceWarmup: jsii.Number(123),
//   	MetricAggregationType: jsii.String("metricAggregationType"),
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   	PolicyType: jsii.String("policyType"),
//   	PredictiveScalingConfiguration: &PredictiveScalingConfigurationProperty{
//   		MetricSpecifications: []interface{}{
//   			&PredictiveScalingMetricSpecificationProperty{
//   				TargetValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				CustomizedCapacityMetricSpecification: &PredictiveScalingCustomizedCapacityMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Id: jsii.String("id"),
//
//   							// the properties below are optional
//   							Expression: jsii.String("expression"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								Stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				CustomizedLoadMetricSpecification: &PredictiveScalingCustomizedLoadMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Id: jsii.String("id"),
//
//   							// the properties below are optional
//   							Expression: jsii.String("expression"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								Stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				CustomizedScalingMetricSpecification: &PredictiveScalingCustomizedScalingMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Id: jsii.String("id"),
//
//   							// the properties below are optional
//   							Expression: jsii.String("expression"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								Stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				PredefinedLoadMetricSpecification: &PredictiveScalingPredefinedLoadMetricProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				PredefinedMetricPairSpecification: &PredictiveScalingPredefinedMetricPairProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				PredefinedScalingMetricSpecification: &PredictiveScalingPredefinedScalingMetricProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		MaxCapacityBreachBehavior: jsii.String("maxCapacityBreachBehavior"),
//   		MaxCapacityBuffer: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   		SchedulingBufferTime: jsii.Number(123),
//   	},
//   	ScalingAdjustment: jsii.Number(123),
//   	StepAdjustments: []interface{}{
//   		&StepAdjustmentProperty{
//   			ScalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			MetricIntervalLowerBound: jsii.Number(123),
//   			MetricIntervalUpperBound: jsii.Number(123),
//   		},
//   	},
//   	TargetTrackingConfiguration: &TargetTrackingConfigurationProperty{
//   		TargetValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		CustomizedMetricSpecification: &CustomizedMetricSpecificationProperty{
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//   			Statistic: jsii.String("statistic"),
//
//   			// the properties below are optional
//   			Dimensions: []interface{}{
//   				&MetricDimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Unit: jsii.String("unit"),
//   		},
//   		DisableScaleIn: jsii.Boolean(false),
//   		PredefinedMetricSpecification: &PredefinedMetricSpecificationProperty{
//   			PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   			// the properties below are optional
//   			ResourceLabel: jsii.String("resourceLabel"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html
//
type CfnScalingPolicyProps struct {
	// The name of the Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-autoscalinggroupname
	//
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Specifies how the scaling adjustment is interpreted (for example, an absolute number or a percentage).
	//
	// The valid values are `ChangeInCapacity` , `ExactCapacity` , and `PercentChangeInCapacity` .
	//
	// Required if the policy type is `StepScaling` or `SimpleScaling` . For more information, see [Scaling adjustment types](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment) in the *Amazon EC2 Auto Scaling User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-adjustmenttype
	//
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// A cooldown period, in seconds, that applies to a specific simple scaling policy.
	//
	// When a cooldown period is specified here, it overrides the default cooldown.
	//
	// Valid only if the policy type is `SimpleScaling` . For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Default: None.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-cooldown
	//
	Cooldown *string `field:"optional" json:"cooldown" yaml:"cooldown"`
	// *Not needed if the default instance warmup is defined for the group.*.
	//
	// The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics. This warm-up period applies to instances launched due to a specific target tracking or step scaling policy. When a warm-up period is specified here, it overrides the default instance warmup.
	//
	// Valid only if the policy type is `TargetTrackingScaling` or `StepScaling` .
	//
	// > The default is to use the value for the default instance warmup defined for the group. If default instance warmup is null, then `EstimatedInstanceWarmup` falls back to the value of default cooldown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-estimatedinstancewarmup
	//
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// The aggregation type for the CloudWatch metrics.
	//
	// The valid values are `Minimum` , `Maximum` , and `Average` . If the aggregation type is null, the value is treated as `Average` .
	//
	// Valid only if the policy type is `StepScaling` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-metricaggregationtype
	//
	MetricAggregationType *string `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is `PercentChangeInCapacity` .
	//
	// For example, suppose that you create a step scaling policy to scale out an Auto Scaling group by 25 percent and you specify a `MinAdjustmentMagnitude` of 2. If the group has 4 instances and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a `MinAdjustmentMagnitude` of 2, Amazon EC2 Auto Scaling scales out the group by 2 instances.
	//
	// Valid only if the policy type is `StepScaling` or `SimpleScaling` . For more information, see [Scaling adjustment types](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// > Some Auto Scaling groups use instance weights. In this case, set the `MinAdjustmentMagnitude` to a value that is at least as large as your largest instance weight.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-minadjustmentmagnitude
	//
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// One of the following policy types:.
	//
	// - `TargetTrackingScaling`
	// - `StepScaling`
	// - `SimpleScaling` (default)
	// - `PredictiveScaling`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-policytype
	//
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// A predictive scaling policy. Provides support for predefined and custom metrics.
	//
	// Predefined metrics include CPU utilization, network in/out, and the Application Load Balancer request count.
	//
	// Required if the policy type is `PredictiveScaling` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-predictivescalingconfiguration
	//
	PredictiveScalingConfiguration interface{} `field:"optional" json:"predictiveScalingConfiguration" yaml:"predictiveScalingConfiguration"`
	// The amount by which to scale, based on the specified adjustment type.
	//
	// A positive value adds to the current capacity while a negative number removes from the current capacity. For exact capacity, you must specify a non-negative value.
	//
	// Required if the policy type is `SimpleScaling` . (Not used with any other policy type.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-scalingadjustment
	//
	ScalingAdjustment *float64 `field:"optional" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// Required if the policy type is `StepScaling` . (Not used with any other policy type.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-stepadjustments
	//
	StepAdjustments interface{} `field:"optional" json:"stepAdjustments" yaml:"stepAdjustments"`
	// A target tracking scaling policy. Provides support for predefined or custom metrics.
	//
	// The following predefined metrics are available:
	//
	// - `ASGAverageCPUUtilization`
	// - `ASGAverageNetworkIn`
	// - `ASGAverageNetworkOut`
	// - `ALBRequestCountPerTarget`
	//
	// If you specify `ALBRequestCountPerTarget` for the metric, you must specify the `ResourceLabel` property with the `PredefinedMetricSpecification` .
	//
	// Required if the policy type is `TargetTrackingScaling` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration
	//
	TargetTrackingConfiguration interface{} `field:"optional" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
}

