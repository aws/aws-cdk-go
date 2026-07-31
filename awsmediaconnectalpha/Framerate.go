package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A video frame rate expressed as a rational number (numerator/denominator).
//
// Use the predefined constants for standard rates, or {@link Framerate.of} for a custom rate.
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
type Framerate interface {
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Framerate
type jsiiProxy_Framerate struct {
	_ byte // padding
}

// Define a custom frame rate.
// Experimental.
func Framerate_Of(numerator *float64, denominator *float64) Framerate {
	_init_.Initialize()

	if err := validateFramerate_OfParameters(numerator, denominator); err != nil {
		panic(err)
	}
	var returns Framerate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		"of",
		[]interface{}{numerator, denominator},
		&returns,
	)

	return returns
}

func Framerate_FPS_24() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		"FPS_24",
		&returns,
	)
	return returns
}

func Framerate_FPS_25() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		"FPS_25",
		&returns,
	)
	return returns
}

func Framerate_FPS_29_97() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		"FPS_29_97",
		&returns,
	)
	return returns
}

func Framerate_FPS_30() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		"FPS_30",
		&returns,
	)
	return returns
}

func Framerate_FPS_50() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		"FPS_50",
		&returns,
	)
	return returns
}

func Framerate_FPS_59_94() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		"FPS_59_94",
		&returns,
	)
	return returns
}

func Framerate_FPS_60() Framerate {
	_init_.Initialize()
	var returns Framerate
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Framerate",
		"FPS_60",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_Framerate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

