package awsxray


// Configuration for the CloudWatch Logs log group that X-Ray creates when this resource enables Transaction Search.
//
// It is used only at that time and is never returned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cloudWatchLogsConfigurationProperty := &CloudWatchLogsConfigurationProperty{
//   	SpansLogGroupKmsKeyId: jsii.String("spansLogGroupKmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-transactionsearchconfig-cloudwatchlogsconfiguration.html
//
type CfnTransactionSearchConfigPropsMixin_CloudWatchLogsConfigurationProperty struct {
	// The ARN of a symmetric customer managed AWS KMS key in the same account and Region, used to encrypt the aws/spans log group.
	//
	// Key IDs and alias ARNs are not supported, so for an AWS::KMS::Key use Fn::GetAtt Arn, not Ref. X-Ray applies the key only when it creates aws/spans while enabling Transaction Search. If aws/spans already exists, for example from an earlier enablement or an earlier create that failed partway, the create fails and X-Ray does not enable Transaction Search or change the log group. Remove this property to use the existing log group with its current encryption, and use the CloudWatch Logs AssociateKmsKey API to change its key. The caller needs kms:DescribeKey on the key, and the key policy must allow CloudWatch Logs to use it. Adding, changing or removing this property replaces the resource, which fails while Transaction Search is enabled
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-transactionsearchconfig-cloudwatchlogsconfiguration.html#cfn-xray-transactionsearchconfig-cloudwatchlogsconfiguration-spansloggroupkmskeyid
	//
	SpansLogGroupKmsKeyId *string `field:"optional" json:"spansLogGroupKmsKeyId" yaml:"spansLogGroupKmsKeyId"`
}

