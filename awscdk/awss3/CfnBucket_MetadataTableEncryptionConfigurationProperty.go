package awss3


// The encryption settings for an S3 Metadata journal table or inventory table configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataTableEncryptionConfigurationProperty := &MetadataTableEncryptionConfigurationProperty{
//   	SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatatableencryptionconfiguration.html
//
type CfnBucket_MetadataTableEncryptionConfigurationProperty struct {
	// The encryption type specified for a metadata table.
	//
	// To specify server-side encryption with AWS Key Management Service ( AWS KMS ) keys (SSE-KMS), use the `aws:kms` value. To specify server-side encryption with Amazon S3 managed keys (SSE-S3), use the `AES256` value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatatableencryptionconfiguration.html#cfn-s3-bucket-metadatatableencryptionconfiguration-ssealgorithm
	//
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// If server-side encryption with AWS Key Management Service ( AWS KMS ) keys (SSE-KMS) is specified, you must also specify the KMS key Amazon Resource Name (ARN).
	//
	// You must specify a customer-managed KMS key that's located in the same Region as the general purpose bucket that corresponds to the metadata table configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metadatatableencryptionconfiguration.html#cfn-s3-bucket-metadatatableencryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

