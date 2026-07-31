package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for Video Range.
//
// Example:
//   var stack Stack
//   var role IRole
//   var sg1 ISecurityGroup
//   var sg2 ISecurityGroup
//   var subnet ISubnet
//
//
//   efaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("efa-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		sg1,
//   	},
//   	Subnet: subnet,
//   	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_EFA(),
//   })
//
//   enaInterface := awsmediaconnectalpha.VpcInterface_Define(&VpcInterfaceDefineProps{
//   	VpcInterfaceName: jsii.String("ena-interface"),
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		sg2,
//   	},
//   	Subnet: subnet,
//   	NetworkInterfaceType: awsmediaconnectalpha.NetworkInterface_ENA(),
//   })
//
//   videoStream := awsmediaconnectalpha.MediaStream_Video(&MediaStreamVideo{
//   	MediaStreamId: jsii.Number(1),
//   	MediaStreamName: jsii.String("video"),
//   	VideoFormat: awsmediaconnectalpha.MediaVideoFormat_UHD_2160P(),
//   	Fmtp: &FmtpVideo{
//   		ExactFramerate: awsmediaconnectalpha.Framerate_FPS_59_94(),
//   		Par: awsmediaconnectalpha.PixelAspectRatio_SQUARE(),
//   		Colorimetry: awsmediaconnectalpha.Colorimetry_BT2020(),
//   		VideoRange: awsmediaconnectalpha.VideoRange_FULL(),
//   		ScanMode: awsmediaconnectalpha.ScanMode_PROGRESSIVE(),
//   		Tcs: awsmediaconnectalpha.Tcs_PQ(),
//   	},
//   })
//
//   flow := awsmediaconnectalpha.NewFlow(stack, jsii.String("MyJpegXsFlow"), &FlowProps{
//   	FlowSize: awsmediaconnectalpha.FlowSize_LARGE_4X(),
//   	 // Required for JPEG XS
//   	VpcInterfaces: []VpcInterfaceConfig{
//   		efaInterface,
//   		enaInterface,
//   	},
//   	MediaStreams: []MediaStream{
//   		videoStream,
//   	},
//   	Source: awsmediaconnectalpha.SourceConfiguration_JpegXs(&SourceJpegXs{
//   		FlowSourceName: jsii.String("jpegxs-source"),
//   		MaxSyncBuffer: jsii.Number(100),
//   		MediaStreamSourceConfigurations: []MediaStreamSourceConfigurationJpegXs{
//   			&MediaStreamSourceConfigurationJpegXs{
//   				Encoding: awsmediaconnectalpha.Encoding_JXSV(),
//   				Port: jsii.Number(5000),
//   				InputInterface: []VpcInterfaceConfig{
//   					efaInterface,
//   					enaInterface,
//   				},
//   				 // 2 interfaces for redundancy
//   				MediaStream: videoStream,
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type VideoRange interface {
	// The video range string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VideoRange
type jsiiProxy_VideoRange struct {
	_ byte // padding
}

func (j *jsiiProxy_VideoRange) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom video range value.
// Experimental.
func VideoRange_Of(value *string) VideoRange {
	_init_.Initialize()

	if err := validateVideoRange_OfParameters(value); err != nil {
		panic(err)
	}
	var returns VideoRange

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.VideoRange",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func VideoRange_FULL() VideoRange {
	_init_.Initialize()
	var returns VideoRange
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.VideoRange",
		"FULL",
		&returns,
	)
	return returns
}

func VideoRange_FULLPROTECT() VideoRange {
	_init_.Initialize()
	var returns VideoRange
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.VideoRange",
		"FULLPROTECT",
		&returns,
	)
	return returns
}

func VideoRange_NARROW() VideoRange {
	_init_.Initialize()
	var returns VideoRange
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.VideoRange",
		"NARROW",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VideoRange) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

