package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties needed to make an anomaly detection alarm from a metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metric Metric
//
//   anomalyDetectionMetricOptions := &AnomalyDetectionMetricOptions{
//   	Metric: metric,
//
//   	// the properties below are optional
//   	Color: jsii.String("color"),
//   	Label: jsii.String("label"),
//   	Period: cdk.Duration_Minutes(jsii.Number(30)),
//   	SearchAccount: jsii.String("searchAccount"),
//   	SearchRegion: jsii.String("searchRegion"),
//   	StdDevs: jsii.Number(123),
//   }
//
type AnomalyDetectionMetricOptions struct {
	// Color for this metric when added to a Graph in a Dashboard.
	// Default: - Automatic color.
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Label for this expression when added to a Graph in a Dashboard.
	//
	// If this expression evaluates to more than one time series (for
	// example, through the use of `METRICS()` or `SEARCH()` expressions),
	// each time series will appear in the graph using a combination of the
	// expression label and the individual metric label. Specify the empty
	// string (`''`) to suppress the expression label and only keep the
	// metric label.
	//
	// You can use [dynamic labels](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html)
	// to show summary information about the displayed time series
	// in the legend. For example, if you use:
	//
	// ```
	// [max: ${MAX}] MyMetric
	// ```
	//
	// As the metric label, the maximum value in the visible range will
	// be shown next to the time series name in the graph's legend. If the
	// math expression produces more than one time series, the maximum
	// will be shown for each individual time series produce by this
	// math expression.
	// Default: - Expression value is used as label.
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The period over which the math expression's statistics are applied.
	//
	// This period overrides all periods in the metrics used in this
	// math expression.
	// Default: Duration.minutes(5)
	//
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Account to evaluate search expressions within.
	//
	// Specifying a searchAccount has no effect to the account used
	// for metrics within the expression (passed via usingMetrics).
	// Default: - Deployment account.
	//
	SearchAccount *string `field:"optional" json:"searchAccount" yaml:"searchAccount"`
	// Region to evaluate search expressions within.
	//
	// Specifying a searchRegion has no effect to the region used
	// for metrics within the expression (passed via usingMetrics).
	// Default: - Deployment region.
	//
	SearchRegion *string `field:"optional" json:"searchRegion" yaml:"searchRegion"`
	// The metric to add the alarm on.
	//
	// Metric objects can be obtained from most resources, or you can construct
	// custom Metric objects by instantiating one.
	Metric IMetric `field:"required" json:"metric" yaml:"metric"`
	// The number of standard deviations to use for the anomaly detection band.
	//
	// The higher the value, the wider the band.
	//
	// - Must be greater than 0. A value of 0 or negative values would not make sense in the context of calculating standard deviations.
	// - There is no strict maximum value defined, as standard deviations can theoretically extend infinitely. However, in practice, values beyond 5 or 6 standard deviations are rarely used, as they would result in an extremely wide anomaly detection band, potentially missing significant anomalies.
	// Default: 2.
	//
	StdDevs *float64 `field:"optional" json:"stdDevs" yaml:"stdDevs"`
}

