package previewawssecuritylakemixins


// Provides encryption details of the Amazon Security Lake object.
//
// The AWS shared responsibility model applies to data protection in Amazon Security Lake . As described in this model, AWS is responsible for protecting the global infrastructure that runs all of the AWS Cloud. You are responsible for maintaining control over your content that is hosted on this infrastructure. For more details, see [Data protection](https://docs.aws.amazon.com//security-lake/latest/userguide/data-protection.html) in the Amazon Security Lake User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-encryptionconfiguration.html
//
type CfnDataLakePropsMixin_EncryptionConfigurationProperty struct {
	// The ID of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-encryptionconfiguration.html#cfn-securitylake-datalake-encryptionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

