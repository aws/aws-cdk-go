package awscloudwatch


// Properties of a rendered metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var metric metric
//   var renderingProperties interface{}
//   var value interface{}
//
//   metricConfig := &MetricConfig{
//   	MathExpression: &MetricExpressionConfig{
//   		Expression: jsii.String("expression"),
//   		Period: jsii.Number(123),
//   		UsingMetrics: map[string]iMetric{
//   			"usingMetricsKey": metric,
//   		},
//
//   		// the properties below are optional
//   		SearchAccount: jsii.String("searchAccount"),
//   		SearchRegion: jsii.String("searchRegion"),
//   	},
//   	MetricStat: &MetricStatConfig{
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   		Period: duration,
//   		Statistic: jsii.String("statistic"),
//
//   		// the properties below are optional
//   		Account: jsii.String("account"),
//   		Dimensions: []dimension{
//   			&dimension{
//   				Name: jsii.String("name"),
//   				Value: value,
//   			},
//   		},
//   		Region: jsii.String("region"),
//   		UnitFilter: awscdk.Aws_cloudwatch.Unit_SECONDS,
//   	},
//   	RenderingProperties: map[string]interface{}{
//   		"renderingPropertiesKey": renderingProperties,
//   	},
//   }
//
// Experimental.
type MetricConfig struct {
	// In case the metric is a math expression, the details of the math expression.
	// Experimental.
	MathExpression *MetricExpressionConfig `field:"optional" json:"mathExpression" yaml:"mathExpression"`
	// In case the metric represents a query, the details of the query.
	// Experimental.
	MetricStat *MetricStatConfig `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Additional properties which will be rendered if the metric is used in a dashboard.
	//
	// Examples are 'label' and 'color', but any key in here will be
	// added to dashboard graphs.
	// Experimental.
	RenderingProperties *map[string]interface{} `field:"optional" json:"renderingProperties" yaml:"renderingProperties"`
}

