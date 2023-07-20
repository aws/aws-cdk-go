package awsconfig


// Properties for defining a `CfnDeliveryChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeliveryChannelProps := &CfnDeliveryChannelProps{
//   	S3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	ConfigSnapshotDeliveryProperties: &ConfigSnapshotDeliveryPropertiesProperty{
//   		DeliveryFrequency: jsii.String("deliveryFrequency"),
//   	},
//   	Name: jsii.String("name"),
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	S3KmsKeyArn: jsii.String("s3KmsKeyArn"),
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html
//
type CfnDeliveryChannelProps struct {
	// The name of the Amazon S3 bucket to which AWS Config delivers configuration snapshots and configuration history files.
	//
	// If you specify a bucket that belongs to another AWS account , that bucket must have policies that grant access permissions to AWS Config . For more information, see [Permissions for the Amazon S3 Bucket](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-policy.html) in the *AWS Config Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3bucketname
	//
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The options for how often AWS Config delivers configuration snapshots to the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties
	//
	ConfigSnapshotDeliveryProperties interface{} `field:"optional" json:"configSnapshotDeliveryProperties" yaml:"configSnapshotDeliveryProperties"`
	// A name for the delivery channel.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the delivery channel name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// Updates are not supported. To change the name, you must run two separate updates. In the first update, delete this resource, and then recreate it with a new name in the second update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The prefix for the specified Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3keyprefix
	//
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service ( AWS KMS ) AWS KMS key (KMS key) used to encrypt objects delivered by AWS Config .
	//
	// Must belong to the same Region as the destination S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3kmskeyarn
	//
	S3KmsKeyArn *string `field:"optional" json:"s3KmsKeyArn" yaml:"s3KmsKeyArn"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which AWS Config sends notifications about configuration changes.
	//
	// If you choose a topic from another account, the topic must have policies that grant access permissions to AWS Config . For more information, see [Permissions for the Amazon SNS Topic](https://docs.aws.amazon.com/config/latest/developerguide/sns-topic-policy.html) in the *AWS Config Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

