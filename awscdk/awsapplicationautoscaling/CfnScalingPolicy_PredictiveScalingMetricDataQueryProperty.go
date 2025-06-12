package awsapplicationautoscaling


// The metric data to return.
//
// Also defines whether this call is returning data for one metric only, or whether it is performing a math expression on the values of returned metric statistics to create a new time series. A time series is a series of data points, each of which is associated with a timestamp.
//
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
	//
	// This expression can use the `Id` of the other metrics to refer to those metrics, and can also use the `Id` of other expressions to use the result of those expressions.
	//
	// Conditional: Within each `MetricDataQuery` object, you must specify either `Expression` or `MetricStat` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A short name that identifies the object's results in the response.
	//
	// This name must be unique among all `MetricDataQuery` objects specified for a single scaling policy. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscores. The first character must be a lowercase letter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A human-readable label for this metric or expression.
	//
	// This is especially useful if this is a math expression, so that you know what the value represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Information about the metric data to return.
	//
	// Conditional: Within each `MetricDataQuery` object, you must specify either `Expression` or `MetricStat` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-metricstat
	//
	MetricStat interface{} `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Indicates whether to return the timestamps and raw data values of this metric.
	//
	// If you use any math expressions, specify `true` for this value for only the final math expression that the metric specification is based on. You must specify `false` for `ReturnData` for all the other metrics and expressions used in the metric specification.
	//
	// If you are only retrieving metrics and not performing any math expressions, do not specify anything for `ReturnData` . This sets it to its default ( `true` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.html#cfn-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery-returndata
	//
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

