package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A network on a MediaConnect Gateway.
//
// Use {@link GatewayNetwork.define} to create a network and reference it from
// gateway, bridge source, and bridge output configurations.
//
// Example:
//   productionNetwork := awsmediaconnectalpha.GatewayNetwork_Define(&GatewayNetworkDefineProps{
//   	Name: jsii.String("production"),
//   	CidrBlock: jsii.String("10.0.0.0/16"),
//   })
//
// Experimental.
type GatewayNetwork interface {
	// A unique IP address range to use for this network in CIDR notation.
	// Experimental.
	CidrBlock() *string
	// The name of the network.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for GatewayNetwork
type jsiiProxy_GatewayNetwork struct {
	_ byte // padding
}

func (j *jsiiProxy_GatewayNetwork) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayNetwork) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Define a new gateway network.
// Experimental.
func GatewayNetwork_Define(props *GatewayNetworkDefineProps) GatewayNetwork {
	_init_.Initialize()

	if err := validateGatewayNetwork_DefineParameters(props); err != nil {
		panic(err)
	}
	var returns GatewayNetwork

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.GatewayNetwork",
		"define",
		[]interface{}{props},
		&returns,
	)

	return returns
}

