package previewawss3mixins


// A bucket-level setting for Amazon S3 general purpose buckets used to prevent the upload of new objects encrypted with the specified server-side encryption type.
//
// For example, blocking an encryption type will block `PutObject` , `CopyObject` , `PostObject` , multipart upload, and replication requests to the bucket for objects with the specified encryption type. However, you can continue to read and list any pre-existing objects already encrypted with the specified encryption type. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/blocking-unblocking-s3-c-encryption-gpb.html) .
//
// This data type is used with the following actions:
//
// - [PutBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html)
// - [GetBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html)
// - [DeleteBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketEncryption.html)
//
// - **Permissions** - You must have the `s3:PutEncryptionConfiguration` permission to block or unblock an encryption type for a bucket.
//
// You must have the `s3:GetEncryptionConfiguration` permission to view a bucket's encryption type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   blockedEncryptionTypesProperty := &BlockedEncryptionTypesProperty{
//   	EncryptionType: []*string{
//   		jsii.String("encryptionType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-blockedencryptiontypes.html
//
type CfnBucketPropsMixin_BlockedEncryptionTypesProperty struct {
	// The object encryption type that you want to block or unblock for an Amazon S3 general purpose bucket.
	//
	// > Currently, this parameter only supports blocking or unblocking server side encryption with customer-provided keys (SSE-C). For more information about SSE-C, see [Using server-side encryption with customer-provided keys (SSE-C)](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerSideEncryptionCustomerKeys.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-blockedencryptiontypes.html#cfn-s3-bucket-blockedencryptiontypes-encryptiontype
	//
	EncryptionType *[]*string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
}

