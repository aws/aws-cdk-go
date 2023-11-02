package awsautoscaling


// `PredictiveScalingConfiguration` is a property of the [AWS::AutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html) resource that specifies a predictive scaling policy for Amazon EC2 Auto Scaling.
//
// For more information, see [Predictive scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingConfigurationProperty := &PredictiveScalingConfigurationProperty{
//   	MetricSpecifications: []interface{}{
//   		&PredictiveScalingMetricSpecificationProperty{
//   			TargetValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			CustomizedCapacityMetricSpecification: &PredictiveScalingCustomizedCapacityMetricProperty{
//   				MetricDataQueries: []interface{}{
//   					&MetricDataQueryProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						Expression: jsii.String("expression"),
//   						Label: jsii.String("label"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//
//   								// the properties below are optional
//   								Dimensions: []interface{}{
//   									&MetricDimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							Stat: jsii.String("stat"),
//
//   							// the properties below are optional
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			CustomizedLoadMetricSpecification: &PredictiveScalingCustomizedLoadMetricProperty{
//   				MetricDataQueries: []interface{}{
//   					&MetricDataQueryProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						Expression: jsii.String("expression"),
//   						Label: jsii.String("label"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//
//   								// the properties below are optional
//   								Dimensions: []interface{}{
//   									&MetricDimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							Stat: jsii.String("stat"),
//
//   							// the properties below are optional
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			CustomizedScalingMetricSpecification: &PredictiveScalingCustomizedScalingMetricProperty{
//   				MetricDataQueries: []interface{}{
//   					&MetricDataQueryProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						Expression: jsii.String("expression"),
//   						Label: jsii.String("label"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//
//   								// the properties below are optional
//   								Dimensions: []interface{}{
//   									&MetricDimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   							},
//   							Stat: jsii.String("stat"),
//
//   							// the properties below are optional
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			PredefinedLoadMetricSpecification: &PredictiveScalingPredefinedLoadMetricProperty{
//   				PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   				// the properties below are optional
//   				ResourceLabel: jsii.String("resourceLabel"),
//   			},
//   			PredefinedMetricPairSpecification: &PredictiveScalingPredefinedMetricPairProperty{
//   				PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   				// the properties below are optional
//   				ResourceLabel: jsii.String("resourceLabel"),
//   			},
//   			PredefinedScalingMetricSpecification: &PredictiveScalingPredefinedScalingMetricProperty{
//   				PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   				// the properties below are optional
//   				ResourceLabel: jsii.String("resourceLabel"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	MaxCapacityBreachBehavior: jsii.String("maxCapacityBreachBehavior"),
//   	MaxCapacityBuffer: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   	SchedulingBufferTime: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.html
//
type CfnScalingPolicy_PredictiveScalingConfigurationProperty struct {
	// This structure includes the metrics and target utilization to use for predictive scaling.
	//
	// This is an array, but we currently only support a single metric specification. That is, you can specify a target value and a single metric pair, or a target value and one scaling metric and one load metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.html#cfn-autoscaling-scalingpolicy-predictivescalingconfiguration-metricspecifications
	//
	MetricSpecifications interface{} `field:"required" json:"metricSpecifications" yaml:"metricSpecifications"`
	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity of the Auto Scaling group.
	//
	// Defaults to `HonorMaxCapacity` if not specified.
	//
	// The following are possible values:
	//
	// - `HonorMaxCapacity` - Amazon EC2 Auto Scaling cannot scale out capacity higher than the maximum capacity. The maximum capacity is enforced as a hard limit.
	// - `IncreaseMaxCapacity` - Amazon EC2 Auto Scaling can scale out capacity higher than the maximum capacity when the forecast capacity is close to or exceeds the maximum capacity. The upper limit is determined by the forecasted capacity and the value for `MaxCapacityBuffer` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.html#cfn-autoscaling-scalingpolicy-predictivescalingconfiguration-maxcapacitybreachbehavior
	//
	MaxCapacityBreachBehavior *string `field:"optional" json:"maxCapacityBreachBehavior" yaml:"maxCapacityBreachBehavior"`
	// The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	//
	// The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer, such that if the forecast capacity is 50, and the maximum capacity is 40, then the effective maximum capacity is 55.
	//
	// If set to 0, Amazon EC2 Auto Scaling may scale capacity higher than the maximum capacity to equal but not exceed forecast capacity.
	//
	// Required if the `MaxCapacityBreachBehavior` property is set to `IncreaseMaxCapacity` , and cannot be used otherwise.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.html#cfn-autoscaling-scalingpolicy-predictivescalingconfiguration-maxcapacitybuffer
	//
	MaxCapacityBuffer *float64 `field:"optional" json:"maxCapacityBuffer" yaml:"maxCapacityBuffer"`
	// The predictive scaling mode.
	//
	// Defaults to `ForecastOnly` if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.html#cfn-autoscaling-scalingpolicy-predictivescalingconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The amount of time, in seconds, by which the instance launch time can be advanced.
	//
	// For example, the forecast says to add capacity at 10:00 AM, and you choose to pre-launch instances by 5 minutes. In that case, the instances will be launched at 9:55 AM. The intention is to give resources time to be provisioned. It can take a few minutes to launch an EC2 instance. The actual amount of time required depends on several factors, such as the size of the instance and whether there are startup scripts to complete.
	//
	// The value must be less than the forecast interval duration of 3600 seconds (60 minutes). Defaults to 300 seconds if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predictivescalingconfiguration.html#cfn-autoscaling-scalingpolicy-predictivescalingconfiguration-schedulingbuffertime
	//
	SchedulingBufferTime *float64 `field:"optional" json:"schedulingBufferTime" yaml:"schedulingBufferTime"`
}

