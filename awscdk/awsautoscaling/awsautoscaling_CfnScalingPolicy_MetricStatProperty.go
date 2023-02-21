package awsautoscaling


// `MetricStat` is a property of the [AWS::AutoScaling::ScalingPolicy MetricDataQuery](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricdataquery.html) property type.
//
// This structure defines the CloudWatch metric to return, along with the statistic, period, and unit.
//
// For more information about the CloudWatch terminology below, see [Amazon CloudWatch concepts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html) in the *Amazon CloudWatch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStatProperty := &metricStatProperty{
//   	metric: &metricProperty{
//   		metricName: jsii.String("metricName"),
//   		namespace: jsii.String("namespace"),
//
//   		// the properties below are optional
//   		dimensions: []interface{}{
//   			&metricDimensionProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	stat: jsii.String("stat"),
//
//   	// the properties below are optional
//   	unit: jsii.String("unit"),
//   }
//
type CfnScalingPolicy_MetricStatProperty struct {
	// The CloudWatch metric to return, including the metric name, namespace, and dimensions.
	//
	// To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html) .
	Metric interface{} `field:"required" json:"metric" yaml:"metric"`
	// The statistic to return.
	//
	// It can include any CloudWatch statistic or extended statistic. For a list of valid values, see the table in [Statistics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html#Statistic) in the *Amazon CloudWatch User Guide* .
	//
	// The most commonly used metrics for predictive scaling are `Average` and `Sum` .
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// The unit to use for the returned data points.
	//
	// For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference* .
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

