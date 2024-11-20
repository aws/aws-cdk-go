package awsapplicationautoscaling


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-metricspecifications
	//
	MetricSpecifications interface{} `field:"required" json:"metricSpecifications" yaml:"metricSpecifications"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-maxcapacitybreachbehavior
	//
	MaxCapacityBreachBehavior *string `field:"optional" json:"maxCapacityBreachBehavior" yaml:"maxCapacityBreachBehavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-maxcapacitybuffer
	//
	MaxCapacityBuffer *float64 `field:"optional" json:"maxCapacityBuffer" yaml:"maxCapacityBuffer"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration-schedulingbuffertime
	//
	SchedulingBufferTime *float64 `field:"optional" json:"schedulingBufferTime" yaml:"schedulingBufferTime"`
}

