package interfacesawspaymentcryptography


// A reference to a Key resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyReference := &KeyReference{
//   	KeyIdentifier: jsii.String("keyIdentifier"),
//   }
//
type KeyReference struct {
	// The KeyIdentifier of the Key resource.
	KeyIdentifier *string `field:"required" json:"keyIdentifier" yaml:"keyIdentifier"`
}

