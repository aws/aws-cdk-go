package awscloudwatch


// Statistic to use over the aggregation period.
//
// Example:
//   var fleet buildFleet
//
//   // Alarm that triggers when the per-second average of not used instances exceed 10%
//   instancesUsedRatio := cloudwatch.NewMathExpression(&mathExpressionProps{
//   	expression: jsii.String("1 - (activeInstances / idleInstances)"),
//   	usingMetrics: map[string]iMetric{
//   		"activeInstances": fleet.metric(jsii.String("ActiveInstances"), &MetricOptions{
//   			"statistic": cloudwatch.Statistic_SUM,
//   		}),
//   		"idleInstances": fleet.metricIdleInstances(),
//   	},
//   })
//   cloudwatch.NewAlarm(this, jsii.String("Alarm"), &alarmProps{
//   	metric: instancesUsedRatio,
//   	threshold: jsii.Number(0.1),
//   	evaluationPeriods: jsii.Number(3),
//   })
//
type Statistic string

const (
	// The count (number) of data points used for the statistical calculation.
	Statistic_SAMPLE_COUNT Statistic = "SAMPLE_COUNT"
	// The value of Sum / SampleCount during the specified period.
	Statistic_AVERAGE Statistic = "AVERAGE"
	// All values submitted for the matching metric added together.
	//
	// This statistic can be useful for determining the total volume of a metric.
	Statistic_SUM Statistic = "SUM"
	// The lowest value observed during the specified period.
	//
	// You can use this value to determine low volumes of activity for your application.
	Statistic_MINIMUM Statistic = "MINIMUM"
	// The highest value observed during the specified period.
	//
	// You can use this value to determine high volumes of activity for your application.
	Statistic_MAXIMUM Statistic = "MAXIMUM"
)

