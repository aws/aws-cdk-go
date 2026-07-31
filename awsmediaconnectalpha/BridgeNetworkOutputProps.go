package awsmediaconnectalpha


// Properties for a bridge network output.
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
type BridgeNetworkOutputProps struct {
	// The IP address where the output will send content.
	// Experimental.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The gateway network this output sends content out of.
	//
	// Use {@link GatewayNetwork.define} to create the network and pass the same
	// instance to the gateway and to each output that uses it.
	// Experimental.
	Network GatewayNetwork `field:"required" json:"network" yaml:"network"`
	// The port to use for the output.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The protocol to use for the output.
	// Experimental.
	Protocol BridgeProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// Time to live (TTL) for the output packets in hops (1-255).
	//
	// TTL represents the maximum number of network hops a packet can traverse
	// before being discarded.
	// Experimental.
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

