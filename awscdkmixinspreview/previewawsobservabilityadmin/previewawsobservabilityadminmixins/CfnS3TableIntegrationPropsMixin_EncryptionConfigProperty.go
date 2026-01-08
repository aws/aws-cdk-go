package previewawsobservabilityadminmixins


// Defines the encryption configuration for S3 Table integrations, including the encryption algorithm and KMS key settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigProperty := &EncryptionConfigProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	SseAlgorithm: jsii.String("sseAlgorithm"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-encryptionconfig.html
//
type CfnS3TableIntegrationPropsMixin_EncryptionConfigProperty struct {
	// The Amazon Resource Name (ARN) of the KMS key used for encryption when using customer-managed keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-encryptionconfig.html#cfn-observabilityadmin-s3tableintegration-encryptionconfig-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The server-side encryption algorithm used for encrypting data in the S3 Table integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-s3tableintegration-encryptionconfig.html#cfn-observabilityadmin-s3tableintegration-encryptionconfig-ssealgorithm
	//
	SseAlgorithm *string `field:"optional" json:"sseAlgorithm" yaml:"sseAlgorithm"`
}

