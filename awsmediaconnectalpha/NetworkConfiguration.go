package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines network configuration for a source — either a public network with a CIDR allowlist, or a VPC interface.
//
// Example:
//   var stack Stack
//
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyFlow"), &FlowProps{
//   	Source: awsmediaconnectalpha.SourceConfiguration_SrtListener(&SourceSrtListener{
//   		FlowSourceName: jsii.String("live-encoder-source"),
//   		Description: jsii.String("Live encoder feed"),
//   		Port: jsii.Number(5000),
//   		MinLatency: awscdk.Duration_Millis(jsii.Number(2000)),
//   		Network: awsmediaconnectalpha.NetworkConfiguration_PublicNetwork(jsii.String("203.0.113.0/24")),
//   	}),
//   })
//
// Experimental.
type NetworkConfiguration interface {
	// The CIDR allowlist for public internet sources, or undefined if using VPC.
	// Experimental.
	AllowlistCidr() *string
	// The VPC interface name, or undefined if using public internet.
	// Experimental.
	VpcInterfaceName() *string
}

// The jsii proxy struct for NetworkConfiguration
type jsiiProxy_NetworkConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkConfiguration) AllowlistCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowlistCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConfiguration) VpcInterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcInterfaceName",
		&returns,
	)
	return returns
}


// Use a public network with a CIDR allowlist.
// Experimental.
func NetworkConfiguration_PublicNetwork(allowlistCidr *string) NetworkConfiguration {
	_init_.Initialize()

	if err := validateNetworkConfiguration_PublicNetworkParameters(allowlistCidr); err != nil {
		panic(err)
	}
	var returns NetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.NetworkConfiguration",
		"publicNetwork",
		[]interface{}{allowlistCidr},
		&returns,
	)

	return returns
}

// Use a VPC interface.
// Experimental.
func NetworkConfiguration_Vpc(vpcInterface *VpcInterfaceConfig) NetworkConfiguration {
	_init_.Initialize()

	if err := validateNetworkConfiguration_VpcParameters(vpcInterface); err != nil {
		panic(err)
	}
	var returns NetworkConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.NetworkConfiguration",
		"vpc",
		[]interface{}{vpcInterface},
		&returns,
	)

	return returns
}

