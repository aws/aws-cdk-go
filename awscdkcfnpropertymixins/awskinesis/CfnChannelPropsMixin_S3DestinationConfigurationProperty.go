package awskinesis


// Configuration for delivery to a vanilla S3 bucket destination.
//
// Exactly one of S3DestinationConfiguration and S3TablesDestinationConfiguration must be specified on the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3DestinationConfigurationProperty := &S3DestinationConfigurationProperty{
//   	DataFreshnessInSeconds: jsii.Number(123),
//   	DeadLetterQueueS3Configuration: &DeadLetterQueueS3ConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	},
//   	StorageConfiguration: &S3StorageConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		CompressionType: jsii.String("compressionType"),
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   		StorageClass: jsii.String("storageClass"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3destinationconfiguration.html
//
type CfnChannelPropsMixin_S3DestinationConfigurationProperty struct {
	// The maximum time in seconds the channel buffers records before delivery if the minimum target file size is not reached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3destinationconfiguration.html#cfn-kinesis-channel-s3destinationconfiguration-datafreshnessinseconds
	//
	// Default: - 300.
	//
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
	// Configuration of the S3 bucket used to capture records that cannot be delivered to the primary destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3destinationconfiguration.html#cfn-kinesis-channel-s3destinationconfiguration-deadletterqueues3configuration
	//
	DeadLetterQueueS3Configuration interface{} `field:"optional" json:"deadLetterQueueS3Configuration" yaml:"deadLetterQueueS3Configuration"`
	// S3 storage configuration that describes the destination bucket and how delivered objects are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3destinationconfiguration.html#cfn-kinesis-channel-s3destinationconfiguration-storageconfiguration
	//
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

