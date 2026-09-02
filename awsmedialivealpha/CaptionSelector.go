package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A caption selector that identifies which captions to extract from the input.
//
// Create with
// the static factory methods — one per caption source format.
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
type CaptionSelector interface {
	// The name of this caption selector, used to associate it with caption outputs.
	//
	// Unique
	// within a channel.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for CaptionSelector
type jsiiProxy_CaptionSelector struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionSelector) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewCaptionSelector_Override(c CaptionSelector, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		[]interface{}{name},
		c,
	)
}

// Select ancillary captions.
// Experimental.
func CaptionSelector_Ancillary(name *string, options *AncillaryCaptionSourceOptions) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_AncillaryParameters(name, options); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"ancillary",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Select ARIB captions.
// Experimental.
func CaptionSelector_Arib(name *string) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_AribParameters(name); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"arib",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Select captions by language code (no specific source format).
// Experimental.
func CaptionSelector_ByLanguage(name *string, languageCode *string) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_ByLanguageParameters(name, languageCode); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"byLanguage",
		[]interface{}{name, languageCode},
		&returns,
	)

	return returns
}

// Select DVB-Sub (image-based) captions.
// Experimental.
func CaptionSelector_DvbSub(name *string, options *DvbSubCaptionSourceOptions) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_DvbSubParameters(name, options); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"dvbSub",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Select embedded (CEA-608/708) captions.
// Experimental.
func CaptionSelector_Embedded(name *string, options *EmbeddedCaptionSourceOptions) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_EmbeddedParameters(name, options); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"embedded",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Select SCTE-20 captions.
// Experimental.
func CaptionSelector_Scte20(name *string, options *Scte20CaptionSourceOptions) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_Scte20Parameters(name, options); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"scte20",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Select SCTE-27 (image-based) captions.
// Experimental.
func CaptionSelector_Scte27(name *string, options *Scte27CaptionSourceOptions) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_Scte27Parameters(name, options); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"scte27",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Select smart subtitles generated by Elemental Inference.
// Experimental.
func CaptionSelector_SmartSubtitle(name *string, options *SmartSubtitleSourceOptions) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_SmartSubtitleParameters(name, options); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"smartSubtitle",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Select Teletext captions.
// Experimental.
func CaptionSelector_Teletext(name *string, options *TeletextCaptionSourceOptions) CaptionSelector {
	_init_.Initialize()

	if err := validateCaptionSelector_TeletextParameters(name, options); err != nil {
		panic(err)
	}
	var returns CaptionSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSelector",
		"teletext",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

