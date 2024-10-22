package awss3express


// Specifies the default server-side encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionRuleProperty := &ServerSideEncryptionRuleProperty{
//   	BucketKeyEnabled: jsii.Boolean(false),
//   	ServerSideEncryptionByDefault: &ServerSideEncryptionByDefaultProperty{
//   		SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   		// the properties below are optional
//   		KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-serversideencryptionrule.html
//
type CfnDirectoryBucket_ServerSideEncryptionRuleProperty struct {
	// Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket.
	//
	// S3 Bucket Keys are always enabled for `GET` and `PUT` operations on a directory bucket and can’t be disabled. It's only allowed to set the `BucketKeyEnabled` element to `true` .
	//
	// S3 Bucket Keys aren't supported, when you copy SSE-KMS encrypted objects from general purpose buckets to directory buckets, from directory buckets to general purpose buckets, or between directory buckets, through [CopyObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html) , [UploadPartCopy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html) , [the Copy operation in Batch Operations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-objects-Batch-Ops) , or [the import jobs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-import-job) . In this case, Amazon S3 makes a call to AWS KMS every time a copy request is made for a KMS-encrypted object.
	//
	// For more information, see [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-UsingKMSEncryption.html#s3-express-sse-kms-bucket-keys) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-serversideencryptionrule.html#cfn-s3express-directorybucket-serversideencryptionrule-bucketkeyenabled
	//
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Specifies the default server-side encryption to apply to new objects in the bucket.
	//
	// If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-serversideencryptionrule.html#cfn-s3express-directorybucket-serversideencryptionrule-serversideencryptionbydefault
	//
	ServerSideEncryptionByDefault interface{} `field:"optional" json:"serverSideEncryptionByDefault" yaml:"serverSideEncryptionByDefault"`
}

