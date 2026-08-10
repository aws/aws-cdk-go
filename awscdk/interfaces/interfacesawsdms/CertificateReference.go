package interfacesawsdms


// A reference to a Certificate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateReference := &CertificateReference{
//   	CertificateArn: jsii.String("certificateArn"),
//   }
//
type CertificateReference struct {
	// The CertificateArn of the Certificate resource.
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
}

