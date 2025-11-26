package previewawsapplicationsignalsmixins


// This structure defines the metric to be used as the service level indicator, along with the statistics, period, and unit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricStatProperty := &MetricStatProperty{
//   	Metric: &MetricProperty{
//   		Dimensions: []interface{}{
//   			&DimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   	},
//   	Period: jsii.Number(123),
//   	Stat: jsii.String("stat"),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html
//
type CfnServiceLevelObjectivePropsMixin_MetricStatProperty struct {
	// The metric to use as the service level indicator, including the metric name, namespace, and dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html#cfn-applicationsignals-servicelevelobjective-metricstat-metric
	//
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// The granularity, in seconds, to be used for the metric.
	//
	// For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a `PutMetricData` call that includes a `StorageResolution` of 1 second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html#cfn-applicationsignals-servicelevelobjective-metricstat-period
	//
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic to use for comparison to the threshold.
	//
	// It can be any CloudWatch statistic or extended statistic. For more information about statistics, see [CloudWatch statistics definitions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html#cfn-applicationsignals-servicelevelobjective-metricstat-stat
	//
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
	// If you omit `Unit` then all data that was collected with any unit is returned, along with the corresponding units that were specified when the data was reported to CloudWatch.
	//
	// If you specify a unit, the operation returns only data that was collected with that unit specified. If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-metricstat.html#cfn-applicationsignals-servicelevelobjective-metricstat-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

