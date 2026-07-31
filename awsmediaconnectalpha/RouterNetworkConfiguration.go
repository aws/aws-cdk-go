package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating Router Network configurations.
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
type RouterNetworkConfiguration interface {
}

// The jsii proxy struct for RouterNetworkConfiguration
type jsiiProxy_RouterNetworkConfiguration struct {
	_ byte // padding
}

// Create a public network configuration.
//
// Returns: RouterNetworkConfiguration instance for public setup.
// Experimental.
func RouterNetworkConfiguration_PublicNetwork(props *PublicNetworkConfigurationProps) RouterNetworkConfiguration {
	_init_.Initialize()

	if err := validateRouterNetworkConfiguration_PublicNetworkParameters(props); err != nil {
		panic(err)
	}
	var returns RouterNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterNetworkConfiguration",
		"publicNetwork",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a VPC network configuration.
//
// Returns: RouterNetworkConfiguration instance for VPC setup.
// Experimental.
func RouterNetworkConfiguration_Vpc(props *VpcNetworkConfigurationProps) RouterNetworkConfiguration {
	_init_.Initialize()

	if err := validateRouterNetworkConfiguration_VpcParameters(props); err != nil {
		panic(err)
	}
	var returns RouterNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterNetworkConfiguration",
		"vpc",
		[]interface{}{props},
		&returns,
	)

	return returns
}

