package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Network configuration for the Runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   runtimeNetworkConfiguration := bedrock_agentcore_alpha.RuntimeNetworkConfiguration_UsingPublicNetwork()
//
// Experimental.
type RuntimeNetworkConfiguration interface {
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

// The jsii proxy struct for RuntimeNetworkConfiguration
type jsiiProxy_RuntimeNetworkConfiguration struct {
	jsiiProxy_NetworkConfiguration
}

func (j *jsiiProxy_RuntimeNetworkConfiguration) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuntimeNetworkConfiguration) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuntimeNetworkConfiguration) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuntimeNetworkConfiguration) VpcSubnets() *awsec2.SubnetSelection {
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
func NewRuntimeNetworkConfiguration(mode *string, scope constructs.Construct, vpcConfig *VpcConfigProps) RuntimeNetworkConfiguration {
	_init_.Initialize()

	if err := validateNewRuntimeNetworkConfigurationParameters(mode, vpcConfig); err != nil {
		panic(err)
	}
	j := jsiiProxy_RuntimeNetworkConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeNetworkConfiguration",
		[]interface{}{mode, scope, vpcConfig},
		&j,
	)

	return &j
}

// Creates a new network configuration.
// Experimental.
func NewRuntimeNetworkConfiguration_Override(r RuntimeNetworkConfiguration, mode *string, scope constructs.Construct, vpcConfig *VpcConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeNetworkConfiguration",
		[]interface{}{mode, scope, vpcConfig},
		r,
	)
}

// Creates a public network configuration.
//
// PUBLIC is the default network mode.
//
// Returns: A RuntimeNetworkConfiguration.
// Run the runtime in a public environment with internet access, suitable for less sensitive or open-use scenarios.
// Experimental.
func RuntimeNetworkConfiguration_UsingPublicNetwork() RuntimeNetworkConfiguration {
	_init_.Initialize()

	var returns RuntimeNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeNetworkConfiguration",
		"usingPublicNetwork",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates a network configuration from a VPC configuration.
//
// Returns: A RuntimeNetworkConfiguration.
// Experimental.
func RuntimeNetworkConfiguration_UsingVpc(scope constructs.Construct, vpcConfig *VpcConfigProps) RuntimeNetworkConfiguration {
	_init_.Initialize()

	if err := validateRuntimeNetworkConfiguration_UsingVpcParameters(scope, vpcConfig); err != nil {
		panic(err)
	}
	var returns RuntimeNetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeNetworkConfiguration",
		"usingVpc",
		[]interface{}{scope, vpcConfig},
		&returns,
	)

	return returns
}

