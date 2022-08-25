package awscloudwatch


// Properties of a rendered metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
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
//   		period: cdk.duration.minutes(jsii.Number(30)),
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
type MetricConfig struct {
	// In case the metric is a math expression, the details of the math expression.
	MathExpression *MetricExpressionConfig `field:"optional" json:"mathExpression" yaml:"mathExpression"`
	// In case the metric represents a query, the details of the query.
	MetricStat *MetricStatConfig `field:"optional" json:"metricStat" yaml:"metricStat"`
	// Additional properties which will be rendered if the metric is used in a dashboard.
	//
	// Examples are 'label' and 'color', but any key in here will be
	// added to dashboard graphs.
	RenderingProperties *map[string]interface{} `field:"optional" json:"renderingProperties" yaml:"renderingProperties"`
}

