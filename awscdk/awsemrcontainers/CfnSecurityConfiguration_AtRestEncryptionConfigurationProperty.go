package awsemrcontainers


// At-rest encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   atRestEncryptionConfigurationProperty := &AtRestEncryptionConfigurationProperty{
//   	LocalDiskEncryptionConfiguration: &LocalDiskEncryptionConfigurationProperty{
//   		AwsKmsKeyId: jsii.String("awsKmsKeyId"),
//   		EncryptionKeyProviderType: jsii.String("encryptionKeyProviderType"),
//   	},
//   	S3EncryptionConfiguration: &S3EncryptionConfigurationProperty{
//   		EncryptionOption: jsii.String("encryptionOption"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-atrestencryptionconfiguration.html
//
type CfnSecurityConfiguration_AtRestEncryptionConfigurationProperty struct {
	// Local disk encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-atrestencryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-atrestencryptionconfiguration-localdiskencryptionconfiguration
	//
	LocalDiskEncryptionConfiguration interface{} `field:"optional" json:"localDiskEncryptionConfiguration" yaml:"localDiskEncryptionConfiguration"`
	// S3 encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-atrestencryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-atrestencryptionconfiguration-s3encryptionconfiguration
	//
	S3EncryptionConfiguration interface{} `field:"optional" json:"s3EncryptionConfiguration" yaml:"s3EncryptionConfiguration"`
}

