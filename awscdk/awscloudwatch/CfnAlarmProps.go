package awscloudwatch


// Properties for defining a `CfnAlarm`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlarmProps := &CfnAlarmProps{
//   	ComparisonOperator: jsii.String("comparisonOperator"),
//   	EvaluationPeriods: jsii.Number(123),
//
//   	// the properties below are optional
//   	ActionsEnabled: jsii.Boolean(false),
//   	AlarmActions: []*string{
//   		jsii.String("alarmActions"),
//   	},
//   	AlarmDescription: jsii.String("alarmDescription"),
//   	AlarmName: jsii.String("alarmName"),
//   	DatapointsToAlarm: jsii.Number(123),
//   	Dimensions: []interface{}{
//   		&DimensionProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	EvaluateLowSampleCountPercentile: jsii.String("evaluateLowSampleCountPercentile"),
//   	ExtendedStatistic: jsii.String("extendedStatistic"),
//   	InsufficientDataActions: []*string{
//   		jsii.String("insufficientDataActions"),
//   	},
//   	MetricName: jsii.String("metricName"),
//   	Metrics: []interface{}{
//   		&MetricDataQueryProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			AccountId: jsii.String("accountId"),
//   			Expression: jsii.String("expression"),
//   			Label: jsii.String("label"),
//   			MetricStat: &MetricStatProperty{
//   				Metric: &MetricProperty{
//   					Dimensions: []interface{}{
//   						&DimensionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					MetricName: jsii.String("metricName"),
//   					Namespace: jsii.String("namespace"),
//   				},
//   				Period: jsii.Number(123),
//   				Stat: jsii.String("stat"),
//
//   				// the properties below are optional
//   				Unit: jsii.String("unit"),
//   			},
//   			Period: jsii.Number(123),
//   			ReturnData: jsii.Boolean(false),
//   		},
//   	},
//   	Namespace: jsii.String("namespace"),
//   	OkActions: []*string{
//   		jsii.String("okActions"),
//   	},
//   	Period: jsii.Number(123),
//   	Statistic: jsii.String("statistic"),
//   	Threshold: jsii.Number(123),
//   	ThresholdMetricId: jsii.String("thresholdMetricId"),
//   	TreatMissingData: jsii.String("treatMissingData"),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html
//
type CfnAlarmProps struct {
	// The arithmetic operation to use when comparing the specified statistic and threshold.
	//
	// The specified statistic value is used as the first operand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-comparisonoperator
	//
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of periods over which data is compared to the specified threshold.
	//
	// If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and `DatapointsToAlarm` is the M.
	//
	// For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *Amazon CloudWatch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-evaluationperiods
	//
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Indicates whether actions should be executed during any changes to the alarm state.
	//
	// The default is TRUE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-actionsenabled
	//
	// Default: - true.
	//
	ActionsEnabled interface{} `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state.
	//
	// Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *Amazon CloudWatch API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-alarmactions
	//
	AlarmActions *[]*string `field:"optional" json:"alarmActions" yaml:"alarmActions"`
	// The description of the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-alarmdescription
	//
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// The name of the alarm.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the alarm name.
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-alarmname
	//
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// The number of datapoints that must be breaching to trigger the alarm.
	//
	// This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for `EvaluationPeriods` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *Amazon CloudWatch User Guide* .
	//
	// If you omit this parameter, CloudWatch uses the same value here that you set for `EvaluationPeriods` , and the alarm goes to alarm state if that many consecutive periods are breaching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-datapointstoalarm
	//
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// The dimensions for the metric associated with the alarm.
	//
	// For an alarm based on a math expression, you can't specify `Dimensions` . Instead, you use `Metrics` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Used only for alarms based on percentiles.
	//
	// If `ignore` , the alarm state does not change during periods with too few data points to be statistically significant. If `evaluate` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-evaluatelowsamplecountpercentile
	//
	EvaluateLowSampleCountPercentile *string `field:"optional" json:"evaluateLowSampleCountPercentile" yaml:"evaluateLowSampleCountPercentile"`
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	//
	// For an alarm based on a metric, you must specify either `Statistic` or `ExtendedStatistic` but not both.
	//
	// For an alarm based on a math expression, you can't specify `ExtendedStatistic` . Instead, you use `Metrics` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-extendedstatistic
	//
	ExtendedStatistic *string `field:"optional" json:"extendedStatistic" yaml:"extendedStatistic"`
	// The actions to execute when this alarm transitions to the `INSUFFICIENT_DATA` state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-insufficientdataactions
	//
	InsufficientDataActions *[]*string `field:"optional" json:"insufficientDataActions" yaml:"insufficientDataActions"`
	// The name of the metric associated with the alarm.
	//
	// This is required for an alarm based on a metric. For an alarm based on a math expression, you use `Metrics` instead and you can't specify `MetricName` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// An array that enables you to create an alarm based on the result of a metric math expression.
	//
	// Each item in the array either retrieves a metric or performs a math expression.
	//
	// If you specify the `Metrics` parameter, you cannot specify `MetricName` , `Dimensions` , `Period` , `Namespace` , `Statistic` , `ExtendedStatistic` , or `Unit` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-metrics
	//
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// The namespace of the metric associated with the alarm.
	//
	// This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify `Namespace` and you use `Metrics` instead.
	//
	// For a list of namespaces for metrics from AWS services, see [AWS Services That Publish CloudWatch Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The actions to execute when this alarm transitions to the `OK` state from any other state.
	//
	// Each action is specified as an Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-okactions
	//
	OkActions *[]*string `field:"optional" json:"okActions" yaml:"okActions"`
	// The period, in seconds, over which the statistic is applied.
	//
	// This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.
	//
	// For an alarm based on a math expression, you can't specify `Period` , and instead you use the `Metrics` parameter.
	//
	// *Minimum:* 10.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-period
	//
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use `ExtendedStatistic` .
	//
	// For an alarm based on a metric, you must specify either `Statistic` or `ExtendedStatistic` but not both.
	//
	// For an alarm based on a math expression, you can't specify `Statistic` . Instead, you use `Metrics` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-statistic
	//
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The value to compare with the specified statistic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// In an alarm based on an anomaly detection model, this is the ID of the `ANOMALY_DETECTION_BAND` function used as the threshold for the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-thresholdmetricid
	//
	ThresholdMetricId *string `field:"optional" json:"thresholdMetricId" yaml:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points.
	//
	// Valid values are `breaching` , `notBreaching` , `ignore` , and `missing` . For more information, see [Configuring How CloudWatch Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon CloudWatch User Guide* .
	//
	// If you omit this parameter, the default behavior of `missing` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-treatmissingdata
	//
	TreatMissingData *string `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
	// The unit of the metric associated with the alarm.
	//
	// Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a `Metrics` array.
	//
	// You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-alarm.html#cfn-cloudwatch-alarm-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

