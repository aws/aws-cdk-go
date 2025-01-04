package awss3


// Describes the default server-side encryption to apply to new objects in the bucket.
//
// If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied. For more information, see [PutBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTencryption.html) .
//
// > - *General purpose buckets* - If you don't specify a customer managed key at configuration, Amazon S3 automatically creates an AWS KMS key ( `aws/s3` ) in your AWS account the first time that you add an object encrypted with SSE-KMS to a bucket. By default, Amazon S3 uses this KMS key for SSE-KMS.
// > - *Directory buckets* - Your SSE-KMS configuration can only support 1 [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk) per directory bucket's lifetime. The [AWS managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk) ( `aws/s3` ) isn't supported.
// > - *Directory buckets* - For directory buckets, there are only two supported options for server-side encryption: SSE-S3 and SSE-KMS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionByDefaultProperty := &ServerSideEncryptionByDefaultProperty{
//   	SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   	// the properties below are optional
//   	KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionbydefault.html
//
type CfnBucket_ServerSideEncryptionByDefaultProperty struct {
	// Server-side encryption algorithm to use for the default encryption.
	//
	// > For directory buckets, there are only two supported values for server-side encryption: `AES256` and `aws:kms` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionbydefault.html#cfn-s3-bucket-serversideencryptionbydefault-ssealgorithm
	//
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// AWS Key Management Service (KMS) customer managed key ID to use for the default encryption.
	//
	// > - *General purpose buckets* - This parameter is allowed if and only if `SSEAlgorithm` is set to `aws:kms` or `aws:kms:dsse` .
	// > - *Directory buckets* - This parameter is allowed if and only if `SSEAlgorithm` is set to `aws:kms` .
	//
	// You can specify the key ID, key alias, or the Amazon Resource Name (ARN) of the KMS key.
	//
	// - Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key ARN: `arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key Alias: `alias/alias-name`
	//
	// If you are using encryption with cross-account or AWS service operations, you must use a fully qualified KMS key ARN. For more information, see [Using encryption for cross-account operations](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html#bucket-encryption-update-bucket-policy) .
	//
	// > - *General purpose buckets* - If you're specifying a customer managed KMS key, we recommend using a fully qualified KMS key ARN. If you use a KMS key alias instead, then AWS KMS resolves the key within the requester’s account. This behavior can result in data that's encrypted with a KMS key that belongs to the requester, and not the bucket owner. Also, if you use a key ID, you can run into a LogDestination undeliverable error when creating a VPC flow log.
	// > - *Directory buckets* - When you specify an [AWS KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk) for encryption in your directory bucket, only use the key ID or key ARN. The key alias format of the KMS key isn't supported. > Amazon S3 only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in AWS KMS](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionbydefault.html#cfn-s3-bucket-serversideencryptionbydefault-kmsmasterkeyid
	//
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

