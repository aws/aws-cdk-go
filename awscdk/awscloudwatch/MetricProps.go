package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a metric.
//
// Example:
//   // Create a metric
//   metric := cloudwatch.NewMetric(&MetricProps{
//   	Namespace: jsii.String("AWS/EC2"),
//   	MetricName: jsii.String("CPUUtilization"),
//   	Statistic: jsii.String("Average"),
//   	Period: awscdk.Duration_Minutes(jsii.Number(5)),
//   })
//
//   // Create an anomaly detection alarm
//   alarm := cloudwatch.NewAnomalyDetectionAlarm(this, jsii.String("AnomalyAlarm"), &AnomalyDetectionAlarmProps{
//   	Metric: metric,
//   	EvaluationPeriods: jsii.Number(1),
//
//   	// Number of standard deviations for the band (default: 2)
//   	StdDevs: jsii.Number(2),
//   	// Alarm outside on either side of the band, or just below or above it (default: outside)
//   	ComparisonOperator: cloudwatch.ComparisonOperator_LESS_THAN_LOWER_OR_GREATER_THAN_UPPER_THRESHOLD,
//   	AlarmDescription: jsii.String("Alarm when metric is outside the expected band"),
//   })
//
type MetricProps struct {
	// Account which this metric comes from.
	// Default: - Deployment account.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to use when this metric is rendered on a graph. The `Color` class has a set of standard colors that can be used here.
	// Default: - Automatic color.
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Dimensions of the metric.
	// Default: - No dimensions.
	//
	DimensionsMap *map[string]*string `field:"optional" json:"dimensionsMap" yaml:"dimensionsMap"`
	// Label for this metric when added to a Graph in a Dashboard.
	//
	// You can use [dynamic labels](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html)
	// to show summary information about the entire displayed time series
	// in the legend. For example, if you use:
	//
	// ```
	// [max: ${MAX}] MyMetric
	// ```
	//
	// As the metric label, the maximum value in the visible range will
	// be shown next to the time series name in the graph's legend.
	// Default: - No label.
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The period over which the specified statistic is applied.
	// Default: Duration.minutes(5)
	//
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Region which this metric comes from.
	// Default: - Deployment region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Account of the stack this metric is attached to.
	// Default: - Deployment account.
	//
	StackAccount *string `field:"optional" json:"stackAccount" yaml:"stackAccount"`
	// Region of the stack this metric is attached to.
	// Default: - Deployment region.
	//
	StackRegion *string `field:"optional" json:"stackRegion" yaml:"stackRegion"`
	// What function to use for aggregating.
	//
	// Use the `aws_cloudwatch.Stats` helper class to construct valid input strings.
	//
	// Can be one of the following:
	//
	// - "Minimum" | "min"
	// - "Maximum" | "max"
	// - "Average" | "avg"
	// - "Sum" | "sum"
	// - "SampleCount | "n"
	// - "pNN.NN"
	// - "tmNN.NN" | "tm(NN.NN%:NN.NN%)"
	// - "iqm"
	// - "wmNN.NN" | "wm(NN.NN%:NN.NN%)"
	// - "tcNN.NN" | "tc(NN.NN%:NN.NN%)"
	// - "tsNN.NN" | "ts(NN.NN%:NN.NN%)"
	// Default: Average.
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Unit used to filter the metric stream.
	//
	// Only refer to datums emitted to the metric stream with the given unit and
	// ignore all others. Only useful when datums are being emitted to the same
	// metric stream under different units.
	//
	// The default is to use all matric datums in the stream, regardless of unit,
	// which is recommended in nearly all cases.
	//
	// CloudWatch does not honor this property for graphs.
	// Default: - All metric datums in the given metric stream.
	//
	Unit Unit `field:"optional" json:"unit" yaml:"unit"`
	// Name of the metric.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Namespace of the metric.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

