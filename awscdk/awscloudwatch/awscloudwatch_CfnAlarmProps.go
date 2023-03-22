package awscloudwatch


// Properties for defining a `CfnAlarm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmProps := &cfnAlarmProps{
//   	comparisonOperator: jsii.String("comparisonOperator"),
//   	evaluationPeriods: jsii.Number(123),
//
//   	// the properties below are optional
//   	actionsEnabled: jsii.Boolean(false),
//   	alarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	alarmDescription: jsii.String("alarmDescription"),
//   	alarmName: jsii.String("alarmName"),
//   	datapointsToAlarm: jsii.Number(123),
//   	dimensions: []interface{}{
//   		&dimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	evaluateLowSampleCountPercentile: jsii.String("evaluateLowSampleCountPercentile"),
//   	extendedStatistic: jsii.String("extendedStatistic"),
//   	insufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	metricName: jsii.String("metricName"),
//   	metrics: []interface{}{
//   		&metricDataQueryProperty{
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			accountId: jsii.String("accountId"),
//   			expression: jsii.String("expression"),
//   			label: jsii.String("label"),
//   			metricStat: &metricStatProperty{
//   				metric: &metricProperty{
//   					dimensions: []interface{}{
//   						&dimensionProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					metricName: jsii.String("metricName"),
//   					namespace: jsii.String("namespace"),
//   				},
//   				period: jsii.Number(123),
//   				stat: jsii.String("stat"),
//
//   				// the properties below are optional
//   				unit: jsii.String("unit"),
//   			},
//   			period: jsii.Number(123),
//   			returnData: jsii.Boolean(false),
//   		},
//   	},
//   	namespace: jsii.String("namespace"),
//   	okActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	period: jsii.Number(123),
//   	statistic: jsii.String("statistic"),
//   	threshold: jsii.Number(123),
//   	thresholdMetricId: jsii.String("thresholdMetricId"),
//   	treatMissingData: jsii.String("treatMissingData"),
//   	unit: jsii.String("unit"),
//   }
//
type CfnAlarmProps struct {
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	//
	// The specified statistic value is used as the first operand.
	//
	// You can specify the following values: `GreaterThanThreshold` , `GreaterThanOrEqualToThreshold` , `LessThanThreshold` , or `LessThanOrEqualToThreshold` .
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of periods over which data is compared to the specified threshold.
	//
	// If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and `DatapointsToAlarm` is the M.
	//
	// For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *Amazon CloudWatch User Guide* .
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Indicates whether actions should be executed during any changes to the alarm state.
	//
	// The default is TRUE.
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
	//
	// Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *Amazon CloudWatch API Reference* .
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// The description of the alarm.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// The name of the alarm.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the alarm name.
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The number of datapoints that must be breaching to trigger the alarm.
	//
	// This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for `EvaluationPeriods` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *Amazon CloudWatch User Guide* .
	//
	// If you omit this parameter, CloudWatch uses the same value here that you set for `EvaluationPeriods` , and the alarm goes to alarm state if that many consecutive periods are breaching.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The dimensions for the metric associated with the alarm.
	//
	// For an alarm based on a math expression, you can't specify `Dimensions` . Instead, you use `Metrics` .
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Used only for alarms based on percentiles.
	//
	// If `ignore` , the alarm state does not change during periods with too few data points to be statistically significant. If `evaluate` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
	EvaluateLowSampleCountPercentile *string `field:"optional" json:"evaluateLowSampleCountPercentile" yaml:"evaluateLowSampleCountPercentile"`
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	//
	// For an alarm based on a metric, you must specify either `Statistic` or `ExtendedStatistic` but not both.
	//
	// For an alarm based on a math expression, you can't specify `ExtendedStatistic` . Instead, you use `Metrics` .
	ExtendedStatistic *string `field:"optional" json:"extendedStatistic" yaml:"extendedStatistic"`
	// The actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// The name of the metric associated with the alarm.
	//
	// This is required for an alarm based on a metric. For an alarm based on a math expression, you use `Metrics` instead and you can't specify `MetricName` .
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// An array that enables you to create an alarm based on the result of a metric math expression.
	//
	// Each item in the array either retrieves a metric or performs a math expression.
	//
	// If you specify the `Metrics` parameter, you cannot specify `MetricName` , `Dimensions` , `Period` , `Namespace` , `Statistic` , `ExtendedStatistic` , or `Unit` .
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// The namespace of the metric associated with the alarm.
	//
	// This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify `Namespace` and you use `Metrics` instead.
	//
	// For a list of namespaces for metrics from AWS services, see [AWS Services That Publish CloudWatch Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The actions to execute when this alarm transitions to the `OK` state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN).
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// The period, in seconds, over which the statistic is applied.
	//
	// This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.
	//
	// For an alarm based on a math expression, you can't specify `Period` , and instead you use the `Metrics` parameter.
	//
	// *Minimum:* 10.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use `ExtendedStatistic` .
	//
	// For an alarm based on a metric, you must specify either `Statistic` or `ExtendedStatistic` but not both.
	//
	// For an alarm based on a math expression, you can't specify `Statistic` . Instead, you use `Metrics` .
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The value to compare with the specified statistic.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// In an alarm based on an anomaly detection model, this is the ID of the `ANOMALY_DETECTION_BAND` function used as the threshold for the alarm.
	ThresholdMetricId *string `field:"optional" json:"thresholdMetricId" yaml:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points.
	//
	// Valid values are `breaching` , `notBreaching` , `ignore` , and `missing` . For more information, see [Configuring How CloudWatch Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon CloudWatch User Guide* .
	//
	// If you omit this parameter, the default behavior of `missing` is used.
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
	// The unit of the metric associated with the alarm.
	//
	// Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a `Metrics` array.
	//
	// You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

