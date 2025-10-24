package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Network configuration for the Browser tool.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("testVPC"))
//
//   codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
//   	CodeInterpreterCustomName: jsii.String("my_sandbox_interpreter"),
//   	Description: jsii.String("Code interpreter with isolated network access"),
//   	NetworkConfiguration: agentcore.BrowserNetworkConfiguration_UsingVpc(this, &VpcConfigProps{
//   		Vpc: vpc,
//   	}),
//   })
//
//   codeInterpreter.connections.AddSecurityGroup(ec2.NewSecurityGroup(this, jsii.String("AdditionalGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   }))
//
// Experimental.
type BrowserNetworkConfiguration interface {
	NetworkConfiguration
	// The connections object to the network.
	// Experimental.
	Connections() awsec2.Connections
	// The network mode to use.
	//
	// Configure the security level for agent
	// execution to control access, isolate resources, and protect sensitive data.
	// Experimental.
	NetworkMode() *string
	// The scope to create the resource in.
	// Experimental.
	Scope() constructs.Construct
	// The VPC subnets to use.
	// Experimental.
	VpcSubnets() *awsec2.SubnetSelection
}

// The jsii proxy struct for BrowserNetworkConfiguration
type jsiiProxy_BrowserNetworkConfiguration struct {
	jsiiProxy_NetworkConfiguration
}

func (j *jsiiProxy_BrowserNetworkConfiguration) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BrowserNetworkConfiguration) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BrowserNetworkConfiguration) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BrowserNetworkConfiguration) VpcSubnets() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcSubnets",
		&returns,
	)
	return returns
}


// Creates a new network configuration.
// Experimental.
func NewBrowserNetworkConfiguration(mode *string, scope constructs.Construct, vpcConfig *VpcConfigProps) BrowserNetworkConfiguration {
	_init_.Initialize()

	if err := validateNewBrowserNetworkConfigurationParameters(mode, vpcConfig); err != nil {
		panic(err)
	}
	j := jsiiProxy_BrowserNetworkConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserNetworkConfiguration",
		[]interface{}{mode, scope, vpcConfig},
		&j,
	)

	return &j
}

// Creates a new network configuration.
// Experimental.
func NewBrowserNetworkConfiguration_Override(b BrowserNetworkConfiguration, mode *string, scope constructs.Construct, vpcConfig *VpcConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserNetworkConfiguration",
		[]interface{}{mode, scope, vpcConfig},
		b,
	)
}

// Creates a public network configuration.
//
// PUBLIC is the default network mode.
//
// Returns: A BrowserNetworkConfiguration.
// Run this tool to operate in a public environment with internet access, suitable for less sensitive or open-use scenarios.
// Experimental.
func BrowserNetworkConfiguration_UsingPublicNetwork() BrowserNetworkConfiguration {
	_init_.Initialize()

	var returns BrowserNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserNetworkConfiguration",
		"usingPublicNetwork",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates a network configuration from a VPC configuration.
//
// Returns: A BrowserNetworkConfiguration.
// Experimental.
func BrowserNetworkConfiguration_UsingVpc(scope constructs.Construct, vpcConfig *VpcConfigProps) BrowserNetworkConfiguration {
	_init_.Initialize()

	if err := validateBrowserNetworkConfiguration_UsingVpcParameters(scope, vpcConfig); err != nil {
		panic(err)
	}
	var returns BrowserNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.BrowserNetworkConfiguration",
		"usingVpc",
		[]interface{}{scope, vpcConfig},
		&returns,
	)

	return returns
}

