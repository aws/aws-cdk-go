package interfacesawscertificatemanager


// A reference to a AcmeExternalAccountBinding resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acmeExternalAccountBindingReference := &AcmeExternalAccountBindingReference{
//   	AcmeExternalAccountBindingArn: jsii.String("acmeExternalAccountBindingArn"),
//   }
//
type AcmeExternalAccountBindingReference struct {
	// The AcmeExternalAccountBindingArn of the AcmeExternalAccountBinding resource.
	AcmeExternalAccountBindingArn *string `field:"required" json:"acmeExternalAccountBindingArn" yaml:"acmeExternalAccountBindingArn"`
}

