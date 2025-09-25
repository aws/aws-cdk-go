package awsacmpca


// A reference to a CertificateAuthorityActivation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateAuthorityActivationReference := &CertificateAuthorityActivationReference{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   }
//
type CertificateAuthorityActivationReference struct {
	// The CertificateAuthorityArn of the CertificateAuthorityActivation resource.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
}

