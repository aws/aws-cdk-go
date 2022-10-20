package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a metric.
//
// Example:
//   hostedZone := route53.NewHostedZone(this, jsii.String("MyHostedZone"), &hostedZoneProps{
//   	zoneName: jsii.String("example.org"),
//   })
//   metric := cloudwatch.NewMetric(&metricProps{
//   	namespace: jsii.String("AWS/Route53"),
//   	metricName: jsii.String("DNSQueries"),
//   	dimensionsMap: map[string]*string{
//   		"HostedZoneId": hostedZone.hostedZoneId,
//   	},
//   })
//
type MetricProps struct {
	// Account which this metric comes from.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to use when this metric is rendered on a graph. The `Color` class has a set of standard colors that can be used here.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Dimensions of the metric.
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
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The period over which the specified statistic is applied.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Region which this metric comes from.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// What function to use for aggregating.
	//
	// Can be one of the following:
	//
	// - "Minimum" | "min"
	// - "Maximum" | "max"
	// - "Average" | "avg"
	// - "Sum" | "sum"
	// - "SampleCount | "n"
	// - "pNN.NN"
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
	Unit Unit `field:"optional" json:"unit" yaml:"unit"`
	// Name of the metric.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Namespace of the metric.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

