package awss3tables


// Configuration specifying how data should be encrypted.
//
// This structure defines the encryption algorithm and optional KMS key to be used for server-side encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	SseAlgorithm: jsii.String("sseAlgorithm"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-encryptionconfiguration.html
//
type CfnTableBucket_EncryptionConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the KMS key to use for encryption.
	//
	// This field is required only when `sseAlgorithm` is set to `aws:kms` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-encryptionconfiguration.html#cfn-s3tables-tablebucket-encryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The server-side encryption algorithm to use.
	//
	// Valid values are `AES256` for S3-managed encryption keys, or `aws:kms` for AWS KMS-managed encryption keys. If you choose SSE-KMS encryption you must grant the S3 Tables maintenance principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption](https://docs.aws.amazon.com//AmazonS3/latest/userguide/s3-tables-kms-permissions.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-encryptionconfiguration.html#cfn-s3tables-tablebucket-encryptionconfiguration-ssealgorithm
	//
	SseAlgorithm *string `field:"optional" json:"sseAlgorithm" yaml:"sseAlgorithm"`
}

