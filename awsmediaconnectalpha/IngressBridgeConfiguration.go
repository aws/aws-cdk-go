package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Ingress bridge configuration.
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
type IngressBridgeConfiguration struct {
	// The maximum expected bitrate (in bps) of the bridge.
	// Experimental.
	MaxBitrate awscdk.Bitrate `field:"required" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum number of outputs on the ingress bridge.
	// Experimental.
	MaxOutputs *float64 `field:"required" json:"maxOutputs" yaml:"maxOutputs"`
	// The network sources for the ingress bridge.
	// Experimental.
	NetworkSources *[]*BridgeNetworkInput `field:"required" json:"networkSources" yaml:"networkSources"`
}

