package awsmediaconnectalpha


// The source of the bridge.
//
// A flow source originates in MediaConnect as an existing cloud flow.
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
type BridgeFlowSource struct {
	// The cloud flow used as a source of this bridge.
	// Experimental.
	Flow IFlow `field:"required" json:"flow" yaml:"flow"`
	// The VPC interface attachment to use for this source.
	// Default: - no VPC interface.
	//
	// Experimental.
	VpcInterface *VpcInterfaceConfig `field:"optional" json:"vpcInterface" yaml:"vpcInterface"`
}

