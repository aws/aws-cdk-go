package interfacesawsdms


// A reference to a Certificate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateReference := &CertificateReference{
//   	CertificateId: jsii.String("certificateId"),
//   }
//
type CertificateReference struct {
	// The Id of the Certificate resource.
	CertificateId *string `field:"required" json:"certificateId" yaml:"certificateId"`
}

