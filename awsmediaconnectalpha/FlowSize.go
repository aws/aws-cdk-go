package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for Flow Size.
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
type FlowSize interface {
	// The flow size string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FlowSize
type jsiiProxy_FlowSize struct {
	_ byte // padding
}

func (j *jsiiProxy_FlowSize) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom flow size value.
// Experimental.
func FlowSize_Of(value *string) FlowSize {
	_init_.Initialize()

	if err := validateFlowSize_OfParameters(value); err != nil {
		panic(err)
	}
	var returns FlowSize

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.FlowSize",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func FlowSize_LARGE() FlowSize {
	_init_.Initialize()
	var returns FlowSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.FlowSize",
		"LARGE",
		&returns,
	)
	return returns
}

func FlowSize_LARGE_4X() FlowSize {
	_init_.Initialize()
	var returns FlowSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.FlowSize",
		"LARGE_4X",
		&returns,
	)
	return returns
}

func FlowSize_MEDIUM() FlowSize {
	_init_.Initialize()
	var returns FlowSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.FlowSize",
		"MEDIUM",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FlowSize) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

