package awsmsk


// S3 storage configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3StorageProperty := &S3StorageProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	CompressionType: jsii.String("compressionType"),
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   	OutputPrefix: jsii.String("outputPrefix"),
//   	StorageClass: jsii.String("storageClass"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3storage.html
//
type CfnChannelPropsMixin_S3StorageProperty struct {
	// ARN of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3storage.html#cfn-msk-channel-s3storage-bucketarn
	//
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// S3 compression type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3storage.html#cfn-msk-channel-s3storage-compressiontype
	//
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// Optional 12-digit AWS account ID expected to own the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3storage.html#cfn-msk-channel-s3storage-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Template for S3 key for output objects, used for partitioning.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3storage.html#cfn-msk-channel-s3storage-outputkeytemplate
	//
	OutputKeyTemplate *string `field:"optional" json:"outputKeyTemplate" yaml:"outputKeyTemplate"`
	// Optional prefix for output objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3storage.html#cfn-msk-channel-s3storage-outputprefix
	//
	OutputPrefix *string `field:"optional" json:"outputPrefix" yaml:"outputPrefix"`
	// S3 storage class.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-s3storage.html#cfn-msk-channel-s3storage-storageclass
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

