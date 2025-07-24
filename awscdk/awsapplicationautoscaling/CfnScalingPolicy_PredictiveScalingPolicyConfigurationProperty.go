package awsapplicationautoscaling


// Represents a predictive scaling policy configuration.
//
// Predictive scaling is supported on Amazon ECS services.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingPolicyConfigurationProperty := &PredictiveScalingPolicyConfigurationProperty{
//   	MetricSpecifications: []interface{}{
//   		&PredictiveScalingMetricSpecificationProperty{
//   			TargetValue: jsii.Number(123),
//
//   			// the properties below are optional
//   			CustomizedCapacityMetricSpecification: &PredictiveScalingCustomizedCapacityMetricProperty{
//   				MetricDataQueries: []interface{}{
//   					&PredictiveScalingMetricDataQueryProperty{
//   						Expression: jsii.String("expression"),
//   						Id: jsii.String("id"),
//   						Label: jsii.String("label"),
//   						MetricStat: &PredictiveScalingMetricStatProperty{
//   							Metric: &PredictiveScalingMetricProperty{
//   								Dimensions: []interface{}{
//   									&PredictiveScalingMetricDimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Stat: jsii.String("stat"),
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			CustomizedLoadMetricSpecification: &PredictiveScalingCustomizedLoadMetricProperty{
//   				MetricDataQueries: []interface{}{
//   					&PredictiveScalingMetricDataQueryProperty{
//   						Expression: jsii.String("expression"),
//   						Id: jsii.String("id"),
//   						Label: jsii.String("label"),
//   						MetricStat: &PredictiveScalingMetricStatProperty{
//   							Metric: &PredictiveScalingMetricProperty{
//   								Dimensions: []interface{}{
//   									&PredictiveScalingMetricDimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Stat: jsii.String("stat"),
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			CustomizedScalingMetricSpecification: &PredictiveScalingCustomizedScalingMetricProperty{
//   				MetricDataQueries: []interface{}{
//   					&PredictiveScalingMetricDataQueryProperty{
//   						Expression: jsii.String("expression"),
//   						Id: jsii.String("id"),
//   						Label: jsii.String("label"),
//   						MetricStat: &PredictiveScalingMetricStatProperty{
//   							Metric: &PredictiveScalingMetricProperty{
//   								Dimensions: []interface{}{
//   									&PredictiveScalingMetricDimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Stat: jsii.String("stat"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html
//
type CfnScalingPolicy_PredictiveScalingPolicyConfigurationProperty struct {
	// This structure includes the metrics and target utilization to use for predictive scaling.
	//
	// This is an array, but we currently only support a single metric specification. That is, you can specify a target value and a single metric pair, or a target value and one scaling metric and one load metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-metricspecifications
	//
	MetricSpecifications interface{} `field:"required" json:"metricSpecifications" yaml:"metricSpecifications"`
	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity.
	//
	// Defaults to `HonorMaxCapacity` if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-maxcapacitybreachbehavior
	//
	MaxCapacityBreachBehavior *string `field:"optional" json:"maxCapacityBreachBehavior" yaml:"maxCapacityBreachBehavior"`
	// The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	//
	// The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer, such that if the forecast capacity is 50, and the maximum capacity is 40, then the effective maximum capacity is 55.
	//
	// Required if the `MaxCapacityBreachBehavior` property is set to `IncreaseMaxCapacity` , and cannot be used otherwise.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-maxcapacitybuffer
	//
	MaxCapacityBuffer *float64 `field:"optional" json:"maxCapacityBuffer" yaml:"maxCapacityBuffer"`
	// The predictive scaling mode.
	//
	// Defaults to `ForecastOnly` if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The amount of time, in seconds, that the start time can be advanced.
	//
	// The value must be less than the forecast interval duration of 3600 seconds (60 minutes). Defaults to 300 seconds if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-schedulingbuffertime
	//
	SchedulingBufferTime *float64 `field:"optional" json:"schedulingBufferTime" yaml:"schedulingBufferTime"`
}

