package interfacesawsacmpca


// A reference to a CertificateAuthority resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateAuthorityReference := &CertificateAuthorityReference{
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   }
//
type CertificateAuthorityReference struct {
	// The Arn of the CertificateAuthority resource.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
}

