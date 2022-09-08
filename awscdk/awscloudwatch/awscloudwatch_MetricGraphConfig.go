package awscloudwatch


// Properties used to construct the Metric identifying part of a Graph.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   metricGraphConfig := &metricGraphConfig{
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//   	period: jsii.Number(123),
//   	renderingProperties: &metricRenderingProperties{
//   		period: jsii.Number(123),
//
//   		// the properties below are optional
//   		color: jsii.String("color"),
//   		label: jsii.String("label"),
//   		stat: jsii.String("stat"),
//   	},
//
//   	// the properties below are optional
//   	color: jsii.String("color"),
//   	dimensions: []dimension{
//   		&dimension{
//   			name: jsii.String("name"),
//   			value: value,
//   		},
//   	},
//   	label: jsii.String("label"),
//   	statistic: jsii.String("statistic"),
//   	unit: awscdk.Aws_cloudwatch.unit_SECONDS,
//   }
//
// Deprecated: Replaced by MetricConfig.
type MetricGraphConfig struct {
	// Name of the metric.
	// Deprecated: Replaced by MetricConfig.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Namespace of the metric.
	// Deprecated: Replaced by MetricConfig.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// How many seconds to aggregate over.
	// Deprecated: Use `period` in `renderingProperties`.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Rendering properties override yAxis parameter of the widget object.
	// Deprecated: Replaced by MetricConfig.
	RenderingProperties *MetricRenderingProperties `field:"required" json:"renderingProperties" yaml:"renderingProperties"`
	// Color for the graph line.
	// Deprecated: Use `color` in `renderingProperties`.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The dimensions to apply to the alarm.
	// Deprecated: Replaced by MetricConfig.
	Dimensions *[]*Dimension `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Label for the metric.
	// Deprecated: Use `label` in `renderingProperties`.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Aggregation function to use (can be either simple or a percentile).
	// Deprecated: Use `stat` in `renderingProperties`.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of the alarm.
	// Deprecated: not used in dashboard widgets.
	Unit Unit `field:"optional" json:"unit" yaml:"unit"`
}

