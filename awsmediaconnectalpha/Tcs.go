package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for Tcs.
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
type Tcs interface {
	// The TCS string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Tcs
type jsiiProxy_Tcs struct {
	_ byte // padding
}

func (j *jsiiProxy_Tcs) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom TCS value.
// Experimental.
func Tcs_Of(value *string) Tcs {
	_init_.Initialize()

	if err := validateTcs_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Tcs

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Tcs_BT2100LINHLG() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"BT2100LINHLG",
		&returns,
	)
	return returns
}

func Tcs_BT2100LINPQ() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"BT2100LINPQ",
		&returns,
	)
	return returns
}

func Tcs_DENSITY() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"DENSITY",
		&returns,
	)
	return returns
}

func Tcs_HLG() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"HLG",
		&returns,
	)
	return returns
}

func Tcs_LINEAR() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"LINEAR",
		&returns,
	)
	return returns
}

func Tcs_PQ() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"PQ",
		&returns,
	)
	return returns
}

func Tcs_SDR() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"SDR",
		&returns,
	)
	return returns
}

func Tcs_ST2065_1() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"ST2065_1",
		&returns,
	)
	return returns
}

func Tcs_ST428_1() Tcs {
	_init_.Initialize()
	var returns Tcs
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.Tcs",
		"ST428_1",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_Tcs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

