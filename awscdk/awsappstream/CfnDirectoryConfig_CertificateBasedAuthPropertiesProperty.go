package awsappstream


// The certificate-based authentication properties used to authenticate SAML 2.0 Identity Provider (IdP) user identities to Active Directory domain-joined streaming instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateBasedAuthPropertiesProperty := &CertificateBasedAuthPropertiesProperty{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	Status: jsii.String("status"),
//   }
//
type CfnDirectoryConfig_CertificateBasedAuthPropertiesProperty struct {
	// The ARN of the AWS Certificate Manager Private CA resource.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The status of the certificate-based authentication properties.
	//
	// Fallback is turned on by default when certificate-based authentication is *Enabled* . Fallback allows users to log in using their AD domain password if certificate-based authentication is unsuccessful, or to unlock a desktop lock screen. *Enabled_no_directory_login_fallback* enables certificate-based authentication, but does not allow users to log in using their AD domain password. Users will be disconnected to re-authenticate using certificates.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

