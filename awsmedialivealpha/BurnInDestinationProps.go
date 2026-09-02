package awsmedialivealpha


// Properties for burn-in captions.
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
type BurnInDestinationProps struct {
	// Caption alignment.
	//
	// With explicit x/y positions, the font is justified relative to them.
	// Default: CaptionAlignment.CENTERED
	//
	// Experimental.
	Alignment CaptionAlignment `field:"optional" json:"alignment" yaml:"alignment"`
	// The color of the rectangle behind the captions.
	// Default: - service default.
	//
	// Experimental.
	BackgroundColor CaptionBackgroundColor `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The opacity of the background rectangle (0 transparent .. 255 opaque).
	// Default: - service default.
	//
	// Experimental.
	BackgroundOpacity *float64 `field:"optional" json:"backgroundOpacity" yaml:"backgroundOpacity"`
	// An external font file (.ttf or .tte) used for burn-in. Provide a `FileLocation` referencing an S3 bucket (`FileLocation.fromBucket`, which auto-grants read access) or a URL (`FileLocation.url`).
	// Default: - service default font.
	//
	// Experimental.
	Font FileLocation `field:"optional" json:"font" yaml:"font"`
	// The color of the burned-in captions.
	// Default: CaptionFontColor.WHITE
	//
	// Experimental.
	FontColor CaptionFontColor `field:"optional" json:"fontColor" yaml:"fontColor"`
	// The opacity of the burned-in captions (0 transparent .. 255 opaque).
	// Default: 255.
	//
	// Experimental.
	FontOpacity *float64 `field:"optional" json:"fontOpacity" yaml:"fontOpacity"`
	// The font resolution in DPI.
	// Default: 96.
	//
	// Experimental.
	FontResolution *float64 `field:"optional" json:"fontResolution" yaml:"fontResolution"`
	// Font size — `CaptionFontSize.AUTO` to scale with the output, or `CaptionFontSize.of(points)` for an exact size in points.
	// Default: CaptionFontSize.AUTO
	//
	// Experimental.
	FontSize CaptionFontSize `field:"optional" json:"fontSize" yaml:"fontSize"`
	// The font outline color.
	// Default: CaptionOutlineColor.BLACK
	//
	// Experimental.
	OutlineColor CaptionOutlineColor `field:"optional" json:"outlineColor" yaml:"outlineColor"`
	// The font outline size in pixels.
	// Default: 2.
	//
	// Experimental.
	OutlineSize *float64 `field:"optional" json:"outlineSize" yaml:"outlineSize"`
	// The color of the shadow cast by the captions.
	// Default: CaptionShadowColor.NONE
	//
	// Experimental.
	ShadowColor CaptionShadowColor `field:"optional" json:"shadowColor" yaml:"shadowColor"`
	// The opacity of the shadow (0 transparent .. 255 opaque).
	// Default: 0.
	//
	// Experimental.
	ShadowOpacity *float64 `field:"optional" json:"shadowOpacity" yaml:"shadowOpacity"`
	// The horizontal offset of the shadow in pixels (negative shifts left).
	// Default: - service default.
	//
	// Experimental.
	ShadowXOffset *float64 `field:"optional" json:"shadowXOffset" yaml:"shadowXOffset"`
	// The vertical offset of the shadow in pixels (negative shifts up).
	// Default: - service default.
	//
	// Experimental.
	ShadowYOffset *float64 `field:"optional" json:"shadowYOffset" yaml:"shadowYOffset"`
	// For Teletext input, the number of lines for the captions bitmap.
	// Default: - service default.
	//
	// Experimental.
	SubtitleRows *string `field:"optional" json:"subtitleRows" yaml:"subtitleRows"`
	// Whether a fixed grid is used to generate the subtitle bitmap (Teletext input).
	// Default: CaptionTeletextGridControl.FIXED
	//
	// Experimental.
	TeletextGridControl CaptionTeletextGridControl `field:"optional" json:"teletextGridControl" yaml:"teletextGridControl"`
	// The horizontal position of the captions in pixels from the left.
	// Default: - determined by `alignment`.
	//
	// Experimental.
	XPosition *float64 `field:"optional" json:"xPosition" yaml:"xPosition"`
	// The vertical position of the captions in pixels from the top.
	// Default: - positioned towards the bottom.
	//
	// Experimental.
	YPosition *float64 `field:"optional" json:"yPosition" yaml:"yPosition"`
}

