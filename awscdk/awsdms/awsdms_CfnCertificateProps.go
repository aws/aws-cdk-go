package awsdms


// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	certificateIdentifier: jsii.String("certificateIdentifier"),
//   	certificatePem: jsii.String("certificatePem"),
//   	certificateWallet: jsii.String("certificateWallet"),
//   }
//
type CfnCertificateProps struct {
	// A customer-assigned name for the certificate.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or contain two consecutive hyphens.
	CertificateIdentifier *string `field:"optional" json:"certificateIdentifier" yaml:"certificateIdentifier"`
	// The contents of a `.pem` file, which contains an X.509 certificate.
	CertificatePem *string `field:"optional" json:"certificatePem" yaml:"certificatePem"`
	// The location of an imported Oracle Wallet certificate for use with SSL.
	//
	// An example is: `filebase64("${path.root}/rds-ca-2019-root.sso")`
	CertificateWallet *string `field:"optional" json:"certificateWallet" yaml:"certificateWallet"`
}

