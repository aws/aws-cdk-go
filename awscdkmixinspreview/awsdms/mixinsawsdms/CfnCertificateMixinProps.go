package mixinsawsdms


// Properties for CfnCertificatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificateMixinProps := &CfnCertificateMixinProps{
//   	CertificateIdentifier: jsii.String("certificateIdentifier"),
//   	CertificatePem: jsii.String("certificatePem"),
//   	CertificateWallet: jsii.String("certificateWallet"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html
//
type CfnCertificateMixinProps struct {
	// A customer-assigned name for the certificate.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificateidentifier
	//
	CertificateIdentifier *string `field:"optional" json:"certificateIdentifier" yaml:"certificateIdentifier"`
	// The contents of a `.pem` file, which contains an X.509 certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificatepem
	//
	CertificatePem *string `field:"optional" json:"certificatePem" yaml:"certificatePem"`
	// The location of an imported Oracle Wallet certificate for use with SSL.
	//
	// An example is: `filebase64("${path.root}/rds-ca-2019-root.sso")`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificatewallet
	//
	CertificateWallet *string `field:"optional" json:"certificateWallet" yaml:"certificateWallet"`
}

