package awsapplicationautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnScalingPolicy_PredictiveScalingMetricStatProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricstat-metric
	//
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// The statistic to return.
	//
	// It can include any CloudWatch statistic or extended statistic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricstat-stat
	//
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// The unit to use for the returned data points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricstat.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricstat-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

