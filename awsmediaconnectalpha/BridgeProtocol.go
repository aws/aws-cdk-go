package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for bridge network source and output protocols.
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
type BridgeProtocol interface {
	// The bridge protocol string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BridgeProtocol
type jsiiProxy_BridgeProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_BridgeProtocol) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom bridge protocol value.
// Experimental.
func BridgeProtocol_Of(value *string) BridgeProtocol {
	_init_.Initialize()

	if err := validateBridgeProtocol_OfParameters(value); err != nil {
		panic(err)
	}
	var returns BridgeProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeProtocol",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func BridgeProtocol_RTP() BridgeProtocol {
	_init_.Initialize()
	var returns BridgeProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeProtocol",
		"RTP",
		&returns,
	)
	return returns
}

func BridgeProtocol_RTP_FEC() BridgeProtocol {
	_init_.Initialize()
	var returns BridgeProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeProtocol",
		"RTP_FEC",
		&returns,
	)
	return returns
}

func BridgeProtocol_UDP() BridgeProtocol {
	_init_.Initialize()
	var returns BridgeProtocol
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.BridgeProtocol",
		"UDP",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BridgeProtocol) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

