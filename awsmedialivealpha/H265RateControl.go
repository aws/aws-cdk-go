package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 rate control. Use the static factory methods to create.
//
// Example:
//   // H.264
//   h264 := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("h264_720p"),
//   	Codec: medialive.VideoCodecSettings_H264(&H264SettingsProps{
//   		RateControl: medialive.H264RateControl_Cbr(&CbrRateControlProps{
//   			Bitrate: awscdk.Bitrate_Mbps(jsii.Number(3)),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   		Profile: medialive.H264Profile_HIGH(),
//   	}),
//   	Width: jsii.Number(1280),
//   	Height: jsii.Number(720),
//   })
//
//   // H.265
//   h265 := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("h265_1080p"),
//   	Codec: medialive.VideoCodecSettings_H265(&H265SettingsProps{
//   		RateControl: medialive.H265RateControl_Qvbr(&QvbrRateControlProps{
//   			MaxBitrate: awscdk.Bitrate_*Mbps(jsii.Number(5)),
//   			QvbrQualityLevel: jsii.Number(7),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   		Profile: medialive.H265Profile_MAIN(),
//   		Tier: medialive.H265Tier_HIGH(),
//   	}),
//   	Width: jsii.Number(1920),
//   	Height: jsii.Number(1080),
//   })
//
// Experimental.
type H265RateControl interface {
}

// The jsii proxy struct for H265RateControl
type jsiiProxy_H265RateControl struct {
	_ byte // padding
}

// Constant bitrate.
// Experimental.
func H265RateControl_Cbr(props *CbrRateControlProps) H265RateControl {
	_init_.Initialize()

	if err := validateH265RateControl_CbrParameters(props); err != nil {
		panic(err)
	}
	var returns H265RateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265RateControl",
		"cbr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Quality-defined variable bitrate.
// Experimental.
func H265RateControl_Qvbr(props *QvbrRateControlProps) H265RateControl {
	_init_.Initialize()

	if err := validateH265RateControl_QvbrParameters(props); err != nil {
		panic(err)
	}
	var returns H265RateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265RateControl",
		"qvbr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Variable bitrate.
// Experimental.
func H265RateControl_Vbr(props *VbrRateControlProps) H265RateControl {
	_init_.Initialize()

	if err := validateH265RateControl_VbrParameters(props); err != nil {
		panic(err)
	}
	var returns H265RateControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265RateControl",
		"vbr",
		[]interface{}{props},
		&returns,
	)

	return returns
}

