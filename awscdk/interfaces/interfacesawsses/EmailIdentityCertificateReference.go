package interfacesawsses


// A reference to a EmailIdentityCertificate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailIdentityCertificateReference := &EmailIdentityCertificateReference{
//   	EmailIdentity: jsii.String("emailIdentity"),
//   	FromAddress: jsii.String("fromAddress"),
//   }
//
type EmailIdentityCertificateReference struct {
	// The EmailIdentity of the EmailIdentityCertificate resource.
	EmailIdentity *string `field:"required" json:"emailIdentity" yaml:"emailIdentity"`
	// The FromAddress of the EmailIdentityCertificate resource.
	FromAddress *string `field:"required" json:"fromAddress" yaml:"fromAddress"`
}

