package previewawsmedialivemixins


// Settings for EBU-TT captions in the output.
//
// The parent of this entity is CaptionDestinationSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ebuTtDDestinationSettingsProperty := &EbuTtDDestinationSettingsProperty{
//   	CopyrightHolder: jsii.String("copyrightHolder"),
//   	DefaultFontSize: jsii.Number(123),
//   	DefaultLineHeight: jsii.Number(123),
//   	FillLineGap: jsii.String("fillLineGap"),
//   	FontFamily: jsii.String("fontFamily"),
//   	StyleControl: jsii.String("styleControl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html
//
type CfnChannelPropsMixin_EbuTtDDestinationSettingsProperty struct {
	// Applies only if you plan to convert these source captions to EBU-TT-D or TTML in an output.
	//
	// Complete this field if you want to include the name of the copyright holder in the copyright metadata tag in the TTML.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-copyrightholder
	//
	CopyrightHolder *string `field:"optional" json:"copyrightHolder" yaml:"copyrightHolder"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-defaultfontsize
	//
	DefaultFontSize *float64 `field:"optional" json:"defaultFontSize" yaml:"defaultFontSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-defaultlineheight
	//
	DefaultLineHeight *float64 `field:"optional" json:"defaultLineHeight" yaml:"defaultLineHeight"`
	// Specifies how to handle the gap between the lines (in multi-line captions).
	//
	// - enabled: Fill with the captions background color (as specified in the input captions).
	// - disabled: Leave the gap unfilled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-filllinegap
	//
	FillLineGap *string `field:"optional" json:"fillLineGap" yaml:"fillLineGap"`
	// Specifies the font family to include in the font data attached to the EBU-TT captions.
	//
	// Valid only if styleControl is set to include. If you leave this field empty, the font family is set to "monospaced". (If styleControl is set to exclude, the font family is always set to "monospaced".) You specify only the font family. All other style information (color, bold, position and so on) is copied from the input captions. The size is always set to 100% to allow the downstream player to choose the size. - Enter a list of font families, as a comma-separated list of font names, in order of preference. The name can be a font family (such as “Arial”), or a generic font family (such as “serif”), or “default” (to let the downstream player choose the font).
	// - Leave blank to set the family to “monospace”.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-fontfamily
	//
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// Specifies the style information (font color, font position, and so on) to include in the font data that is attached to the EBU-TT captions.
	//
	// - include: Take the style information (font color, font position, and so on) from the source captions and include that information in the font data attached to the EBU-TT captions. This option is valid only if the source captions are Embedded or Teletext.
	// - exclude: In the font data attached to the EBU-TT captions, set the font family to "monospaced". Do not include any other style information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-stylecontrol
	//
	StyleControl *string `field:"optional" json:"styleControl" yaml:"styleControl"`
}

