package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Color space settings for H.265 video.
//
// Example:
//   hdr := medialive.EncodeConfiguration_Video(&VideoEncodeProps{
//   	Name: jsii.String("h265_hdr"),
//   	Codec: medialive.VideoCodecSettings_H265(&H265SettingsProps{
//   		RateControl: medialive.H265RateControl_Qvbr(&QvbrRateControlProps{
//   			MaxBitrate: awscdk.Bitrate_Mbps(jsii.Number(8)),
//   			QvbrQualityLevel: jsii.Number(8),
//   		}),
//   		Framerate: medialive.Framerate_FPS_30(),
//   		SceneChangeDetect: medialive.H265SceneChangeDetect_ENABLED(),
//   		ColorSpaceSettings: medialive.H265ColorSpaceSettings_Hlg2020(),
//   	}),
//   	Width: jsii.Number(1920),
//   	Height: jsii.Number(1080),
//   })
//
// Experimental.
type H265ColorSpaceSettings interface {
}

// The jsii proxy struct for H265ColorSpaceSettings
type jsiiProxy_H265ColorSpaceSettings struct {
	_ byte // padding
}

// Dolby Vision 8.1 color space.
// Experimental.
func H265ColorSpaceSettings_DolbyVision81() H265ColorSpaceSettings {
	_init_.Initialize()

	var returns H265ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265ColorSpaceSettings",
		"dolbyVision81",
		nil, // no parameters
		&returns,
	)

	return returns
}

// HDR10 color space.
// Experimental.
func H265ColorSpaceSettings_Hdr10(props *Hdr10SettingsProps) H265ColorSpaceSettings {
	_init_.Initialize()

	if err := validateH265ColorSpaceSettings_Hdr10Parameters(props); err != nil {
		panic(err)
	}
	var returns H265ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265ColorSpaceSettings",
		"hdr10",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// HLG 2020 color space.
// Experimental.
func H265ColorSpaceSettings_Hlg2020() H265ColorSpaceSettings {
	_init_.Initialize()

	var returns H265ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265ColorSpaceSettings",
		"hlg2020",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Pass through the source color space with no conversion.
// Experimental.
func H265ColorSpaceSettings_Passthrough() H265ColorSpaceSettings {
	_init_.Initialize()

	var returns H265ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265ColorSpaceSettings",
		"passthrough",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Convert to Rec.601 color space.
// Experimental.
func H265ColorSpaceSettings_Rec601() H265ColorSpaceSettings {
	_init_.Initialize()

	var returns H265ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265ColorSpaceSettings",
		"rec601",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Convert to Rec.709 color space.
// Experimental.
func H265ColorSpaceSettings_Rec709() H265ColorSpaceSettings {
	_init_.Initialize()

	var returns H265ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265ColorSpaceSettings",
		"rec709",
		nil, // no parameters
		&returns,
	)

	return returns
}

