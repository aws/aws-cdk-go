package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The output caption format for a caption encode.
//
// Use the static factory methods to select one of
// the supported destination types.
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
type CaptionDestination interface {
}

// The jsii proxy struct for CaptionDestination
type jsiiProxy_CaptionDestination struct {
	_ byte // padding
}

// Experimental.
func NewCaptionDestination_Override(c CaptionDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		nil, // no parameters
		c,
	)
}

// ARIB captions.
// Experimental.
func CaptionDestination_Arib() CaptionDestination {
	_init_.Initialize()

	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"arib",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Burned-in captions rendered into the video.
// Experimental.
func CaptionDestination_BurnIn(props *BurnInDestinationProps) CaptionDestination {
	_init_.Initialize()

	if err := validateCaptionDestination_BurnInParameters(props); err != nil {
		panic(err)
	}
	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"burnIn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// DVB-Sub bitmap captions.
// Experimental.
func CaptionDestination_DvbSub(props *DvbSubDestinationProps) CaptionDestination {
	_init_.Initialize()

	if err := validateCaptionDestination_DvbSubParameters(props); err != nil {
		panic(err)
	}
	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"dvbSub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// EBU-TT-D sidecar captions.
// Experimental.
func CaptionDestination_EbuTtD(props *EbuTtDDestinationProps) CaptionDestination {
	_init_.Initialize()

	if err := validateCaptionDestination_EbuTtDParameters(props); err != nil {
		panic(err)
	}
	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"ebuTtD",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Embedded (CEA-608/708) captions.
// Experimental.
func CaptionDestination_Embedded() CaptionDestination {
	_init_.Initialize()

	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"embedded",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Embedded plus SCTE-20 captions.
// Experimental.
func CaptionDestination_EmbeddedPlusScte20() CaptionDestination {
	_init_.Initialize()

	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"embeddedPlusScte20",
		nil, // no parameters
		&returns,
	)

	return returns
}

// RTMP CaptionInfo captions.
// Experimental.
func CaptionDestination_RtmpCaptionInfo() CaptionDestination {
	_init_.Initialize()

	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"rtmpCaptionInfo",
		nil, // no parameters
		&returns,
	)

	return returns
}

// SCTE-20 plus embedded captions.
// Experimental.
func CaptionDestination_Scte20PlusEmbedded() CaptionDestination {
	_init_.Initialize()

	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"scte20PlusEmbedded",
		nil, // no parameters
		&returns,
	)

	return returns
}

// SMPTE-TT sidecar captions.
// Experimental.
func CaptionDestination_SmpteTt() CaptionDestination {
	_init_.Initialize()

	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"smpteTt",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Teletext captions.
// Experimental.
func CaptionDestination_Teletext() CaptionDestination {
	_init_.Initialize()

	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"teletext",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TTML sidecar captions.
// Experimental.
func CaptionDestination_Ttml(props *TtmlDestinationProps) CaptionDestination {
	_init_.Initialize()

	if err := validateCaptionDestination_TtmlParameters(props); err != nil {
		panic(err)
	}
	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"ttml",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// WebVTT sidecar captions.
// Experimental.
func CaptionDestination_Webvtt(props *WebvttDestinationProps) CaptionDestination {
	_init_.Initialize()

	if err := validateCaptionDestination_WebvttParameters(props); err != nil {
		panic(err)
	}
	var returns CaptionDestination

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionDestination",
		"webvtt",
		[]interface{}{props},
		&returns,
	)

	return returns
}

