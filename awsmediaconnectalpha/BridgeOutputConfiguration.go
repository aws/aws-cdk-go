package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for a bridge output.
//
// Example:
//   var stack Stack
//   var gateway Gateway
//   var flow Flow
//   var vpcInterface VpcInterfaceConfig
//   var productionNetwork GatewayNetwork
//
//
//   egressBridge := awsmediaconnectalpha.NewBridge(stack, jsii.String("MyEgressBridge"), &BridgeProps{
//   	BridgeName: jsii.String("my-egress-bridge"),
//   	Config: awsmediaconnectalpha.BridgeConfiguration_Egress(&EgressBridgeConfiguration{
//   		MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   		FlowSources: []BridgeFlowInput{
//   			&BridgeFlowInput{
//   				Name: jsii.String("cloud-source"),
//   				Source: &BridgeFlowSource{
//   					Flow: flow,
//   					VpcInterface: vpcInterface,
//   				},
//   			},
//   		},
//   		NetworkOutputs: []BridgeNetworkOutput{
//   			&BridgeNetworkOutput{
//   				Name: jsii.String("on-prem-output"),
//   				Output: awsmediaconnectalpha.BridgeOutputConfiguration_Network(&BridgeNetworkOutputProps{
//   					IpAddress: jsii.String("192.168.1.200"),
//   					Port: jsii.Number(5001),
//   					Network: productionNetwork,
//   					Protocol: awsmediaconnectalpha.BridgeProtocol_RTP(),
//   					Ttl: jsii.Number(50),
//   				}),
//   			},
//   		},
//   	}),
//   	Gateway: gateway,
//   })
//
// Experimental.
type BridgeOutputConfiguration interface {
}

// The jsii proxy struct for BridgeOutputConfiguration
type jsiiProxy_BridgeOutputConfiguration struct {
	_ byte // padding
}

// Create a network output configuration for a bridge.
// Experimental.
func BridgeOutputConfiguration_Network(props *BridgeNetworkOutputProps) BridgeOutputConfiguration {
	_init_.Initialize()

	if err := validateBridgeOutputConfiguration_NetworkParameters(props); err != nil {
		panic(err)
	}
	var returns BridgeOutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeOutputConfiguration",
		"network",
		[]interface{}{props},
		&returns,
	)

	return returns
}

