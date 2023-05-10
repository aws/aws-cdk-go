package awscloudwatch


// Custom rendering properties that override the default rendering properties specified in the yAxis parameter of the widget object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricRenderingProperties := &MetricRenderingProperties{
//   	Period: jsii.Number(123),
//
//   	// the properties below are optional
//   	Color: jsii.String("color"),
//   	Label: jsii.String("label"),
//   	Stat: jsii.String("stat"),
//   }
//
// Deprecated: Replaced by MetricConfig.
type MetricRenderingProperties struct {
	// How many seconds to aggregate over.
	// Deprecated: Replaced by MetricConfig.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to use when this metric is rendered on a graph. The `Color` class has a set of standard colors that can be used here.
	// Deprecated: Replaced by MetricConfig.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Label for the metric.
	// Deprecated: Replaced by MetricConfig.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Aggregation function to use (can be either simple or a percentile).
	// Deprecated: Replaced by MetricConfig.
	Stat *string `field:"optional" json:"stat" yaml:"stat"`
}

