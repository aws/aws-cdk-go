package awsstepfunctions


// Task Metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metricDimensions interface{}
//
//   taskMetricsConfig := &TaskMetricsConfig{
//   	MetricDimensions: map[string]interface{}{
//   		"metricDimensionsKey": metricDimensions,
//   	},
//   	MetricPrefixPlural: jsii.String("metricPrefixPlural"),
//   	MetricPrefixSingular: jsii.String("metricPrefixSingular"),
//   }
//
type TaskMetricsConfig struct {
	// The dimensions to attach to metrics.
	// Default: - No metrics.
	//
	MetricDimensions *map[string]interface{} `field:"optional" json:"metricDimensions" yaml:"metricDimensions"`
	// Prefix for plural metric names of activity actions.
	// Default: - No such metrics.
	//
	MetricPrefixPlural *string `field:"optional" json:"metricPrefixPlural" yaml:"metricPrefixPlural"`
	// Prefix for singular metric names of activity actions.
	// Default: - No such metrics.
	//
	MetricPrefixSingular *string `field:"optional" json:"metricPrefixSingular" yaml:"metricPrefixSingular"`
}

