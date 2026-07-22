package interfacesawsdirectconnect


// A reference to a TransitVirtualInterface resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitVirtualInterfaceReference := &TransitVirtualInterfaceReference{
//   	VirtualInterfaceArn: jsii.String("virtualInterfaceArn"),
//   	VirtualInterfaceId: jsii.String("virtualInterfaceId"),
//   }
//
type TransitVirtualInterfaceReference struct {
	// The VirtualInterfaceArn of the TransitVirtualInterface resource.
	VirtualInterfaceArn *string `field:"required" json:"virtualInterfaceArn" yaml:"virtualInterfaceArn"`
	// The VirtualInterfaceId of the TransitVirtualInterface resource.
	VirtualInterfaceId *string `field:"required" json:"virtualInterfaceId" yaml:"virtualInterfaceId"`
}

