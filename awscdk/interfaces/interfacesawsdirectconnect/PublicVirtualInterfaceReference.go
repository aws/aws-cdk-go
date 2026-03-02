package interfacesawsdirectconnect


// A reference to a PublicVirtualInterface resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicVirtualInterfaceReference := &PublicVirtualInterfaceReference{
//   	VirtualInterfaceArn: jsii.String("virtualInterfaceArn"),
//   }
//
type PublicVirtualInterfaceReference struct {
	// The VirtualInterfaceArn of the PublicVirtualInterface resource.
	VirtualInterfaceArn *string `field:"required" json:"virtualInterfaceArn" yaml:"virtualInterfaceArn"`
}

