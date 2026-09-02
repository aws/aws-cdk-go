package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base interface for an encode configuration (video, audio, or caption).
//
// The same EncodeConfiguration instance can be shared across multiple output groups within a channel.
// The channel automatically deduplicates encode descriptions by name at synth time.
//
// Example:
//   // Define a caption selector on the input attachment (see Input Attachment Settings below)
//   captionSelector := medialive.CaptionSelector_Embedded(jsii.String("captions"))
//
//   // WebVTT captions — packaged alongside the video encode in the same output
//   webvtt := medialive.EncodeConfiguration_Caption(&CaptionEncodeProps{
//   	Name: jsii.String("eng_webvtt"),
//   	CaptionSelectorName: captionSelector.Name,
//   	LanguageCode: jsii.String("eng"),
//   	LanguageDescription: jsii.String("English"),
//   	Destination: medialive.CaptionDestination_Webvtt(),
//   })
//
//   // Burned-in captions — rendered into the video, styled via the burn-in options
//   burnIn := medialive.EncodeConfiguration_Caption(&CaptionEncodeProps{
//   	Name: jsii.String("eng_burnin"),
//   	CaptionSelectorName: captionSelector.*Name,
//   	Destination: medialive.CaptionDestination_BurnIn(&BurnInDestinationProps{
//   		Alignment: medialive.CaptionAlignment_CENTERED(),
//   		FontColor: medialive.CaptionFontColor_WHITE(),
//   		OutlineColor: medialive.CaptionOutlineColor_BLACK(),
//   		FontSize: medialive.CaptionFontSize_AUTO(),
//   	}),
//   })
//
// Experimental.
type EncodeConfiguration interface {
	// The unique name for this encode, used to reference it from outputs.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for EncodeConfiguration
type jsiiProxy_EncodeConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_EncodeConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewEncodeConfiguration_Override(e EncodeConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.EncodeConfiguration",
		nil, // no parameters
		e,
	)
}

// Create an audio encode configuration.
// Experimental.
func EncodeConfiguration_Audio(props *AudioEncodeProps) EncodeConfiguration {
	_init_.Initialize()

	if err := validateEncodeConfiguration_AudioParameters(props); err != nil {
		panic(err)
	}
	var returns EncodeConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.EncodeConfiguration",
		"audio",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a caption encode configuration.
// Experimental.
func EncodeConfiguration_Caption(props *CaptionEncodeProps) EncodeConfiguration {
	_init_.Initialize()

	if err := validateEncodeConfiguration_CaptionParameters(props); err != nil {
		panic(err)
	}
	var returns EncodeConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.EncodeConfiguration",
		"caption",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a video encode configuration.
// Experimental.
func EncodeConfiguration_Video(props *VideoEncodeProps) EncodeConfiguration {
	_init_.Initialize()

	if err := validateEncodeConfiguration_VideoParameters(props); err != nil {
		panic(err)
	}
	var returns EncodeConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.EncodeConfiguration",
		"video",
		[]interface{}{props},
		&returns,
	)

	return returns
}

