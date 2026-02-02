package awsemrcontainers


// TLS certificate configuration for in-transit encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tLSCertificateConfigurationProperty := &TLSCertificateConfigurationProperty{
//   	CertificateProviderType: jsii.String("certificateProviderType"),
//   	PrivateKeySecretArn: jsii.String("privateKeySecretArn"),
//   	PublicKeySecretArn: jsii.String("publicKeySecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-tlscertificateconfiguration.html
//
type CfnSecurityConfiguration_TLSCertificateConfigurationProperty struct {
	// The certificate provider type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-tlscertificateconfiguration.html#cfn-emrcontainers-securityconfiguration-tlscertificateconfiguration-certificateprovidertype
	//
	CertificateProviderType *string `field:"optional" json:"certificateProviderType" yaml:"certificateProviderType"`
	// The ARN of the secret containing the private key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-tlscertificateconfiguration.html#cfn-emrcontainers-securityconfiguration-tlscertificateconfiguration-privatekeysecretarn
	//
	PrivateKeySecretArn *string `field:"optional" json:"privateKeySecretArn" yaml:"privateKeySecretArn"`
	// The ARN of the secret containing the public key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-tlscertificateconfiguration.html#cfn-emrcontainers-securityconfiguration-tlscertificateconfiguration-publickeysecretarn
	//
	PublicKeySecretArn *string `field:"optional" json:"publicKeySecretArn" yaml:"publicKeySecretArn"`
}

