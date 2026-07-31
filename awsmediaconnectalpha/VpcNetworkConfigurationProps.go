package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for VPC network configuration.
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
type VpcNetworkConfigurationProps struct {
	// Security groups to associate with the network interface.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Subnet where the network interface will be created.
	// Experimental.
	Subnet awsec2.ISubnet `field:"required" json:"subnet" yaml:"subnet"`
}

