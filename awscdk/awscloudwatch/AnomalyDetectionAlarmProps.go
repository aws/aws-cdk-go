package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for Anomaly Detection Alarms.
//
// Example:
//   // Create a metric
//   metric := cloudwatch.NewMetric(&MetricProps{
//   	Namespace: jsii.String("AWS/EC2"),
//   	MetricName: jsii.String("CPUUtilization"),
//   	Statistic: jsii.String("Average"),
//   	Period: awscdk.Duration_Minutes(jsii.Number(5)),
//   })
//
//   // Create an anomaly detection alarm
//   alarm := cloudwatch.NewAnomalyDetectionAlarm(this, jsii.String("AnomalyAlarm"), &AnomalyDetectionAlarmProps{
//   	Metric: metric,
//   	EvaluationPeriods: jsii.Number(1),
//
//   	// Number of standard deviations for the band (default: 2)
//   	StdDevs: jsii.Number(2),
//   	// Alarm outside on either side of the band, or just below or above it (default: outside)
//   	ComparisonOperator: cloudwatch.ComparisonOperator_LESS_THAN_LOWER_OR_GREATER_THAN_UPPER_THRESHOLD,
//   	AlarmDescription: jsii.String("Alarm when metric is outside the expected band"),
//   })
//
type AnomalyDetectionAlarmProps struct {
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The metric to add the alarm on.
	//
	// Metric objects can be obtained from most resources, or you can construct
	// custom Metric objects by instantiating one.
	Metric IMetric `field:"required" json:"metric" yaml:"metric"`
	// Whether the actions for this alarm are enabled.
	// Default: true.
	//
	ActionsEnabled *bool `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Description for the alarm.
	// Default: No description.
	//
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Name of the alarm.
	// Default: Automatically generated name.
	//
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// Comparison operator to use to check if metric is breaching.
	//
	// Must be one of the anomaly detection operators:
	// - LESS_THAN_LOWER_OR_GREATER_THAN_UPPER_THRESHOLD
	// - GREATER_THAN_UPPER_THRESHOLD
	// - LESS_THAN_LOWER_THRESHOLD.
	// Default: LESS_THAN_LOWER_OR_GREATER_THAN_UPPER_THRESHOLD.
	//
	ComparisonOperator ComparisonOperator `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of datapoints that must be breaching to trigger the alarm.
	//
	// This is used only if you are setting an "M
	// out of N" alarm. In that case, this value is the M. For more information, see Evaluating an Alarm in the Amazon
	// CloudWatch User Guide.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation
	//
	// Default: ``evaluationPeriods``.
	//
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// Specifies whether to evaluate the data and potentially change the alarm state if there are too few data points to be statistically significant.
	//
	// Used only for alarms that are based on percentiles.
	// Default: - Not configured.
	//
	EvaluateLowSampleCountPercentile *string `field:"optional" json:"evaluateLowSampleCountPercentile" yaml:"evaluateLowSampleCountPercentile"`
	// The period over which the specified statistic is applied.
	//
	// Cannot be used with `MathExpression` objects.
	// Default: - The period from the metric.
	//
	// Deprecated: Use `metric.with({ period: ... })` to encode the period into the Metric object
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// What function to use for aggregating.
	//
	// Can be one of the following:
	//
	// - "Minimum" | "min"
	// - "Maximum" | "max"
	// - "Average" | "avg"
	// - "Sum" | "sum"
	// - "SampleCount | "n"
	// - "pNN.NN"
	//
	// Cannot be used with `MathExpression` objects.
	// Default: - The statistic from the metric.
	//
	// Deprecated: Use `metric.with({ statistic: ... })` to encode the period into the Metric object
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The number of standard deviations to use for the anomaly detection band.
	//
	// The higher the value, the wider the band.
	//
	// - Must be greater than 0. A value of 0 or negative values would not make sense in the context of calculating standard deviations.
	// - There is no strict maximum value defined, as standard deviations can theoretically extend infinitely. However, in practice, values beyond 5 or 6 standard deviations are rarely used, as they would result in an extremely wide anomaly detection band, potentially missing significant anomalies.
	// Default: 2.
	//
	StdDevs *float64 `field:"optional" json:"stdDevs" yaml:"stdDevs"`
	// Sets how this alarm is to handle missing data points.
	// Default: TreatMissingData.Missing
	//
	TreatMissingData TreatMissingData `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

