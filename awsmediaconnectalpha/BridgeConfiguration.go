package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Bridge configuration to set ingress and egress options on the bridge.
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
type BridgeConfiguration interface {
}

// The jsii proxy struct for BridgeConfiguration
type jsiiProxy_BridgeConfiguration struct {
	_ byte // padding
}

// An egress bridge is a cloud-to-ground bridge.
//
// The content comes from an existing MediaConnect flow and is delivered to your premises.
// Experimental.
func BridgeConfiguration_Egress(config *EgressBridgeConfiguration) BridgeConfiguration {
	_init_.Initialize()

	if err := validateBridgeConfiguration_EgressParameters(config); err != nil {
		panic(err)
	}
	var returns BridgeConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeConfiguration",
		"egress",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// An ingress bridge is a ground-to-cloud bridge.
//
// The content originates at your premises and is delivered to the cloud.
// Experimental.
func BridgeConfiguration_Ingress(config *IngressBridgeConfiguration) BridgeConfiguration {
	_init_.Initialize()

	if err := validateBridgeConfiguration_IngressParameters(config); err != nil {
		panic(err)
	}
	var returns BridgeConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeConfiguration",
		"ingress",
		[]interface{}{config},
		&returns,
	)

	return returns
}

