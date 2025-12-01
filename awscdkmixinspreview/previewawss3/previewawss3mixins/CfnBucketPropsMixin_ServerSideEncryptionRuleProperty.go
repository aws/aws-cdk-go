package previewawss3mixins


// Specifies the default server-side encryption configuration.
//
// > - *General purpose buckets* - If you're specifying a customer managed KMS key, we recommend using a fully qualified KMS key ARN. If you use a KMS key alias instead, then AWS  resolves the key within the requester’s account. This behavior can result in data that's encrypted with a KMS key that belongs to the requester, and not the bucket owner.
// > - *Directory buckets* - When you specify an [AWS  customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk) for encryption in your directory bucket, only use the key ID or key ARN. The key alias format of the KMS key isn't supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serverSideEncryptionRuleProperty := &ServerSideEncryptionRuleProperty{
//   	BlockedEncryptionTypes: &BlockedEncryptionTypesProperty{
//   		EncryptionType: []*string{
//   			jsii.String("encryptionType"),
//   		},
//   	},
//   	BucketKeyEnabled: jsii.Boolean(false),
//   	ServerSideEncryptionByDefault: &ServerSideEncryptionByDefaultProperty{
//   		KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionrule.html
//
type CfnBucketPropsMixin_ServerSideEncryptionRuleProperty struct {
	// A bucket-level setting for Amazon S3 general purpose buckets used to prevent the upload of new objects encrypted with the specified server-side encryption type.
	//
	// For example, blocking an encryption type will block `PutObject` , `CopyObject` , `PostObject` , multipart upload, and replication requests to the bucket for objects with the specified encryption type. However, you can continue to read and list any pre-existing objects already encrypted with the specified encryption type. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/blocking-unblocking-s3-c-encryption-gpb.html) .
	//
	// > Currently, this parameter only supports blocking or unblocking server-side encryption with customer-provided keys (SSE-C). For more information about SSE-C, see [Using server-side encryption with customer-provided keys (SSE-C)](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerSideEncryptionCustomerKeys.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionrule.html#cfn-s3-bucket-serversideencryptionrule-blockedencryptiontypes
	//
	BlockedEncryptionTypes interface{} `field:"optional" json:"blockedEncryptionTypes" yaml:"blockedEncryptionTypes"`
	// Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket.
	//
	// Existing objects are not affected. Setting the `BucketKeyEnabled` element to `true` causes Amazon S3 to use an S3 Bucket Key. By default, S3 Bucket Key is not enabled.
	//
	// For more information, see [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionrule.html#cfn-s3-bucket-serversideencryptionrule-bucketkeyenabled
	//
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Specifies the default server-side encryption to apply to new objects in the bucket.
	//
	// If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-serversideencryptionrule.html#cfn-s3-bucket-serversideencryptionrule-serversideencryptionbydefault
	//
	ServerSideEncryptionByDefault interface{} `field:"optional" json:"serverSideEncryptionByDefault" yaml:"serverSideEncryptionByDefault"`
}

