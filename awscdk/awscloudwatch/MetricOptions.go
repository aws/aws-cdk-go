package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties of a metric that can be changed.
//
// Example:
//   var matchmakingRuleSet matchmakingRuleSet
//
//   // Alarm that triggers when the per-second average of not placed matches exceed 10%
//   ruleEvaluationRatio := cloudwatch.NewMathExpression(&MathExpressionProps{
//   	Expression: jsii.String("1 - (ruleEvaluationsPassed / ruleEvaluationsFailed)"),
//   	UsingMetrics: map[string]iMetric{
//   		"ruleEvaluationsPassed": matchmakingRuleSet.metricRuleEvaluationsPassed(&MetricOptions{
//   			"statistic": cloudwatch.Statistic_SUM,
//   		}),
//   		"ruleEvaluationsFailed": matchmakingRuleSet.metric(jsii.String("ruleEvaluationsFailed")),
//   	},
//   })
//   cloudwatch.NewAlarm(this, jsii.String("Alarm"), &AlarmProps{
//   	Metric: ruleEvaluationRatio,
//   	Threshold: jsii.Number(0.1),
//   	EvaluationPeriods: jsii.Number(3),
//   })
//
type MetricOptions struct {
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
}

