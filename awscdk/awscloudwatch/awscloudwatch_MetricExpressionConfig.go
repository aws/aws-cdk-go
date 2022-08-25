package awscloudwatch


// Properties for a concrete metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metric metric
//
//   metricExpressionConfig := &metricExpressionConfig{
//   	expression: jsii.String("expression"),
//   	period: jsii.Number(123),
//   	usingMetrics: map[string]iMetric{
//   		"usingMetricsKey": metric,
//   	},
//
//   	// the properties below are optional
//   	searchAccount: jsii.String("searchAccount"),
//   	searchRegion: jsii.String("searchRegion"),
//   }
//
type MetricExpressionConfig struct {
	// Math expression for the metric.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// How many seconds to aggregate over.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Metrics used in the math expression.
	UsingMetrics *map[string]IMetric `field:"required" json:"usingMetrics" yaml:"usingMetrics"`
	// Account to evaluate search expressions within.
	SearchAccount *string `field:"optional" json:"searchAccount" yaml:"searchAccount"`
	// Region to evaluate search expressions within.
	SearchRegion *string `field:"optional" json:"searchRegion" yaml:"searchRegion"`
}

