package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Network interface options.
//
// Example:
//   var stack Stack
//   var role IRole
//   var securityGroup ISecurityGroup
//   var subnet ISubnet
//
//
//   efaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("efa-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Subnet: subnet,
//   	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_EFA(),
//   })
//
//   videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
//   	MediaStreamId: jsii.Number(1),
//   	MediaStreamName: jsii.String("video"),
//   	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_HD_1080P(),
//   	Fmtp: &FmtpVideo{
//   		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_29_97(),
//   		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
//   	},
//   })
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyCdiFlow"), &FlowProps{
//   	FlowSize: awsmediaconnectalpha.FlowSize_LARGE_4X(),
//   	 // Required for CDI and JPEG XS
//   	VpcInterfaces: []VpcInterfaceConfig{
//   		efaInterface,
//   	},
//   	MediaStreams: []MediaStream{
//   		videoStream,
//   	},
//   	Source: awsmediaconnectalpha.SourceConfiguration_Cdi(&SourceCdi{
//   		FlowSourceName: jsii.String("cdi-source"),
//   		VpcInterface: efaInterface,
//   		Port: jsii.Number(5000),
//   		MaxSyncBuffer: jsii.Number(100),
//   		MediaStreamSourceConfigurations: []MediaStreamSourceConfigurationCdi{
//   			&MediaStreamSourceConfigurationCdi{
//   				Encoding: awsmediaconnectalpha.Encoding_RAW(),
//   				MediaStream: videoStream,
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type NetworkInterface interface {
	// The network interface string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkInterface
type jsiiProxy_NetworkInterface struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkInterface) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom network interface value.
// Experimental.
func NetworkInterface_Of(value *string) NetworkInterface {
	_init_.Initialize()

	if err := validateNetworkInterface_OfParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkInterface

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.NetworkInterface",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkInterface_EFA() NetworkInterface {
	_init_.Initialize()
	var returns NetworkInterface
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.NetworkInterface",
		"EFA",
		&returns,
	)
	return returns
}

func NetworkInterface_ENA() NetworkInterface {
	_init_.Initialize()
	var returns NetworkInterface
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.NetworkInterface",
		"ENA",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkInterface) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

