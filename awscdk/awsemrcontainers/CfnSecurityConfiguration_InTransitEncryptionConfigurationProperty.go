package awsemrcontainers


// In-transit encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inTransitEncryptionConfigurationProperty := &InTransitEncryptionConfigurationProperty{
//   	TlsCertificateConfiguration: &TLSCertificateConfigurationProperty{
//   		CertificateProviderType: jsii.String("certificateProviderType"),
//   		PrivateKeySecretArn: jsii.String("privateKeySecretArn"),
//   		PublicKeySecretArn: jsii.String("publicKeySecretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-intransitencryptionconfiguration.html
//
type CfnSecurityConfiguration_InTransitEncryptionConfigurationProperty struct {
	// TLS certificate configuration for in-transit encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-intransitencryptionconfiguration.html#cfn-emrcontainers-securityconfiguration-intransitencryptionconfiguration-tlscertificateconfiguration
	//
	TlsCertificateConfiguration interface{} `field:"optional" json:"tlsCertificateConfiguration" yaml:"tlsCertificateConfiguration"`
}

