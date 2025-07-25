package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// A CloudWatch search expression for dynamically finding and graphing multiple related metrics.
//
// Search expressions allow you to search for and graph multiple related metrics from a single expression.
// This is particularly useful in dynamic environments where the exact metric names or dimensions
// may not be known at deployment time.
//
// Example:
// ```ts
// const searchExpression = new cloudwatch.SearchExpression({
//   expression: "SEARCH('{AWS/EC2,InstanceId} CPUUtilization', 'Average', 300)",
//   label: 'EC2 CPU Utilization',
//   period: Duration.minutes(5),
// });
// ```
//
// This class does not represent a resource, so hence is not a construct. Instead,
// SearchExpression is an abstraction that makes it easy to specify metrics for use in graphs.
//
// Example:
//   cpuUtilization := cloudwatch.NewSearchExpression(&SearchExpressionProps{
//   	Expression: jsii.String("SEARCH('{AWS/EC2,InstanceId} MetricName=\"CPUUtilization\"', 'Average', 900)"),
//   	Label: jsii.String("EC2 CPU Utilization"),
//   	Color: jsii.String("#ff7f0e"),
//   })
//
type SearchExpression interface {
	IMetric
	// Hex color code (e.g. '#00ff00'), to use when rendering the resulting metrics in a graph. If multiple time series are returned, color is assigned to the first metric, color for the other metrics is automatically assigned.
	Color() *string
	// The search expression defining the metrics to be retrieved.
	Expression() *string
	// The label is used as a prefix for the title of each metric returned by the search expression.
	Label() *string
	// The aggregation period for the metrics produced by the Search Expression.
	Period() awscdk.Duration
	// Account to evaluate search expressions within.
	SearchAccount() *string
	// Region to evaluate search expressions within.
	SearchRegion() *string
	// Warnings generated by this search expression.
	// Deprecated: - use warningsV2.
	Warnings() *[]*string
	// Warnings generated by this search expression.
	WarningsV2() *map[string]*string
	// Inspect the details of the metric object.
	ToMetricConfig() *MetricConfig
	// Returns a string representation of an object.
	ToString() *string
	// Return a copy of SearchExpression with properties changed.
	//
	// All properties except expression can be changed.
	With(props *SearchExpressionOptions) SearchExpression
}

// The jsii proxy struct for SearchExpression
type jsiiProxy_SearchExpression struct {
	jsiiProxy_IMetric
}

func (j *jsiiProxy_SearchExpression) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchExpression) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchExpression) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchExpression) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchExpression) SearchAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchExpression) SearchRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchExpression) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SearchExpression) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}


func NewSearchExpression(props *SearchExpressionProps) SearchExpression {
	_init_.Initialize()

	if err := validateNewSearchExpressionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SearchExpression{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.SearchExpression",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewSearchExpression_Override(s SearchExpression, props *SearchExpressionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.SearchExpression",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SearchExpression) ToMetricConfig() *MetricConfig {
	var returns *MetricConfig

	_jsii_.Invoke(
		s,
		"toMetricConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchExpression) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SearchExpression) With(props *SearchExpressionOptions) SearchExpression {
	if err := s.validateWithParameters(props); err != nil {
		panic(err)
	}
	var returns SearchExpression

	_jsii_.Invoke(
		s,
		"with",
		[]interface{}{props},
		&returns,
	)

	return returns
}

