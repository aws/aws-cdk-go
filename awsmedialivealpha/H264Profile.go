package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 profile.
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
type H264Profile interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264Profile
type jsiiProxy_H264Profile struct {
	_ byte // padding
}

func (j *jsiiProxy_H264Profile) Value() *string {
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
func H264Profile_Of(value *string) H264Profile {
	_init_.Initialize()

	if err := validateH264Profile_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264Profile

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264Profile",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264Profile_BASELINE() H264Profile {
	_init_.Initialize()
	var returns H264Profile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Profile",
		"BASELINE",
		&returns,
	)
	return returns
}

func H264Profile_HIGH() H264Profile {
	_init_.Initialize()
	var returns H264Profile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Profile",
		"HIGH",
		&returns,
	)
	return returns
}

func H264Profile_HIGH_10BIT() H264Profile {
	_init_.Initialize()
	var returns H264Profile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Profile",
		"HIGH_10BIT",
		&returns,
	)
	return returns
}

func H264Profile_HIGH_422() H264Profile {
	_init_.Initialize()
	var returns H264Profile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Profile",
		"HIGH_422",
		&returns,
	)
	return returns
}

func H264Profile_HIGH_422_10BIT() H264Profile {
	_init_.Initialize()
	var returns H264Profile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Profile",
		"HIGH_422_10BIT",
		&returns,
	)
	return returns
}

func H264Profile_MAIN() H264Profile {
	_init_.Initialize()
	var returns H264Profile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Profile",
		"MAIN",
		&returns,
	)
	return returns
}

