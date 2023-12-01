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
	// Account ID for owner of detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-anomalyvisibilitytime
	//
	AnomalyVisibilityTime *float64 `field:"optional" json:"anomalyVisibilityTime" yaml:"anomalyVisibilityTime"`
	// Name of detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-detectorname
	//
	DetectorName *string `field:"optional" json:"detectorName" yaml:"detectorName"`
	// How often log group is evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-evaluationfrequency
	//
	EvaluationFrequency *string `field:"optional" json:"evaluationFrequency" yaml:"evaluationFrequency"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-filterpattern
	//
	FilterPattern *string `field:"optional" json:"filterPattern" yaml:"filterPattern"`
	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// List of Arns for the given log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loganomalydetector.html#cfn-logs-loganomalydetector-loggrouparnlist
	//
	LogGroupArnList *[]*string `field:"optional" json:"logGroupArnList" yaml:"logGroupArnList"`
}

