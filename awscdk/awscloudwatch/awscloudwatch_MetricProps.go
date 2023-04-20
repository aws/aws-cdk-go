package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for a metric.
//
// Example:
//   hostedZone := route53.NewHostedZone(this, jsii.String("MyHostedZone"), &HostedZoneProps{
//   	ZoneName: jsii.String("example.org"),
//   })
//   metric := cloudwatch.NewMetric(&MetricProps{
//   	Namespace: jsii.String("AWS/Route53"),
//   	MetricName: jsii.String("DNSQueries"),
//   	DimensionsMap: map[string]*string{
//   		"HostedZoneId": hostedZone.hostedZoneId,
//   	},
//   })
//
// Experimental.
type MetricProps struct {
	// Account which this metric comes from.
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to use when this metric is rendered on a graph. The `Color` class has a set of standard colors that can be used here.
	// Experimental.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Dimensions of the metric.
	// Deprecated: Use 'dimensionsMap' instead.
	Dimensions *map[string]interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Dimensions of the metric.
	// Experimental.
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
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The period over which the specified statistic is applied.
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Region which this metric comes from.
	// Experimental.
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
	// Experimental.
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
	// Experimental.
	Unit Unit `field:"optional" json:"unit" yaml:"unit"`
	// Name of the metric.
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Namespace of the metric.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

