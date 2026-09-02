package awsmsk


// S3 destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3DestinationConfigurationProperty := &S3DestinationConfigurationProperty{
//   	DataFreshnessInSeconds: jsii.Number(123),
//   	DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   		BucketArn: jsii.String("bucketArn"),
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	},
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   	Storage: &S3StorageProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		CompressionType: jsii.String("compressionType"),
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   		OutputPrefix: jsii.String("outputPrefix"),
//   		StorageClass: jsii.String("storageClass"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html
//
type CfnChannelPropsMixin_S3DestinationConfigurationProperty struct {
	// Data freshness in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html#cfn-msk-channel-s3destinationconfiguration-datafreshnessinseconds
	//
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
	// Dead letter queue S3 configuration of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html#cfn-msk-channel-s3destinationconfiguration-deadletterqueues3
	//
	DeadLetterQueueS3 interface{} `field:"optional" json:"deadLetterQueueS3" yaml:"deadLetterQueueS3"`
	// The Amazon Resource Name (ARN) of an IAM role used by MSK to access S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html#cfn-msk-channel-s3destinationconfiguration-serviceexecutionrolearn
	//
	ServiceExecutionRoleArn *string `field:"optional" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// S3 storage configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html#cfn-msk-channel-s3destinationconfiguration-storage
	//
	Storage interface{} `field:"optional" json:"storage" yaml:"storage"`
}

