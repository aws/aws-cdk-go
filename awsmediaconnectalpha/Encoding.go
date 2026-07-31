package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Encoding options.
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
type Encoding interface {
	// The encoding string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Encoding
type jsiiProxy_Encoding struct {
	_ byte // padding
}

func (j *jsiiProxy_Encoding) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom encoding value.
// Experimental.
func Encoding_Of(value *string) Encoding {
	_init_.Initialize()

	if err := validateEncoding_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Encoding

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.Encoding",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Encoding_JXSV() Encoding {
	_init_.Initialize()
	var returns Encoding
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Encoding",
		"JXSV",
		&returns,
	)
	return returns
}

func Encoding_PCM() Encoding {
	_init_.Initialize()
	var returns Encoding
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Encoding",
		"PCM",
		&returns,
	)
	return returns
}

func Encoding_RAW() Encoding {
	_init_.Initialize()
	var returns Encoding
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Encoding",
		"RAW",
		&returns,
	)
	return returns
}

func Encoding_SMPTE291() Encoding {
	_init_.Initialize()
	var returns Encoding
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Encoding",
		"SMPTE291",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Encoding) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

