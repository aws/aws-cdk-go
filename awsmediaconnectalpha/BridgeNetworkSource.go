package awsmediaconnectalpha


// Bridge network source options.
//
// Example:
//   var stack Stack
//
//
//   productionNetwork := awsmediaconnectalpha.GatewayNetwork_Define(&GatewayNetworkDefineProps{
//   	CidrBlock: jsii.String("192.168.1.0/24"),
//   	Name: jsii.String("production-network"),
//   })
//
//   gateway := awsmediaconnectalpha.NewGateway(stack, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	EgressCidrBlocks: []*string{
//   		jsii.String("10.0.0.0/16"),
//   	},
//   	Networks: []GatewayNetwork{
//   		productionNetwork,
//   	},
//   })
//
//   ingressBridge := awsmediaconnectalpha.NewBridge(stack, jsii.String("MyIngressBridge"), &BridgeProps{
//   	BridgeName: jsii.String("my-ingress-bridge"),
//   	Config: awsmediaconnectalpha.BridgeConfiguration_Ingress(&IngressBridgeConfiguration{
//   		MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   		MaxOutputs: jsii.Number(2),
//   		NetworkSources: []BridgeNetworkInput{
//   			&BridgeNetworkInput{
//   				Name: jsii.String("on-prem-source"),
//   				Source: &BridgeNetworkSource{
//   					Protocol: awsmediaconnectalpha.BridgeProtocol_RTP(),
//   					Network: productionNetwork,
//   					MulticastIp: jsii.String("239.1.1.1"),
//   					Port: jsii.Number(5000),
//   				},
//   			},
//   		},
//   	}),
//   	Gateway: gateway,
//   })
//
// Experimental.
type BridgeNetworkSource struct {
	// The network source multicast IP.
	//
	// Must be a valid Multicast IP address.
	// Experimental.
	MulticastIp *string `field:"required" json:"multicastIp" yaml:"multicastIp"`
	// The gateway network this source listens on.
	//
	// Use {@link GatewayNetwork.define} to create the network and pass the same
	// instance to the gateway and to each source that uses it.
	// Experimental.
	Network GatewayNetwork `field:"required" json:"network" yaml:"network"`
	// The network source port.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network source protocol.
	// Experimental.
	Protocol BridgeProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// The setting related to the multicast source.
	//
	// The IP address of the source for source-specific multicast (SSM).
	// Default: - no multicast source IP.
	//
	// Experimental.
	MulticastSourceIp *string `field:"optional" json:"multicastSourceIp" yaml:"multicastSourceIp"`
}

