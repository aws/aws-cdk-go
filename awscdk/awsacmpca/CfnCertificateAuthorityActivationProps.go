package awsacmpca


// Properties for defining a `CfnCertificateAuthorityActivation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateAuthorityActivationProps := &CfnCertificateAuthorityActivationProps{
//   	Certificate: jsii.String("certificate"),
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//
//   	// the properties below are optional
//   	CertificateChain: jsii.String("certificateChain"),
//   	Status: jsii.String("status"),
//   }
//
type CfnCertificateAuthorityActivationProps struct {
	// The Base64 PEM-encoded certificate authority certificate.
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// The Amazon Resource Name (ARN) of your private CA.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The Base64 PEM-encoded certificate chain that chains up to the root CA certificate that you used to sign your private CA certificate.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// Status of your private CA.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

