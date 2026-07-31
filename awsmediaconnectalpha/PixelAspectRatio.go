package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The pixel aspect ratio (PAR) of the video.
//
// Use the predefined constants for standard ratios, or {@link PixelAspectRatio.of} for
// a custom ratio.
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
type PixelAspectRatio interface {
	// Returns the string value in `numerator:denominator` form.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PixelAspectRatio
type jsiiProxy_PixelAspectRatio struct {
	_ byte // padding
}

// Define a pixel aspect ratio.
// Experimental.
func PixelAspectRatio_Of(numerator *float64, denominator *float64) PixelAspectRatio {
	_init_.Initialize()

	if err := validatePixelAspectRatio_OfParameters(numerator, denominator); err != nil {
		panic(err)
	}
	var returns PixelAspectRatio

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.PixelAspectRatio",
		"of",
		[]interface{}{numerator, denominator},
		&returns,
	)

	return returns
}

func PixelAspectRatio_SQUARE() PixelAspectRatio {
	_init_.Initialize()
	var returns PixelAspectRatio
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.PixelAspectRatio",
		"SQUARE",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PixelAspectRatio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

