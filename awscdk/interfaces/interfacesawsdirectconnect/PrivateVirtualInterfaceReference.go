package interfacesawsdirectconnect


// A reference to a PrivateVirtualInterface resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateVirtualInterfaceReference := &PrivateVirtualInterfaceReference{
//   	VirtualInterfaceArn: jsii.String("virtualInterfaceArn"),
//   }
//
type PrivateVirtualInterfaceReference struct {
	// The VirtualInterfaceArn of the PrivateVirtualInterface resource.
	VirtualInterfaceArn *string `field:"required" json:"virtualInterfaceArn" yaml:"virtualInterfaceArn"`
}

