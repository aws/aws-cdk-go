package awsmediaconnectalpha


// Properties for defining a Gateway network.
//
// Example:
//   var stack Stack
//
//
//   productionNetwork := awsmediaconnectalpha.GatewayNetwork_Define(&GatewayNetworkDefineProps{
//   	CidrBlock: jsii.String("192.168.1.0/24"),
//   	Name: jsii.String("production-network"),
//   })
//
//   gateway := awsmediaconnectalpha.NewGateway(stack, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	EgressCidrBlocks: []*string{
//   		jsii.String("10.0.0.0/16"),
//   	},
//   	Networks: []GatewayNetwork{
//   		productionNetwork,
//   	},
//   })
//
// Experimental.
type GatewayNetworkDefineProps struct {
	// A unique IP address range to use for this network.
	//
	// Must be in CIDR notation
	// (for example, `10.0.0.0/16`).
	// Experimental.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// The name of the network.
	//
	// Used to reference this network from bridge sources
	// and outputs, and must be unique among the networks on the gateway.
	//
	// Maximum 64 characters; alphanumeric, hyphens, and underscores only.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

