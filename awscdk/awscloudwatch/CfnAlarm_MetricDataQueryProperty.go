package awscloudwatch


// The `MetricDataQuery` property type specifies the metric data to return, and whether this call is just retrieving a batch set of data for one metric, or is performing a math expression on metric data.
//
// Any expression used must return a single time series. For more information, see [Metric Math Syntax and Functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax) in the *Amazon CloudWatch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDataQueryProperty := &MetricDataQueryProperty{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	AccountId: jsii.String("accountId"),
//   	Expression: jsii.String("expression"),
//   	Label: jsii.String("label"),
//   	MetricStat: &MetricStatProperty{
//   		Metric: &MetricProperty{
//   			Dimensions: []interface{}{
//   				&DimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//   		},
//   		Period: jsii.Number(123),
//   		Stat: jsii.String("stat"),
//
//   		// the properties below are optional
//   		Unit: jsii.String("unit"),
//   	},
//   	Period: jsii.Number(123),
//   	ReturnData: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html
//
type CfnAlarm_MetricDataQueryProperty struct {
	// A short name used to tie this object to the results in the response.
	//
	// This name must be unique within a single call to `GetMetricData` . If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html#cfn-cloudwatch-alarm-metricdataquery-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html#cfn-cloudwatch-alarm-metricdataquery-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The math expression to be performed on the returned data, if this object is performing a math expression.
	//
	// This expression can use the `Id` of the other metrics to refer to those metrics, and can also use the `Id` of other expressions to use the result of those expressions. For more information about metric math expressions, see [Metric Math Syntax and Functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax) in the *Amazon CloudWatch User Guide* .
	//
	// Within each MetricDataQuery object, you must specify either `Expression` or `MetricStat` but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html#cfn-cloudwatch-alarm-metricdataquery-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A human-readable label for this metric or expression.
	//
	// This is especially useful if this is an expression, so that you know what the value represents. If the metric or expression is shown in a CloudWatch dashboard widget, the label is shown. If `Label` is omitted, CloudWatch generates a default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html#cfn-cloudwatch-alarm-metricdataquery-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The metric to be returned, along with statistics, period, and units.
	//
	// Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
	//
	// Within one MetricDataQuery object, you must specify either `Expression` or `MetricStat` but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html#cfn-cloudwatch-alarm-metricdataquery-metricstat
	//
	MetricStat interface{} `field:"optional" json:"metricStat" yaml:"metricStat"`
	// The granularity, in seconds, of the returned data points.
	//
	// For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 20, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a `PutMetricData` operation that includes a `StorageResolution of 1 second` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html#cfn-cloudwatch-alarm-metricdataquery-period
	//
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// This option indicates whether to return the timestamps and raw data values of this metric.
	//
	// When you create an alarm based on a metric math expression, specify `True` for this value for only the one math expression that the alarm is based on. You must specify `False` for `ReturnData` for all the other metrics and expressions used in the alarm.
	//
	// This field is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-metricdataquery.html#cfn-cloudwatch-alarm-metricdataquery-returndata
	//
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

