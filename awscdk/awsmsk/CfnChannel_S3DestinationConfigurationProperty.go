package awsmsk


// S3 destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationConfigurationProperty := &S3DestinationConfigurationProperty{
//   	DeadLetterQueueS3: &DeadLetterQueueS3Property{
//   		BucketArn: jsii.String("bucketArn"),
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//
//   		// the properties below are optional
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	},
//   	ServiceExecutionRoleArn: jsii.String("serviceExecutionRoleArn"),
//   	Storage: &S3StorageProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		CompressionType: jsii.String("compressionType"),
//   		StorageClass: jsii.String("storageClass"),
//
//   		// the properties below are optional
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   		OutputPrefix: jsii.String("outputPrefix"),
//   	},
//
//   	// the properties below are optional
//   	DataFreshnessInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html
//
type CfnChannel_S3DestinationConfigurationProperty struct {
	// Dead letter queue S3 configuration of the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html#cfn-msk-channel-s3destinationconfiguration-deadletterqueues3
	//
	DeadLetterQueueS3 interface{} `field:"required" json:"deadLetterQueueS3" yaml:"deadLetterQueueS3"`
	// The Amazon Resource Name (ARN) of an IAM role used by MSK to access S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html#cfn-msk-channel-s3destinationconfiguration-serviceexecutionrolearn
	//
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// S3 storage configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html#cfn-msk-channel-s3destinationconfiguration-storage
	//
	Storage interface{} `field:"required" json:"storage" yaml:"storage"`
	// Data freshness in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3destinationconfiguration.html#cfn-msk-channel-s3destinationconfiguration-datafreshnessinseconds
	//
	DataFreshnessInSeconds *float64 `field:"optional" json:"dataFreshnessInSeconds" yaml:"dataFreshnessInSeconds"`
}

