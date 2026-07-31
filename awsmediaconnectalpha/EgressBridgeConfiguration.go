package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Egress bridge configuration.
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
type EgressBridgeConfiguration struct {
	// The flow sources for the egress bridge.
	// Experimental.
	FlowSources *[]*BridgeFlowInput `field:"required" json:"flowSources" yaml:"flowSources"`
	// The maximum expected bitrate (in bps) of the bridge.
	// Experimental.
	MaxBitrate awscdk.Bitrate `field:"required" json:"maxBitrate" yaml:"maxBitrate"`
	// The network outputs for the egress bridge.
	// Experimental.
	NetworkOutputs *[]*BridgeNetworkOutput `field:"required" json:"networkOutputs" yaml:"networkOutputs"`
}

