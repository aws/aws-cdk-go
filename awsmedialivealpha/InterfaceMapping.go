package awsmedialivealpha


// A mapping between a logical interface name and a network ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   interfaceMapping := &InterfaceMapping{
//   	LogicalInterfaceName: jsii.String("logicalInterfaceName"),
//   	NetworkId: jsii.String("networkId"),
//   }
//
// Experimental.
type InterfaceMapping struct {
	// The logical interface name.
	// Experimental.
	LogicalInterfaceName *string `field:"required" json:"logicalInterfaceName" yaml:"logicalInterfaceName"`
	// The network ID to map to.
	// Experimental.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
}

