package awsobservabilityadmin


// Defines the encryption configuration for S3 Table integrations, including the encryption algorithm and KMS key settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigProperty := &EncryptionConfigProperty{
//   	SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-encryptionconfig.html
//
type CfnS3TableIntegration_EncryptionConfigProperty struct {
	// The server-side encryption algorithm used for encrypting data in the S3 Table integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-encryptionconfig.html#cfn-observabilityadmin-s3tableintegration-encryptionconfig-ssealgorithm
	//
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// The Amazon Resource Name (ARN) of the KMS key used for encryption when using customer-managed keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-encryptionconfig.html#cfn-observabilityadmin-s3tableintegration-encryptionconfig-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

