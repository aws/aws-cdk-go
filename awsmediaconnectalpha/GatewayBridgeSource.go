package awsmediaconnectalpha


// Options for Gateway Bridge Source.
//
// Example:
//   var stack Stack
//   var bridge Bridge
//   var role IRole
//   var securityGroup ISecurityGroup
//   var subnet ISubnet
//
//
//   vpcInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("bridge-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Subnet: subnet,
//   })
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_GatewayBridge(&GatewayBridgeSource{
//   		Bridge: bridge,
//   		VpcInterface: vpcInterface,
//   	}),
//   	VpcInterfaces: []VpcInterfaceConfig{
//   		vpcInterface,
//   	},
//   })
//
// Experimental.
type GatewayBridgeSource struct {
	// The bridge feeding this flow.
	// Experimental.
	Bridge IBridge `field:"required" json:"bridge" yaml:"bridge"`
	// The VPC interface attachment to use for this bridge source.
	// Default: - no VPC interface.
	//
	// Experimental.
	VpcInterface *VpcInterfaceConfig `field:"optional" json:"vpcInterface" yaml:"vpcInterface"`
}

