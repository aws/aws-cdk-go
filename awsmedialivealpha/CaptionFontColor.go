package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Font color for burn-in and DVB-Sub captions.
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
type CaptionFontColor interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionFontColor
type jsiiProxy_CaptionFontColor struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionFontColor) Value() *string {
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
func CaptionFontColor_Of(value *string) CaptionFontColor {
	_init_.Initialize()

	if err := validateCaptionFontColor_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionFontColor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionFontColor",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionFontColor_BLACK() CaptionFontColor {
	_init_.Initialize()
	var returns CaptionFontColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionFontColor",
		"BLACK",
		&returns,
	)
	return returns
}

func CaptionFontColor_BLUE() CaptionFontColor {
	_init_.Initialize()
	var returns CaptionFontColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionFontColor",
		"BLUE",
		&returns,
	)
	return returns
}

func CaptionFontColor_GREEN() CaptionFontColor {
	_init_.Initialize()
	var returns CaptionFontColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionFontColor",
		"GREEN",
		&returns,
	)
	return returns
}

func CaptionFontColor_RED() CaptionFontColor {
	_init_.Initialize()
	var returns CaptionFontColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionFontColor",
		"RED",
		&returns,
	)
	return returns
}

func CaptionFontColor_WHITE() CaptionFontColor {
	_init_.Initialize()
	var returns CaptionFontColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionFontColor",
		"WHITE",
		&returns,
	)
	return returns
}

func CaptionFontColor_YELLOW() CaptionFontColor {
	_init_.Initialize()
	var returns CaptionFontColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionFontColor",
		"YELLOW",
		&returns,
	)
	return returns
}

