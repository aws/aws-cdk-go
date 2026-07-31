package awsmediaconnectalpha


// Properties for Router Network Interface.
//
// Example:
//   var stack Stack
//   var securityGroup ISecurityGroup
//   var subnet ISubnet
//
//
//   privateInterface := awsmediaconnectalpha.NewRouterNetworkInterface(stack, jsii.String("PrivateInterface"), &RouterNetworkInterfaceProps{
//   	RouterNetworkInterfaceName: jsii.String("private-interface"),
//   	Configuration: awsmediaconnectalpha.RouterNetworkConfiguration_Vpc(&VpcNetworkConfigurationProps{
//   		SecurityGroups: []ISecurityGroup{
//   			securityGroup,
//   		},
//   		Subnet: subnet,
//   	}),
//   })
//
// Experimental.
type RouterNetworkInterfaceProps struct {
	// Network configuration for the router network interface.
	// Experimental.
	Configuration RouterNetworkConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The AWS Region where the router network interface will be created.
	// Default: - Same region as the stack.
	//
	// Experimental.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// The name of the router network interface.
	// Default: - Generated automatically.
	//
	// Experimental.
	RouterNetworkInterfaceName *string `field:"optional" json:"routerNetworkInterfaceName" yaml:"routerNetworkInterfaceName"`
	// Tags to add to the network interface.
	// Default: - No tagging.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

