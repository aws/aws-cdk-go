package awscloudwatch


// This structure defines the metric to be returned, along with the statistics, period, and units.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricStatProperty := &MetricStatProperty{
//   	Metric: &MetricProperty{
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//
//   		// the properties below are optional
//   		Dimensions: []interface{}{
//   			&DimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Period: jsii.Number(123),
//   	Stat: jsii.String("stat"),
//
//   	// the properties below are optional
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html
//
type CfnAnomalyDetector_MetricStatProperty struct {
	// The metric to return, including the metric name, namespace, and dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html#cfn-cloudwatch-anomalydetector-metricstat-metric
	//
	Metric interface{} `field:"required" json:"metric" yaml:"metric"`
	// The granularity, in seconds, of the returned data points.
	//
	// For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 20, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a `PutMetricData` call that includes a `StorageResolution` of 1 second.
	//
	// If the `StartTime` parameter specifies a time stamp that is greater than 3 hours ago, you must specify the period as follows or no data points in that time range is returned:
	//
	// - Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds (1 minute).
	// - Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5 minutes).
	// - Start time greater than 63 days ago - Use a multiple of 3600 seconds (1 hour).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html#cfn-cloudwatch-anomalydetector-metricstat-period
	//
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The statistic to return.
	//
	// It can include any CloudWatch statistic or extended statistic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html#cfn-cloudwatch-anomalydetector-metricstat-stat
	//
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// When you are using a `Put` operation, this defines what unit you want to use when storing the metric.
	//
	// In a `Get` operation, if you omit `Unit` then all data that was collected with any unit is returned, along with the corresponding units that were specified when the data was reported to CloudWatch. If you specify a unit, the operation returns only data that was collected with that unit specified. If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-anomalydetector-metricstat.html#cfn-cloudwatch-anomalydetector-metricstat-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

