package previewawsapplicationautoscalingmixins


// This structure defines the CloudWatch metric to return, along with the statistic and unit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   predictiveScalingMetricStatProperty := &PredictiveScalingMetricStatProperty{
//   	Metric: &PredictiveScalingMetricProperty{
//   		Dimensions: []interface{}{
//   			&PredictiveScalingMetricDimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   	},
//   	Stat: jsii.String("stat"),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricstat.html
//
type CfnScalingPolicyPropsMixin_PredictiveScalingMetricStatProperty struct {
	// The CloudWatch metric to return, including the metric name, namespace, and dimensions.
	//
	// To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricstat-metric
	//
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// The statistic to return.
	//
	// It can include any CloudWatch statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *Amazon CloudWatch User Guide* .
	//
	// The most commonly used metrics for predictive scaling are `Average` and `Sum` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricstat-stat
	//
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// The unit to use for the returned data points.
	//
	// For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricstat-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

