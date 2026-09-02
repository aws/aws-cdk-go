package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 scene change detection.
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
type H265SceneChangeDetect interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265SceneChangeDetect
type jsiiProxy_H265SceneChangeDetect struct {
	_ byte // padding
}

func (j *jsiiProxy_H265SceneChangeDetect) Value() *string {
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
func H265SceneChangeDetect_Of(value *string) H265SceneChangeDetect {
	_init_.Initialize()

	if err := validateH265SceneChangeDetect_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265SceneChangeDetect

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265SceneChangeDetect",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265SceneChangeDetect_DISABLED() H265SceneChangeDetect {
	_init_.Initialize()
	var returns H265SceneChangeDetect
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265SceneChangeDetect",
		"DISABLED",
		&returns,
	)
	return returns
}

func H265SceneChangeDetect_ENABLED() H265SceneChangeDetect {
	_init_.Initialize()
	var returns H265SceneChangeDetect
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265SceneChangeDetect",
		"ENABLED",
		&returns,
	)
	return returns
}

