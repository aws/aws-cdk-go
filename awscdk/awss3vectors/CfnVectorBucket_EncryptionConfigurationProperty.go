package awss3vectors


// The encryption configuration for the vector bucket.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-vectorbucket-encryptionconfiguration.html
//
type CfnVectorBucket_EncryptionConfigurationProperty struct {
	// AWS Key Management Service (KMS) customer managed key ID to use for the encryption configuration.
	//
	// This parameter is allowed if and only if sseType is set to aws:kms.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-vectorbucket-encryptionconfiguration.html#cfn-s3vectors-vectorbucket-encryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The server-side encryption type to use for the encryption configuration of the vector bucket.
	//
	// By default, if you don't specify, all new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3 managed keys (SSE-S3), specifically AES256.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3vectors-vectorbucket-encryptionconfiguration.html#cfn-s3vectors-vectorbucket-encryptionconfiguration-ssetype
	//
	// Default: - "AES256".
	//
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

