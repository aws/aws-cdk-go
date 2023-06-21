package awscloudwatch


// Comparison operator for evaluating alarms.
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
type ComparisonOperator string

const (
	// Specified statistic is greater than or equal to the threshold.
	ComparisonOperator_GREATER_THAN_OR_EQUAL_TO_THRESHOLD ComparisonOperator = "GREATER_THAN_OR_EQUAL_TO_THRESHOLD"
	// Specified statistic is strictly greater than the threshold.
	ComparisonOperator_GREATER_THAN_THRESHOLD ComparisonOperator = "GREATER_THAN_THRESHOLD"
	// Specified statistic is strictly less than the threshold.
	ComparisonOperator_LESS_THAN_THRESHOLD ComparisonOperator = "LESS_THAN_THRESHOLD"
	// Specified statistic is less than or equal to the threshold.
	ComparisonOperator_LESS_THAN_OR_EQUAL_TO_THRESHOLD ComparisonOperator = "LESS_THAN_OR_EQUAL_TO_THRESHOLD"
	// Specified statistic is lower than or greater than the anomaly model band.
	//
	// Used only for alarms based on anomaly detection models.
	ComparisonOperator_LESS_THAN_LOWER_OR_GREATER_THAN_UPPER_THRESHOLD ComparisonOperator = "LESS_THAN_LOWER_OR_GREATER_THAN_UPPER_THRESHOLD"
	// Specified statistic is greater than the anomaly model band.
	//
	// Used only for alarms based on anomaly detection models.
	ComparisonOperator_GREATER_THAN_UPPER_THRESHOLD ComparisonOperator = "GREATER_THAN_UPPER_THRESHOLD"
	// Specified statistic is lower than the anomaly model band.
	//
	// Used only for alarms based on anomaly detection models.
	ComparisonOperator_LESS_THAN_LOWER_THRESHOLD ComparisonOperator = "LESS_THAN_LOWER_THRESHOLD"
)

