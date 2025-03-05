package awskinesisfirehose


// The structure that configures parameters such as `ThroughputHintInMBs` for a stream configured with Direct PUT as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directPutSourceConfigurationProperty := &DirectPutSourceConfigurationProperty{
//   	ThroughputHintInMBs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-directputsourceconfiguration.html
//
type CfnDeliveryStream_DirectPutSourceConfigurationProperty struct {
	// The value that you configure for this parameter is for information purpose only and does not affect Firehose delivery throughput limit.
	//
	// You can use the [Firehose Limits form](https://docs.aws.amazon.com/https://support.console.aws.amazon.com/support/home#/case/create%3FissueType=service-limit-increase%26limitType=kinesis-firehose-limits) to request a throughput limit increase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-directputsourceconfiguration.html#cfn-kinesisfirehose-deliverystream-directputsourceconfiguration-throughputhintinmbs
	//
	ThroughputHintInMBs *float64 `field:"optional" json:"throughputHintInMBs" yaml:"throughputHintInMBs"`
}

