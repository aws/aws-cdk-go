package awslogs


// Properties for defining a `CfnLogAnomalyDetector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLogAnomalyDetectorProps := &CfnLogAnomalyDetectorProps{
//   	AccountId: jsii.String("accountId"),
//   	AnomalyVisibilityTime: jsii.Number(123),
//   	DetectorName: jsii.String("detectorName"),
//   	EvaluationFrequency: jsii.String("evaluationFrequency"),
//   	FilterPattern: jsii.String("filterPattern"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogGroupArnList: []*string{
//   		jsii.String("logGroupArnList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html
//
type CfnLogAnomalyDetectorProps struct {
	// The ID of the account to create the anomaly detector in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The number of days to have visibility on an anomaly.
	//
	// After this time period has elapsed for an anomaly, it will be automatically baselined and the anomaly detector will treat new occurrences of a similar anomaly as normal. Therefore, if you do not correct the cause of an anomaly during the time period specified in `AnomalyVisibilityTime` , it will be considered normal going forward and will not be detected as an anomaly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-anomalyvisibilitytime
	//
	AnomalyVisibilityTime *float64 `field:"optional" json:"anomalyVisibilityTime" yaml:"anomalyVisibilityTime"`
	// A name for this anomaly detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-detectorname
	//
	DetectorName *string `field:"optional" json:"detectorName" yaml:"detectorName"`
	// Specifies how often the anomaly detector is to run and look for anomalies.
	//
	// Set this value according to the frequency that the log group receives new logs. For example, if the log group receives new log events every 10 minutes, then 15 minutes might be a good setting for `EvaluationFrequency` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-evaluationfrequency
	//
	EvaluationFrequency *string `field:"optional" json:"evaluationFrequency" yaml:"evaluationFrequency"`
	// You can use this parameter to limit the anomaly detection model to examine only log events that match the pattern you specify here.
	//
	// For more information, see [Filter and Pattern Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-filterpattern
	//
	FilterPattern *string `field:"optional" json:"filterPattern" yaml:"filterPattern"`
	// Optionally assigns a AWS KMS key to secure this anomaly detector and its findings.
	//
	// If a key is assigned, the anomalies found and the model used by this detector are encrypted at rest with the key. If a key is assigned to an anomaly detector, a user must have permissions for both this key and for the anomaly detector to retrieve information about the anomalies that it finds.
	//
	// For more information about using a AWS KMS key and to see the required IAM policy, see [Use a AWS KMS key with an anomaly detector](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/LogsAnomalyDetection-KMS.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The ARN of the log group that is associated with this anomaly detector.
	//
	// You can specify only one log group ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-loggrouparnlist
	//
	LogGroupArnList *[]*string `field:"optional" json:"logGroupArnList" yaml:"logGroupArnList"`
}

