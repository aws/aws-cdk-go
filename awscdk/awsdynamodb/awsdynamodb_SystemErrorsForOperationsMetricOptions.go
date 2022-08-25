package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Options for configuring a system errors metric that considers multiple operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemErrorsForOperationsMetricOptions := &systemErrorsForOperationsMetricOptions{
//   	account: jsii.String("account"),
//   	color: jsii.String("color"),
//   	dimensionsMap: map[string]*string{
//   		"dimensionsMapKey": jsii.String("dimensionsMap"),
//   	},
//   	label: jsii.String("label"),
//   	operations: []operation{
//   		awscdk.Aws_dynamodb.*operation_GET_ITEM,
//   	},
//   	period: cdk.duration.minutes(jsii.Number(30)),
//   	region: jsii.String("region"),
//   	statistic: jsii.String("statistic"),
//   	unit: awscdk.Aws_cloudwatch.unit_SECONDS,
//   }
//
type SystemErrorsForOperationsMetricOptions struct {
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
	Unit awscloudwatch.Unit `field:"optional" json:"unit" yaml:"unit"`
	// The operations to apply the metric to.
	Operations *[]Operation `field:"optional" json:"operations" yaml:"operations"`
}

