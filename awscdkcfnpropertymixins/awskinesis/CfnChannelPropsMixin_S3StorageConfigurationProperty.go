package awskinesis


// S3 storage configuration that describes the destination bucket and how delivered objects are stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3StorageConfigurationProperty := &S3StorageConfigurationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	CompressionType: jsii.String("compressionType"),
//   	ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   	OutputKeyTemplate: jsii.String("outputKeyTemplate"),
//   	StorageClass: jsii.String("storageClass"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3storageconfiguration.html
//
type CfnChannelPropsMixin_S3StorageConfigurationProperty struct {
	// The ARN of the S3 bucket for record delivery.
	//
	// Different channels can deliver to the same bucket. Buckets can be cross-account but must be in the same region as the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3storageconfiguration.html#cfn-kinesis-channel-s3storageconfiguration-bucketarn
	//
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The compression algorithm applied to delivered objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3storageconfiguration.html#cfn-kinesis-channel-s3storageconfiguration-compressiontype
	//
	CompressionType *string `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The AWS account ID of the expected owner of the destination S3 bucket.
	//
	// Used to verify bucket ownership before delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3storageconfiguration.html#cfn-kinesis-channel-s3storageconfiguration-expectedbucketowner
	//
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Optional template for the S3 object key path.
	//
	// Supports placeholders in the form !{name}: !{channel-name}, !{channel-id}, !{stream-name}, !{yyyy}, !{yy}, !{MM}, !{dd}, !{HH}, !{mm}, and !{extension} (a literal file extension can be supplied as !{extension:.json.gz}). When omitted, the service uses the default 'kinesis-channel/!{channel-name}/!{channel-id}/!{yyyy}/!{MM}/!{dd}/!{HH}/!{channel-name}-!{channel-id}-!{yyyy}-!{MM}-!{dd}-!{HH}-!{mm}!{extension}'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3storageconfiguration.html#cfn-kinesis-channel-s3storageconfiguration-outputkeytemplate
	//
	OutputKeyTemplate *string `field:"optional" json:"outputKeyTemplate" yaml:"outputKeyTemplate"`
	// The S3 storage class for delivered objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-channel-s3storageconfiguration.html#cfn-kinesis-channel-s3storageconfiguration-storageclass
	//
	// Default: - "STANDARD".
	//
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

