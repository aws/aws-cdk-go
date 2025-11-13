package interfacesawsec2


// A reference to a DHCPOptions resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dHCPOptionsReference := &DHCPOptionsReference{
//   	DhcpOptionsId: jsii.String("dhcpOptionsId"),
//   }
//
type DHCPOptionsReference struct {
	// The DhcpOptionsId of the DHCPOptions resource.
	DhcpOptionsId *string `field:"required" json:"dhcpOptionsId" yaml:"dhcpOptionsId"`
}

