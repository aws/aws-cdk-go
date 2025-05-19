package awskinesisfirehose


// The `DynamicPartitioningConfiguration` property type specifies the configuration of the dynamic partitioning mechanism that creates targeted data sets from the streaming data by partitioning it based on partition keys.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamicPartitioningConfigurationProperty := &DynamicPartitioningConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	RetryOptions: &RetryOptionsProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dynamicpartitioningconfiguration.html
//
type CfnDeliveryStream_DynamicPartitioningConfigurationProperty struct {
	// Specifies whether dynamic partitioning is enabled for this Kinesis Data Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dynamicpartitioningconfiguration.html#cfn-kinesisfirehose-deliverystream-dynamicpartitioningconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the retry behavior in case Kinesis Data Firehose is unable to deliver data to an Amazon S3 prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dynamicpartitioningconfiguration.html#cfn-kinesisfirehose-deliverystream-dynamicpartitioningconfiguration-retryoptions
	//
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
}

