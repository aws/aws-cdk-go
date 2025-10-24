package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Network configuration for the Code Interpreter tool.
//
// Example:
//   // Create a custom execution role
//   executionRole := iam.NewRole(this, jsii.String("CodeInterpreterExecutionRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("bedrock-agentcore.amazonaws.com")),
//   })
//
//   // Create code interpreter with custom execution role
//   codeInterpreter := agentcore.NewCodeInterpreterCustom(this, jsii.String("MyCodeInterpreter"), &CodeInterpreterCustomProps{
//   	CodeInterpreterCustomName: jsii.String("my_code_interpreter"),
//   	Description: jsii.String("Code interpreter with custom execution role"),
//   	NetworkConfiguration: agentcore.CodeInterpreterNetworkConfiguration_UsingPublicNetwork(),
//   	ExecutionRole: executionRole,
//   })
//
// Experimental.
type CodeInterpreterNetworkConfiguration interface {
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

// The jsii proxy struct for CodeInterpreterNetworkConfiguration
type jsiiProxy_CodeInterpreterNetworkConfiguration struct {
	jsiiProxy_NetworkConfiguration
}

func (j *jsiiProxy_CodeInterpreterNetworkConfiguration) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeInterpreterNetworkConfiguration) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeInterpreterNetworkConfiguration) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeInterpreterNetworkConfiguration) VpcSubnets() *awsec2.SubnetSelection {
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
func NewCodeInterpreterNetworkConfiguration(mode *string, scope constructs.Construct, vpcConfig *VpcConfigProps) CodeInterpreterNetworkConfiguration {
	_init_.Initialize()

	if err := validateNewCodeInterpreterNetworkConfigurationParameters(mode, vpcConfig); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeInterpreterNetworkConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterNetworkConfiguration",
		[]interface{}{mode, scope, vpcConfig},
		&j,
	)

	return &j
}

// Creates a new network configuration.
// Experimental.
func NewCodeInterpreterNetworkConfiguration_Override(c CodeInterpreterNetworkConfiguration, mode *string, scope constructs.Construct, vpcConfig *VpcConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterNetworkConfiguration",
		[]interface{}{mode, scope, vpcConfig},
		c,
	)
}

// Creates a public network configuration.
//
// Returns: A CodeInterpreterNetworkConfiguration.
// Run this tool to operate in a public environment with internet access, suitable for less sensitive or open-use scenarios.
// Experimental.
func CodeInterpreterNetworkConfiguration_UsingPublicNetwork() CodeInterpreterNetworkConfiguration {
	_init_.Initialize()

	var returns CodeInterpreterNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterNetworkConfiguration",
		"usingPublicNetwork",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates a sandbox network configuration.
//
// Returns: A CodeInterpreterNetworkConfiguration.
// Run this tool in a restricted environment with limited Permissions and Encryption to enhance safety and reduce potential risks.
// Experimental.
func CodeInterpreterNetworkConfiguration_UsingSandboxNetwork() CodeInterpreterNetworkConfiguration {
	_init_.Initialize()

	var returns CodeInterpreterNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterNetworkConfiguration",
		"usingSandboxNetwork",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates a network configuration from a VPC configuration.
//
// Returns: A CodeInterpreterNetworkConfiguration.
// Experimental.
func CodeInterpreterNetworkConfiguration_UsingVpc(scope constructs.Construct, vpcConfig *VpcConfigProps) CodeInterpreterNetworkConfiguration {
	_init_.Initialize()

	if err := validateCodeInterpreterNetworkConfiguration_UsingVpcParameters(scope, vpcConfig); err != nil {
		panic(err)
	}
	var returns CodeInterpreterNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.CodeInterpreterNetworkConfiguration",
		"usingVpc",
		[]interface{}{scope, vpcConfig},
		&returns,
	)

	return returns
}

