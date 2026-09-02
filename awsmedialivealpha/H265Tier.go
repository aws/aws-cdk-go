package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 tier.
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
type H265Tier interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265Tier
type jsiiProxy_H265Tier struct {
	_ byte // padding
}

func (j *jsiiProxy_H265Tier) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func H265Tier_Of(value *string) H265Tier {
	_init_.Initialize()

	if err := validateH265Tier_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265Tier

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265Tier",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265Tier_HIGH() H265Tier {
	_init_.Initialize()
	var returns H265Tier
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Tier",
		"HIGH",
		&returns,
	)
	return returns
}

func H265Tier_MAIN() H265Tier {
	_init_.Initialize()
	var returns H265Tier
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Tier",
		"MAIN",
		&returns,
	)
	return returns
}

