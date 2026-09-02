package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Font outline color for burn-in and DVB-Sub captions.
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
type CaptionOutlineColor interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionOutlineColor
type jsiiProxy_CaptionOutlineColor struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionOutlineColor) Value() *string {
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
func CaptionOutlineColor_Of(value *string) CaptionOutlineColor {
	_init_.Initialize()

	if err := validateCaptionOutlineColor_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionOutlineColor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionOutlineColor",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionOutlineColor_BLACK() CaptionOutlineColor {
	_init_.Initialize()
	var returns CaptionOutlineColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionOutlineColor",
		"BLACK",
		&returns,
	)
	return returns
}

func CaptionOutlineColor_BLUE() CaptionOutlineColor {
	_init_.Initialize()
	var returns CaptionOutlineColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionOutlineColor",
		"BLUE",
		&returns,
	)
	return returns
}

func CaptionOutlineColor_GREEN() CaptionOutlineColor {
	_init_.Initialize()
	var returns CaptionOutlineColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionOutlineColor",
		"GREEN",
		&returns,
	)
	return returns
}

func CaptionOutlineColor_RED() CaptionOutlineColor {
	_init_.Initialize()
	var returns CaptionOutlineColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionOutlineColor",
		"RED",
		&returns,
	)
	return returns
}

func CaptionOutlineColor_WHITE() CaptionOutlineColor {
	_init_.Initialize()
	var returns CaptionOutlineColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionOutlineColor",
		"WHITE",
		&returns,
	)
	return returns
}

func CaptionOutlineColor_YELLOW() CaptionOutlineColor {
	_init_.Initialize()
	var returns CaptionOutlineColor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionOutlineColor",
		"YELLOW",
		&returns,
	)
	return returns
}

