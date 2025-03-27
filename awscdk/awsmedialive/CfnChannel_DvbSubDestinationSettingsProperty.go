package awsmedialive


// The settings for DVB Sub captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbSubDestinationSettingsProperty := &DvbSubDestinationSettingsProperty{
//   	Alignment: jsii.String("alignment"),
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	BackgroundOpacity: jsii.Number(123),
//   	Font: &InputLocationProperty{
//   		PasswordParam: jsii.String("passwordParam"),
//   		Uri: jsii.String("uri"),
//   		Username: jsii.String("username"),
//   	},
//   	FontColor: jsii.String("fontColor"),
//   	FontOpacity: jsii.Number(123),
//   	FontResolution: jsii.Number(123),
//   	FontSize: jsii.String("fontSize"),
//   	OutlineColor: jsii.String("outlineColor"),
//   	OutlineSize: jsii.Number(123),
//   	ShadowColor: jsii.String("shadowColor"),
//   	ShadowOpacity: jsii.Number(123),
//   	ShadowXOffset: jsii.Number(123),
//   	ShadowYOffset: jsii.Number(123),
//   	TeletextGridControl: jsii.String("teletextGridControl"),
//   	XPosition: jsii.Number(123),
//   	YPosition: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html
//
type CfnChannel_DvbSubDestinationSettingsProperty struct {
	// If no explicit xPosition or yPosition is provided, setting the alignment to centered places the captions at the bottom center of the output.
	//
	// Similarly, setting a left alignment aligns captions to the bottom left of the output. If x and y positions are specified in conjunction with the alignment parameter, the font is justified (either left or centered) relative to those coordinates. Selecting "smart" justification left-justifies live subtitles and center-justifies pre-recorded subtitles. This option is not valid for source captions that are STL or 608/embedded. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-alignment
	//
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Specifies the color of the rectangle behind the captions.
	//
	// All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-backgroundcolor
	//
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Specifies the opacity of the background rectangle.
	//
	// 255 is opaque; 0 is transparent. Keeping this parameter blank is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-backgroundopacity
	//
	BackgroundOpacity *float64 `field:"optional" json:"backgroundOpacity" yaml:"backgroundOpacity"`
	// The external font file that is used for captions burn-in.
	//
	// The file extension must be .ttf or .tte. Although you can select output fonts for many different types of input captions, embedded, STL, and Teletext sources use a strict grid system. Using external fonts with these captions sources could cause an unexpected display of proportional fonts. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-font
	//
	Font interface{} `field:"optional" json:"font" yaml:"font"`
	// Specifies the color of the burned-in captions.
	//
	// This option is not valid for source captions that are STL, 608/embedded, or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-fontcolor
	//
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Specifies the opacity of the burned-in captions.
	//
	// 255 is opaque; 0 is transparent. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-fontopacity
	//
	FontOpacity *float64 `field:"optional" json:"fontOpacity" yaml:"fontOpacity"`
	// The font resolution in DPI (dots per inch).
	//
	// The default is 96 dpi. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-fontresolution
	//
	FontResolution *float64 `field:"optional" json:"fontResolution" yaml:"fontResolution"`
	// When set to auto, fontSize scales depending on the size of the output.
	//
	// Providing a positive integer specifies the exact font size in points. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-fontsize
	//
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Specifies the font outline color.
	//
	// This option is not valid for source captions that are either 608/embedded or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-outlinecolor
	//
	OutlineColor *string `field:"optional" json:"outlineColor" yaml:"outlineColor"`
	// Specifies the font outline size in pixels.
	//
	// This option is not valid for source captions that are either 608/embedded or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-outlinesize
	//
	OutlineSize *float64 `field:"optional" json:"outlineSize" yaml:"outlineSize"`
	// Specifies the color of the shadow that is cast by the captions.
	//
	// All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-shadowcolor
	//
	ShadowColor *string `field:"optional" json:"shadowColor" yaml:"shadowColor"`
	// Specifies the opacity of the shadow.
	//
	// 255 is opaque; 0 is transparent. Keeping this parameter blank is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-shadowopacity
	//
	ShadowOpacity *float64 `field:"optional" json:"shadowOpacity" yaml:"shadowOpacity"`
	// Specifies the horizontal offset of the shadow relative to the captions in pixels.
	//
	// A value of -2 would result in a shadow offset 2 pixels to the left. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-shadowxoffset
	//
	ShadowXOffset *float64 `field:"optional" json:"shadowXOffset" yaml:"shadowXOffset"`
	// Specifies the vertical offset of the shadow relative to the captions in pixels.
	//
	// A value of -2 would result in a shadow offset 2 pixels above the text. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-shadowyoffset
	//
	ShadowYOffset *float64 `field:"optional" json:"shadowYOffset" yaml:"shadowYOffset"`
	// Controls whether a fixed grid size is used to generate the output subtitles bitmap.
	//
	// This applies to only Teletext inputs and DVB-Sub/Burn-in outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-teletextgridcontrol
	//
	TeletextGridControl *string `field:"optional" json:"teletextGridControl" yaml:"teletextGridControl"`
	// Specifies the horizontal position of the captions relative to the left side of the output in pixels.
	//
	// A value of 10 would result in the captions starting 10 pixels from the left of the output. If no explicit xPosition is provided, the horizontal captions position is determined by the alignment parameter. This option is not valid for source captions that are STL, 608/embedded, or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-xposition
	//
	XPosition *float64 `field:"optional" json:"xPosition" yaml:"xPosition"`
	// Specifies the vertical position of the captions relative to the top of the output in pixels.
	//
	// A value of 10 would result in the captions starting 10 pixels from the top of the output. If no explicit yPosition is provided, the captions are positioned towards the bottom of the output. This option is not valid for source captions that are STL, 608/embedded, or Teletext. These source settings are already pre-defined by the captions stream. All burn-in and DVB-Sub font settings must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubdestinationsettings.html#cfn-medialive-channel-dvbsubdestinationsettings-yposition
	//
	YPosition *float64 `field:"optional" json:"yPosition" yaml:"yPosition"`
}

