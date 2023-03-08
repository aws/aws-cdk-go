package awskinesisfirehose


// The `DynamicPartitioningConfiguration` property type specifies the configuration of the dynamic partitioning mechanism that creates targeted data sets from the streaming data by partitioning it based on partition keys.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamicPartitioningConfigurationProperty := &dynamicPartitioningConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   	retryOptions: &retryOptionsProperty{
//   		durationInSeconds: jsii.Number(123),
//   	},
//   }
//
type CfnDeliveryStream_DynamicPartitioningConfigurationProperty struct {
	// Specifies whether dynamic partitioning is enabled for this Kinesis Data Firehose delivery stream.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the retry behavior in case Kinesis Data Firehose is unable to deliver data to an Amazon S3 prefix.
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
}

