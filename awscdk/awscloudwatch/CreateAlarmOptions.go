package awscloudwatch


// Properties needed to make an alarm from a metric.
//
// Example:
//   var channelGroup ChannelGroup
//   var channel Channel
//   var endpoint OriginEndpoint
//
//
//   // Create a CloudWatch alarm on channel group egress bytes
//   alarm := channelGroup.metricEgressBytes().CreateAlarm(this, jsii.String("HighEgress"), &CreateAlarmOptions{
//   	Threshold: jsii.Number(1000),
//   	EvaluationPeriods: jsii.Number(1),
//   })
//
//   // Monitor channel ingress response time
//   channel.metricIngressResponseTime().CreateAlarm(this, jsii.String("SlowIngress"), &CreateAlarmOptions{
//   	Threshold: jsii.Number(1000),
//   	EvaluationPeriods: jsii.Number(2),
//   })
//
//   // Track origin endpoint request count
//   requestMetric := endpoint.metricEgressRequestCount(&MetricOptions{
//   	Statistic: jsii.String("sum"),
//   	Period: awscdk.Duration_Minutes(jsii.Number(5)),
//   })
//
type CreateAlarmOptions struct {
	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The value against which the specified statistic is compared.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
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
	// Comparison to use to check if metric is breaching.
	// Default: GreaterThanOrEqualToThreshold.
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
	// Sets how this alarm is to handle missing data points.
	// Default: TreatMissingData.Missing
	//
	TreatMissingData TreatMissingData `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

