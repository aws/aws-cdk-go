package previewawsapprunnermixins


// Describes a custom encryption key that AWS App Runner uses to encrypt copies of the source repository and service logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKey: jsii.String("kmsKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-encryptionconfiguration.html
//
type CfnServicePropsMixin_EncryptionConfigurationProperty struct {
	// The ARN of the KMS key that's used for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-encryptionconfiguration.html#cfn-apprunner-service-encryptionconfiguration-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

