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
//   metricExpressionConfig := &MetricExpressionConfig{
//   	Expression: jsii.String("expression"),
//   	Period: jsii.Number(123),
//   	UsingMetrics: map[string]iMetric{
//   		"usingMetricsKey": metric,
//   	},
//
//   	// the properties below are optional
//   	SearchAccount: jsii.String("searchAccount"),
//   	SearchRegion: jsii.String("searchRegion"),
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
	// Default: - Deployment account.
	//
	SearchAccount *string `field:"optional" json:"searchAccount" yaml:"searchAccount"`
	// Region to evaluate search expressions within.
	// Default: - Deployment region.
	//
	SearchRegion *string `field:"optional" json:"searchRegion" yaml:"searchRegion"`
}

