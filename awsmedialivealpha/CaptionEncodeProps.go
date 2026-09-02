package awsmedialivealpha


// Properties for a caption encode configuration.
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
type CaptionEncodeProps struct {
	// The name of the caption selector in the input to use as the source.
	// Experimental.
	CaptionSelectorName *string `field:"required" json:"captionSelectorName" yaml:"captionSelectorName"`
	// The output caption format.
	//
	// Use the `CaptionDestination` factory methods (e.g.
	// `CaptionDestination.burnIn()`, `.webvtt()`, `.embedded()`).
	// Experimental.
	Destination CaptionDestination `field:"required" json:"destination" yaml:"destination"`
	// A unique name for this caption encode.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether this caption track implements accessibility features.
	// Default: - The captions do not implement accessibility features.
	//
	// Experimental.
	Accessibility CaptionAccessibility `field:"optional" json:"accessibility" yaml:"accessibility"`
	// The DASH roles to assign to this captions output.
	//
	// Applies only when the output is configured
	// for DVB DASH accessibility signaling.
	// Default: - no DASH roles.
	//
	// Experimental.
	CaptionDashRoles *[]CaptionDashRole `field:"optional" json:"captionDashRoles" yaml:"captionDashRoles"`
	// DVB DASH accessibility signaling for this captions output.
	// Default: - no DVB DASH accessibility signaling.
	//
	// Experimental.
	DvbDashAccessibility DvbDashAccessibility `field:"optional" json:"dvbDashAccessibility" yaml:"dvbDashAccessibility"`
	// The ISO 639-2 language code for the captions (e.g. 'eng', 'spa').
	// Default: - no language code.
	//
	// Experimental.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Human-readable description of the captions (e.g. 'English', 'Spanish').
	// Default: - no language description.
	//
	// Experimental.
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
}

