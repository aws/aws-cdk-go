package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a SearchExpression.
//
// Example:
//   cpuUtilization := cloudwatch.NewSearchExpression(&SearchExpressionProps{
//   	Expression: jsii.String("SEARCH('{AWS/EC2,InstanceId} MetricName=\"CPUUtilization\"', 'Average', 900)"),
//   	Label: jsii.String("EC2 CPU Utilization"),
//   	Color: jsii.String("#ff7f0e"),
//   })
//
type SearchExpressionProps struct {
	// Color for the metric produced by the search expression.
	//
	// If the search expression produces more than one time series, the color is assigned to the first one.
	// Other metrics are assigned colors automatically.
	// Default: - Automatically assigned.
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Label for this search expression when added to a Graph in a Dashboard.
	//
	// If this expression evaluates to more than one time series,
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
	// search expression produces more than one time series, the maximum
	// will be shown for each individual time series produce by this
	// search expression.
	// Default: - No label.
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The period over which the search expression's statistics are applied.
	//
	// This period overrides the period defined within the search expression.
	// Default: Duration.minutes(5)
	//
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Account to evaluate search expressions within.
	// Default: - Deployment account.
	//
	SearchAccount *string `field:"optional" json:"searchAccount" yaml:"searchAccount"`
	// Region to evaluate search expressions within.
	// Default: - Deployment region.
	//
	SearchRegion *string `field:"optional" json:"searchRegion" yaml:"searchRegion"`
	// The search expression defining the metrics to be retrieved.
	//
	// A search expression cannot be used within an Alarm.
	//
	// A search expression allows you to retrieve and graph multiple related metrics in a single statement.
	// It can return up to 500 time series.
	//
	// Examples:
	// - `SEARCH('{AWS/EC2,InstanceId} CPUUtilization', 'Average', 300)`
	// - `SEARCH('{AWS/ApplicationELB,LoadBalancer} RequestCount', 'Sum', 60)`
	// - `SEARCH('{MyNamespace,ServiceName} Errors', 'Sum')`
	//
	// For more information about search expression syntax, see:
	// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/search-expression-syntax.html
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

