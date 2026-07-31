package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for Media Video format.
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
type MediaVideoFormat interface {
	// The video format string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaVideoFormat
type jsiiProxy_MediaVideoFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_MediaVideoFormat) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom video format value.
// Experimental.
func MediaVideoFormat_Of(value *string) MediaVideoFormat {
	_init_.Initialize()

	if err := validateMediaVideoFormat_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MediaVideoFormat

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.MediaVideoFormat",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MediaVideoFormat_HD_1080I() MediaVideoFormat {
	_init_.Initialize()
	var returns MediaVideoFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.MediaVideoFormat",
		"HD_1080I",
		&returns,
	)
	return returns
}

func MediaVideoFormat_HD_1080P() MediaVideoFormat {
	_init_.Initialize()
	var returns MediaVideoFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.MediaVideoFormat",
		"HD_1080P",
		&returns,
	)
	return returns
}

func MediaVideoFormat_HD_720P() MediaVideoFormat {
	_init_.Initialize()
	var returns MediaVideoFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.MediaVideoFormat",
		"HD_720P",
		&returns,
	)
	return returns
}

func MediaVideoFormat_SD_480P() MediaVideoFormat {
	_init_.Initialize()
	var returns MediaVideoFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.MediaVideoFormat",
		"SD_480P",
		&returns,
	)
	return returns
}

func MediaVideoFormat_UHD_2160P() MediaVideoFormat {
	_init_.Initialize()
	var returns MediaVideoFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.MediaVideoFormat",
		"UHD_2160P",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaVideoFormat) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

