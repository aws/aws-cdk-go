package awsautoscaling


// Properties for defining a `CfnScalingPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScalingPolicyProps := &cfnScalingPolicyProps{
//   	autoScalingGroupName: jsii.String("autoScalingGroupName"),
//
//   	// the properties below are optional
//   	adjustmentType: jsii.String("adjustmentType"),
//   	cooldown: jsii.String("cooldown"),
//   	estimatedInstanceWarmup: jsii.Number(123),
//   	metricAggregationType: jsii.String("metricAggregationType"),
//   	minAdjustmentMagnitude: jsii.Number(123),
//   	policyType: jsii.String("policyType"),
//   	predictiveScalingConfiguration: &predictiveScalingConfigurationProperty{
//   		metricSpecifications: []interface{}{
//   			&predictiveScalingMetricSpecificationProperty{
//   				targetValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				customizedCapacityMetricSpecification: &predictiveScalingCustomizedCapacityMetricProperty{
//   					metricDataQueries: []interface{}{
//   						&metricDataQueryProperty{
//   							id: jsii.String("id"),
//
//   							// the properties below are optional
//   							expression: jsii.String("expression"),
//   							label: jsii.String("label"),
//   							metricStat: &metricStatProperty{
//   								metric: &metricProperty{
//   									metricName: jsii.String("metricName"),
//   									namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									dimensions: []interface{}{
//   										&metricDimensionProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								unit: jsii.String("unit"),
//   							},
//   							returnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				customizedLoadMetricSpecification: &predictiveScalingCustomizedLoadMetricProperty{
//   					metricDataQueries: []interface{}{
//   						&metricDataQueryProperty{
//   							id: jsii.String("id"),
//
//   							// the properties below are optional
//   							expression: jsii.String("expression"),
//   							label: jsii.String("label"),
//   							metricStat: &metricStatProperty{
//   								metric: &metricProperty{
//   									metricName: jsii.String("metricName"),
//   									namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									dimensions: []interface{}{
//   										&metricDimensionProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								unit: jsii.String("unit"),
//   							},
//   							returnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				customizedScalingMetricSpecification: &predictiveScalingCustomizedScalingMetricProperty{
//   					metricDataQueries: []interface{}{
//   						&metricDataQueryProperty{
//   							id: jsii.String("id"),
//
//   							// the properties below are optional
//   							expression: jsii.String("expression"),
//   							label: jsii.String("label"),
//   							metricStat: &metricStatProperty{
//   								metric: &metricProperty{
//   									metricName: jsii.String("metricName"),
//   									namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									dimensions: []interface{}{
//   										&metricDimensionProperty{
//   											name: jsii.String("name"),
//   											value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								unit: jsii.String("unit"),
//   							},
//   							returnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				predefinedLoadMetricSpecification: &predictiveScalingPredefinedLoadMetricProperty{
//   					predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					resourceLabel: jsii.String("resourceLabel"),
//   				},
//   				predefinedMetricPairSpecification: &predictiveScalingPredefinedMetricPairProperty{
//   					predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					resourceLabel: jsii.String("resourceLabel"),
//   				},
//   				predefinedScalingMetricSpecification: &predictiveScalingPredefinedScalingMetricProperty{
//   					predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					resourceLabel: jsii.String("resourceLabel"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		maxCapacityBreachBehavior: jsii.String("maxCapacityBreachBehavior"),
//   		maxCapacityBuffer: jsii.Number(123),
//   		mode: jsii.String("mode"),
//   		schedulingBufferTime: jsii.Number(123),
//   	},
//   	scalingAdjustment: jsii.Number(123),
//   	stepAdjustments: []interface{}{
//   		&stepAdjustmentProperty{
//   			scalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			metricIntervalLowerBound: jsii.Number(123),
//   			metricIntervalUpperBound: jsii.Number(123),
//   		},
//   	},
//   	targetTrackingConfiguration: &targetTrackingConfigurationProperty{
//   		targetValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		customizedMetricSpecification: &customizedMetricSpecificationProperty{
//   			metricName: jsii.String("metricName"),
//   			namespace: jsii.String("namespace"),
//   			statistic: jsii.String("statistic"),
//
//   			// the properties below are optional
//   			dimensions: []interface{}{
//   				&metricDimensionProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			unit: jsii.String("unit"),
//   		},
//   		disableScaleIn: jsii.Boolean(false),
//   		predefinedMetricSpecification: &predefinedMetricSpecificationProperty{
//   			predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   			// the properties below are optional
//   			resourceLabel: jsii.String("resourceLabel"),
//   		},
//   	},
//   }
//
type CfnScalingPolicyProps struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Specifies how the scaling adjustment is interpreted (for example, an absolute number or a percentage).
	//
	// The valid values are `ChangeInCapacity` , `ExactCapacity` , and `PercentChangeInCapacity` .
	//
	// Required if the policy type is `StepScaling` or `SimpleScaling` . For more information, see [Scaling adjustment types](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment) in the *Amazon EC2 Auto Scaling User Guide* .
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// A cooldown period, in seconds, that applies to a specific simple scaling policy.
	//
	// When a cooldown period is specified here, it overrides the default cooldown.
	//
	// Valid only if the policy type is `SimpleScaling` . For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// Default: None.
	Cooldown *string `field:"optional" json:"cooldown" yaml:"cooldown"`
	// *Not needed if the default instance warmup is defined for the group.*.
	//
	// The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics. This warm-up period applies to instances launched due to a specific target tracking or step scaling policy. When a warm-up period is specified here, it overrides the default instance warmup.
	//
	// Valid only if the policy type is `TargetTrackingScaling` or `StepScaling` .
	//
	// > The default is to use the value for the default instance warmup defined for the group. If default instance warmup is null, then `EstimatedInstanceWarmup` falls back to the value of default cooldown.
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// The aggregation type for the CloudWatch metrics.
	//
	// The valid values are `Minimum` , `Maximum` , and `Average` . If the aggregation type is null, the value is treated as `Average` .
	//
	// Valid only if the policy type is `StepScaling` .
	MetricAggregationType *string `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is `PercentChangeInCapacity` .
	//
	// For example, suppose that you create a step scaling policy to scale out an Auto Scaling group by 25 percent and you specify a `MinAdjustmentMagnitude` of 2. If the group has 4 instances and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a `MinAdjustmentMagnitude` of 2, Amazon EC2 Auto Scaling scales out the group by 2 instances.
	//
	// Valid only if the policy type is `StepScaling` or `SimpleScaling` . For more information, see [Scaling adjustment types](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment) in the *Amazon EC2 Auto Scaling User Guide* .
	//
	// > Some Auto Scaling groups use instance weights. In this case, set the `MinAdjustmentMagnitude` to a value that is at least as large as your largest instance weight.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// One of the following policy types:.
	//
	// - `TargetTrackingScaling`
	// - `StepScaling`
	// - `SimpleScaling` (default)
	// - `PredictiveScaling`.
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// A predictive scaling policy. Provides support for predefined and custom metrics.
	//
	// Predefined metrics include CPU utilization, network in/out, and the Application Load Balancer request count.
	//
	// Required if the policy type is `PredictiveScaling` .
	PredictiveScalingConfiguration interface{} `field:"optional" json:"predictiveScalingConfiguration" yaml:"predictiveScalingConfiguration"`
	// The amount by which to scale, based on the specified adjustment type.
	//
	// A positive value adds to the current capacity while a negative number removes from the current capacity. For exact capacity, you must specify a positive value.
	//
	// Required if the policy type is `SimpleScaling` . (Not used with any other policy type.)
	ScalingAdjustment *float64 `field:"optional" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// Required if the policy type is `StepScaling` . (Not used with any other policy type.)
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
	TargetTrackingConfiguration interface{} `field:"optional" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
}

