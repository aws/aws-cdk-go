package mixinsawss3vectors


// Specifies the encryption configuration for the vector bucket.
//
// By default, all new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3 managed keys (SSE-S3), specifically AES256.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	SseType: jsii.String("sseType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-vectorbucket-encryptionconfiguration.html
//
type CfnVectorBucketPropsMixin_EncryptionConfigurationProperty struct {
	// AWS Key Management Service (KMS) customer managed key ARN to use for the encryption configuration.
	//
	// This parameter is required if and only if `SseType` is set to `aws:kms` .
	//
	// You must specify the full ARN of the KMS key. Key IDs or key aliases aren't supported.
	//
	// > Amazon S3 Vectors only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in AWS KMS](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-vectorbucket-encryptionconfiguration.html#cfn-s3vectors-vectorbucket-encryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The server-side encryption type to use for the encryption configuration of the vector bucket.
	//
	// Valid values are `AES256` for Amazon S3 managed keys and `aws:kms` for AWS KMS keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-vectorbucket-encryptionconfiguration.html#cfn-s3vectors-vectorbucket-encryptionconfiguration-ssetype
	//
	// Default: - "AES256".
	//
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

