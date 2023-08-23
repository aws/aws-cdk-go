package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a metric.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   metric := cloudwatch.NewMetric(&MetricProps{
//   	Namespace: jsii.String("MyNamespace"),
//   	MetricName: jsii.String("MyMetric"),
//   	DimensionsMap: map[string]*string{
//   		"MyDimension": jsii.String("MyDimensionValue"),
//   	},
//   })
//   alarm := cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &AlarmProps{
//   	Metric: metric,
//   	Threshold: jsii.Number(100),
//   	EvaluationPeriods: jsii.Number(3),
//   	DatapointsToAlarm: jsii.Number(2),
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
//   	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
//   	Actions: []iAction{
//   		actions.NewCloudWatchSetAlarmStateAction(alarm, &CloudWatchSetAlarmStateActionProps{
//   			Reason: jsii.String("AWS Iot Rule action is triggered"),
//   			AlarmStateToSet: cloudwatch.AlarmState_ALARM,
//   		}),
//   	},
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

