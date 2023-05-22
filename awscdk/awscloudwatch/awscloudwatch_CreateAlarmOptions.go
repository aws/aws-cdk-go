package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties needed to make an alarm from a metric.
//
// Example:
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//   var myHostedZone hostedZone
//
//   certificate := acm.NewCertificate(this, jsii.String("Certificate"), &CertificateProps{
//   	DomainName: jsii.String("hello.example.com"),
//   	Validation: acm.CertificateValidation_FromDns(myHostedZone),
//   })
//   certificate.metricDaysToExpiry().CreateAlarm(this, jsii.String("Alarm"), &CreateAlarmOptions{
//   	ComparisonOperator: cloudwatch.ComparisonOperator_LESS_THAN_THRESHOLD,
//   	EvaluationPeriods: jsii.Number(1),
//   	Threshold: jsii.Number(45),
//   })
//
// Experimental.
type CreateAlarmOptions struct {
	// The number of periods over which data is compared to the specified threshold.
	// Experimental.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// The value against which the specified statistic is compared.
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Whether the actions for this alarm are enabled.
	// Experimental.
	ActionsEnabled *bool `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Description for the alarm.
	// Experimental.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// Name of the alarm.
	// Experimental.
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
	// Comparison to use to check if metric is breaching.
	// Experimental.
	ComparisonOperator ComparisonOperator `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The number of datapoints that must be breaching to trigger the alarm.
	//
	// This is used only if you are setting an "M
	// out of N" alarm. In that case, this value is the M. For more information, see Evaluating an Alarm in the Amazon
	// CloudWatch User Guide.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation
	//
	// Experimental.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// Specifies whether to evaluate the data and potentially change the alarm state if there are too few data points to be statistically significant.
	//
	// Used only for alarms that are based on percentiles.
	// Experimental.
	EvaluateLowSampleCountPercentile *string `field:"optional" json:"evaluateLowSampleCountPercentile" yaml:"evaluateLowSampleCountPercentile"`
	// The period over which the specified statistic is applied.
	//
	// Cannot be used with `MathExpression` objects.
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
	// Deprecated: Use `metric.with({ statistic: ... })` to encode the period into the Metric object
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Sets how this alarm is to handle missing data points.
	// Experimental.
	TreatMissingData TreatMissingData `field:"optional" json:"treatMissingData" yaml:"treatMissingData"`
}

