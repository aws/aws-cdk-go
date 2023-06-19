package awscloudwatch


// Statistic to use over the aggregation period.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var deliveryStream deliveryStream
//
//
//   // Alarm that triggers when the per-second average of incoming bytes exceeds 90% of the current service limit
//   incomingBytesPercentOfLimit := cloudwatch.NewMathExpression(&MathExpressionProps{
//   	Expression: jsii.String("incomingBytes / 300 / bytePerSecLimit"),
//   	UsingMetrics: map[string]iMetric{
//   		"incomingBytes": deliveryStream.metricIncomingBytes(&MetricOptions{
//   			"statistic": cloudwatch.Statistic_SUM,
//   		}),
//   		"bytePerSecLimit": deliveryStream.metric(jsii.String("BytesPerSecondLimit")),
//   	},
//   })
//
//   cloudwatch.NewAlarm(this, jsii.String("Alarm"), &AlarmProps{
//   	Metric: incomingBytesPercentOfLimit,
//   	Threshold: jsii.Number(0.9),
//   	EvaluationPeriods: jsii.Number(3),
//   })
//
// Experimental.
type Statistic string

const (
	// The count (number) of data points used for the statistical calculation.
	// Experimental.
	Statistic_SAMPLE_COUNT Statistic = "SAMPLE_COUNT"
	// The value of Sum / SampleCount during the specified period.
	// Experimental.
	Statistic_AVERAGE Statistic = "AVERAGE"
	// All values submitted for the matching metric added together.
	//
	// This statistic can be useful for determining the total volume of a metric.
	// Experimental.
	Statistic_SUM Statistic = "SUM"
	// The lowest value observed during the specified period.
	//
	// You can use this value to determine low volumes of activity for your application.
	// Experimental.
	Statistic_MINIMUM Statistic = "MINIMUM"
	// The highest value observed during the specified period.
	//
	// You can use this value to determine high volumes of activity for your application.
	// Experimental.
	Statistic_MAXIMUM Statistic = "MAXIMUM"
)

