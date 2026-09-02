package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Font size for burn-in and DVB-Sub captions.
//
// Use {@link CaptionFontSize.AUTO} to scale the font size with the output resolution, or
// {@link CaptionFontSize.of} to set an exact size in points.
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
type CaptionFontSize interface {
}

// The jsii proxy struct for CaptionFontSize
type jsiiProxy_CaptionFontSize struct {
	_ byte // padding
}

// An exact font size, in points.
// Experimental.
func CaptionFontSize_Of(points *float64) CaptionFontSize {
	_init_.Initialize()

	if err := validateCaptionFontSize_OfParameters(points); err != nil {
		panic(err)
	}
	var returns CaptionFontSize

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionFontSize",
		"of",
		[]interface{}{points},
		&returns,
	)

	return returns
}

func CaptionFontSize_AUTO() CaptionFontSize {
	_init_.Initialize()
	var returns CaptionFontSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionFontSize",
		"AUTO",
		&returns,
	)
	return returns
}

