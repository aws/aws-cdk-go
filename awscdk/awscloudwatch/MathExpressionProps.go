package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a MathExpression.
//
// Example:
//   var matchmakingRuleSet matchmakingRuleSet
//
//   // Alarm that triggers when the per-second average of not placed matches exceed 10%
//   ruleEvaluationRatio := cloudwatch.NewMathExpression(&MathExpressionProps{
//   	Expression: jsii.String("1 - (ruleEvaluationsPassed / ruleEvaluationsFailed)"),
//   	UsingMetrics: map[string]iMetric{
//   		"ruleEvaluationsPassed": matchmakingRuleSet.metricRuleEvaluationsPassed(&MetricOptions{
//   			"statistic": cloudwatch.Statistic_SUM,
//   		}),
//   		"ruleEvaluationsFailed": matchmakingRuleSet.metric(jsii.String("ruleEvaluationsFailed")),
//   	},
//   })
//   cloudwatch.NewAlarm(this, jsii.String("Alarm"), &AlarmProps{
//   	Metric: ruleEvaluationRatio,
//   	Threshold: jsii.Number(0.1),
//   	EvaluationPeriods: jsii.Number(3),
//   })
//
type MathExpressionProps struct {
	// Color for this metric when added to a Graph in a Dashboard.
	// Default: - Automatic color.
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Label for this expression when added to a Graph in a Dashboard.
	//
	// If this expression evaluates to more than one time series (for
	// example, through the use of `METRICS()` or `SEARCH()` expressions),
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
	// math expression produces more than one time series, the maximum
	// will be shown for each individual time series produce by this
	// math expression.
	// Default: - Expression value is used as label.
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The period over which the math expression's statistics are applied.
	//
	// This period overrides all periods in the metrics used in this
	// math expression.
	// Default: Duration.minutes(5)
	//
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Account to evaluate search expressions within.
	//
	// Specifying a searchAccount has no effect to the account used
	// for metrics within the expression (passed via usingMetrics).
	// Default: - Deployment account.
	//
	SearchAccount *string `field:"optional" json:"searchAccount" yaml:"searchAccount"`
	// Region to evaluate search expressions within.
	//
	// Specifying a searchRegion has no effect to the region used
	// for metrics within the expression (passed via usingMetrics).
	// Default: - Deployment region.
	//
	SearchRegion *string `field:"optional" json:"searchRegion" yaml:"searchRegion"`
	// The expression defining the metric.
	//
	// When an expression contains a SEARCH function, it cannot be used
	// within an Alarm.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The metrics used in the expression, in a map.
	//
	// The key is the identifier that represents the given metric in the
	// expression, and the value is the actual Metric object.
	//
	// The `period` of each metric in `usingMetrics` is ignored and instead overridden
	// by the `period` specified for the `MathExpression` construct. Even if no `period`
	// is specified for the `MathExpression`, it will be overridden by the default
	// value (`Duration.minutes(5)`).
	//
	// Example:
	//
	// ```ts
	// declare const metrics: elbv2.IApplicationLoadBalancerMetrics;
	// new cloudwatch.MathExpression({
	//   expression:  'm1+m2',
	//   label: 'AlbErrors',
	//   usingMetrics: {
	//     m1: metrics.custom('HTTPCode_ELB_500_Count', {
	//       period: Duration.minutes(1), // <- This period will be ignored
	//       statistic: 'Sum',
	//       label: 'HTTPCode_ELB_500_Count',
	//     }),
	//     m2: metrics.custom('HTTPCode_ELB_502_Count', {
	//       period: Duration.minutes(1), // <- This period will be ignored
	//       statistic: 'Sum',
	//       label: 'HTTPCode_ELB_502_Count',
	//     }),
	//   },
	//   period: Duration.minutes(3), // <- This overrides the period of each metric in `usingMetrics`
	//                                //    (Even if not specified, it is overridden by the default value)
	// });
	// ```.
	// Default: - Empty map.
	//
	UsingMetrics *map[string]IMetric `field:"optional" json:"usingMetrics" yaml:"usingMetrics"`
}

