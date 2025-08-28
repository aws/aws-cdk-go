package awscloudwatch


// Properties for Alarms.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var alias alias
//
//   // or add alarms to an existing group
//   var blueGreenAlias alias
//
//   alarm := cloudwatch.NewAlarm(this, jsii.String("Errors"), &AlarmProps{
//   	ComparisonOperator: cloudwatch.ComparisonOperator_GREATER_THAN_THRESHOLD,
//   	Threshold: jsii.Number(1),
//   	EvaluationPeriods: jsii.Number(1),
//   	Metric: alias.metricErrors(),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
//   	Alias: Alias,
//   	DeploymentConfig: codedeploy.LambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
//   	Alarms: []iAlarm{
//   		alarm,
//   	},
//   })
//   deploymentGroup.AddAlarm(cloudwatch.NewAlarm(this, jsii.String("BlueGreenErrors"), &AlarmProps{
//   	ComparisonOperator: cloudwatch.ComparisonOperator_GREATER_THAN_THRESHOLD,
//   	Threshold: jsii.Number(1),
//   	EvaluationPeriods: jsii.Number(1),
//   	Metric: blueGreenAlias.metricErrors(),
//   }))
//
type AlarmProps struct {
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
	// The metric to add the alarm on.
	//
	// Metric objects can be obtained from most resources, or you can construct
	// custom Metric objects by instantiating one.
	Metric IMetric `field:"required" json:"metric" yaml:"metric"`
}

