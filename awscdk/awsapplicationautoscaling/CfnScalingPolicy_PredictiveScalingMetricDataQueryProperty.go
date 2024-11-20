package awsapplicationautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveScalingMetricDataQueryProperty := &PredictiveScalingMetricDataQueryProperty{
//   	Expression: jsii.String("expression"),
//   	Id: jsii.String("id"),
//   	Label: jsii.String("label"),
//   	MetricStat: &PredictiveScalingMetricStatProperty{
//   		Metric: &PredictiveScalingMetricProperty{
//   			Dimensions: []interface{}{
//   				&PredictiveScalingMetricDimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//   		},
//   		Stat: jsii.String("stat"),
//   		Unit: jsii.String("unit"),
//   	},
//   	ReturnData: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html
//
type CfnScalingPolicy_PredictiveScalingMetricDataQueryProperty struct {
	// The math expression to perform on the returned data, if this object is performing a math expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A short name that identifies the object's results in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A human-readable label for this metric or expression.
	//
	// This is especially useful if this is a math expression, so that you know what the value represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-metricstat
	//
	MetricStat interface{} `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Indicates whether to return the timestamps and raw data values of this metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-returndata
	//
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

