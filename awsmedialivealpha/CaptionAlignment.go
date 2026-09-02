package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Caption alignment for burn-in and DVB-Sub outputs.
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
type CaptionAlignment interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionAlignment
type jsiiProxy_CaptionAlignment struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionAlignment) Value() *string {
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
func CaptionAlignment_Of(value *string) CaptionAlignment {
	_init_.Initialize()

	if err := validateCaptionAlignment_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionAlignment

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionAlignment",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionAlignment_CENTERED() CaptionAlignment {
	_init_.Initialize()
	var returns CaptionAlignment
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionAlignment",
		"CENTERED",
		&returns,
	)
	return returns
}

func CaptionAlignment_LEFT() CaptionAlignment {
	_init_.Initialize()
	var returns CaptionAlignment
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionAlignment",
		"LEFT",
		&returns,
	)
	return returns
}

func CaptionAlignment_SMART() CaptionAlignment {
	_init_.Initialize()
	var returns CaptionAlignment
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionAlignment",
		"SMART",
		&returns,
	)
	return returns
}

