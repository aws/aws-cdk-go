package awss3vectors


// The encryption configuration for a vector bucket or index.
//
// By default, if you don't specify, all new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3 managed keys (SSE-S3), specifically `AES256` . You can optionally override bucket level encryption settings, and set a specific encryption configuration for a vector index at the time of index creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	SseType: jsii.String("sseType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-index-encryptionconfiguration.html
//
type CfnIndex_EncryptionConfigurationProperty struct {
	// AWS Key Management Service (KMS) customer managed key ID to use for the encryption configuration.
	//
	// This parameter is allowed if and only if `sseType` is set to `aws:kms` .
	//
	// To specify the KMS key, you must use the format of the KMS key Amazon Resource Name (ARN).
	//
	// For example, specify Key ARN in the following format: `arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-index-encryptionconfiguration.html#cfn-s3vectors-index-encryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The server-side encryption type to use for the encryption configuration of the vector bucket.
	//
	// By default, if you don't specify, all new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3 managed keys (SSE-S3), specifically `AES256` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-index-encryptionconfiguration.html#cfn-s3vectors-index-encryptionconfiguration-ssetype
	//
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

