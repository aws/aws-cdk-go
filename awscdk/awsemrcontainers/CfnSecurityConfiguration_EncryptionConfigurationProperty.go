package awsemrcontainers


// Encryption configuration for the security configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	AtRestEncryptionConfiguration: &AtRestEncryptionConfigurationProperty{
//   		LocalDiskEncryptionConfiguration: &LocalDiskEncryptionConfigurationProperty{
//   			AwsKmsKeyId: jsii.String("awsKmsKeyId"),
//   			EncryptionKeyProviderType: jsii.String("encryptionKeyProviderType"),
//   		},
//   		S3EncryptionConfiguration: &S3EncryptionConfigurationProperty{
//   			EncryptionOption: jsii.String("encryptionOption"),
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   	},
//   	InTransitEncryptionConfiguration: &InTransitEncryptionConfigurationProperty{
//   		TlsCertificateConfiguration: &TLSCertificateConfigurationProperty{
//   			CertificateProviderType: jsii.String("certificateProviderType"),
//   			PrivateKeySecretArn: jsii.String("privateKeySecretArn"),
//   			PublicKeySecretArn: jsii.String("publicKeySecretArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-encryptionconfiguration.html
//
type CfnSecurityConfiguration_EncryptionConfigurationProperty struct {
	// At-rest encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-encryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-encryptionconfiguration-atrestencryptionconfiguration
	//
	AtRestEncryptionConfiguration interface{} `field:"optional" json:"atRestEncryptionConfiguration" yaml:"atRestEncryptionConfiguration"`
	// In-transit encryption configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-encryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-encryptionconfiguration-intransitencryptionconfiguration
	//
	InTransitEncryptionConfiguration interface{} `field:"optional" json:"inTransitEncryptionConfiguration" yaml:"inTransitEncryptionConfiguration"`
}

