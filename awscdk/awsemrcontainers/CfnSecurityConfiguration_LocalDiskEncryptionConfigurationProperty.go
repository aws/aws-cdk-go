package awsemrcontainers


// Local disk encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localDiskEncryptionConfigurationProperty := &LocalDiskEncryptionConfigurationProperty{
//   	AwsKmsKeyId: jsii.String("awsKmsKeyId"),
//   	EncryptionKeyProviderType: jsii.String("encryptionKeyProviderType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-localdiskencryptionconfiguration.html
//
type CfnSecurityConfiguration_LocalDiskEncryptionConfigurationProperty struct {
	// The AWS KMS key ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-localdiskencryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-localdiskencryptionconfiguration-awskmskeyid
	//
	AwsKmsKeyId *string `field:"optional" json:"awsKmsKeyId" yaml:"awsKmsKeyId"`
	// The encryption key provider type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-localdiskencryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-localdiskencryptionconfiguration-encryptionkeyprovidertype
	//
	EncryptionKeyProviderType *string `field:"optional" json:"encryptionKeyProviderType" yaml:"encryptionKeyProviderType"`
}

