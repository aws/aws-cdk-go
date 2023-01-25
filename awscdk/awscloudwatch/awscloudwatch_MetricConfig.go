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
//   metricConfig := &metricConfig{
//   	mathExpression: &metricExpressionConfig{
//   		expression: jsii.String("expression"),
//   		period: jsii.Number(123),
//   		usingMetrics: map[string]iMetric{
//   			"usingMetricsKey": metric,
//   		},
//
//   		// the properties below are optional
//   		searchAccount: jsii.String("searchAccount"),
//   		searchRegion: jsii.String("searchRegion"),
//   	},
//   	metricStat: &metricStatConfig{
//   		metricName: jsii.String("metricName"),
//   		namespace: jsii.String("namespace"),
//   		period: duration,
//   		statistic: jsii.String("statistic"),
//
//   		// the properties below are optional
//   		account: jsii.String("account"),
//   		dimensions: []dimension{
//   			&dimension{
//   				name: jsii.String("name"),
//   				value: value,
//   			},
//   		},
//   		region: jsii.String("region"),
//   		unitFilter: awscdk.Aws_cloudwatch.unit_SECONDS,
//   	},
//   	renderingProperties: map[string]interface{}{
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

