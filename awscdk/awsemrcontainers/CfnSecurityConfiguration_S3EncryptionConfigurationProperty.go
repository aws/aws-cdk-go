package awsemrcontainers


// S3 encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3EncryptionConfigurationProperty := &S3EncryptionConfigurationProperty{
//   	EncryptionOption: jsii.String("encryptionOption"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-s3encryptionconfiguration.html
//
type CfnSecurityConfiguration_S3EncryptionConfigurationProperty struct {
	// The S3 encryption option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-s3encryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-s3encryptionconfiguration-encryptionoption
	//
	EncryptionOption *string `field:"optional" json:"encryptionOption" yaml:"encryptionOption"`
	// The KMS key ID for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-s3encryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-s3encryptionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

