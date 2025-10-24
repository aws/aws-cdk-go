package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Abstract base class for network configuration.
// Experimental.
type NetworkConfiguration interface {
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

// The jsii proxy struct for NetworkConfiguration
type jsiiProxy_NetworkConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkConfiguration) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConfiguration) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConfiguration) Scope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConfiguration) VpcSubnets() *awsec2.SubnetSelection {
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
func NewNetworkConfiguration_Override(n NetworkConfiguration, mode *string, scope constructs.Construct, vpcConfig *VpcConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.NetworkConfiguration",
		[]interface{}{mode, scope, vpcConfig},
		n,
	)
}

