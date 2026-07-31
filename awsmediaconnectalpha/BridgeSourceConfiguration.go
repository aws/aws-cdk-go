package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating bridge source configurations.
//
// Example:
//   var stack Stack
//   var bridge Bridge
//   var flow Flow
//
//
//   // Add a flow source to an egress bridge (requires failover to be enabled)
//   additionalSource := awsmediaconnectalpha.NewBridgeSource(stack, jsii.String("AdditionalSource"), &BridgeSourceProps{
//   	BridgeSourceName: jsii.String("backup-source"),
//   	Bridge: bridge,
//   	Source: awsmediaconnectalpha.BridgeSourceConfiguration_Flow(&BridgeFlowSource{
//   		Flow: flow,
//   	}),
//   })
//
// Experimental.
type BridgeSourceConfiguration interface {
}

// The jsii proxy struct for BridgeSourceConfiguration
type jsiiProxy_BridgeSourceConfiguration struct {
	_ byte // padding
}

// Experimental.
func NewBridgeSourceConfiguration(flowConfig *BridgeFlowSource, networkConfig *BridgeNetworkSource) BridgeSourceConfiguration {
	_init_.Initialize()

	if err := validateNewBridgeSourceConfigurationParameters(flowConfig, networkConfig); err != nil {
		panic(err)
	}
	j := jsiiProxy_BridgeSourceConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeSourceConfiguration",
		[]interface{}{flowConfig, networkConfig},
		&j,
	)

	return &j
}

// Experimental.
func NewBridgeSourceConfiguration_Override(b BridgeSourceConfiguration, flowConfig *BridgeFlowSource, networkConfig *BridgeNetworkSource) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeSourceConfiguration",
		[]interface{}{flowConfig, networkConfig},
		b,
	)
}

// Create a flow source configuration for a bridge source.
//
// Returns: Bridge source configuration with flow source.
// Experimental.
func BridgeSourceConfiguration_Flow(source *BridgeFlowSource) BridgeSourceConfiguration {
	_init_.Initialize()

	if err := validateBridgeSourceConfiguration_FlowParameters(source); err != nil {
		panic(err)
	}
	var returns BridgeSourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeSourceConfiguration",
		"flow",
		[]interface{}{source},
		&returns,
	)

	return returns
}

// Create a network source configuration for a bridge source.
//
// Returns: Bridge source configuration with network source.
// Experimental.
func BridgeSourceConfiguration_Network(source *BridgeNetworkSource) BridgeSourceConfiguration {
	_init_.Initialize()

	if err := validateBridgeSourceConfiguration_NetworkParameters(source); err != nil {
		panic(err)
	}
	var returns BridgeSourceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeSourceConfiguration",
		"network",
		[]interface{}{source},
		&returns,
	)

	return returns
}

