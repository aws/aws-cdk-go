package mixinsawsmedialive


// Information about the embedded captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   embeddedSourceSettingsProperty := &EmbeddedSourceSettingsProperty{
//   	Convert608To708: jsii.String("convert608To708"),
//   	Scte20Detection: jsii.String("scte20Detection"),
//   	Source608ChannelNumber: jsii.Number(123),
//   	Source608TrackNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html
//
type CfnChannelPropsMixin_EmbeddedSourceSettingsProperty struct {
	// If this is upconvert, 608 data is both passed through the "608 compatibility bytes" fields of the 708 wrapper as well as translated into 708.
	//
	// If 708 data is present in the source content, it is discarded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html#cfn-medialive-channel-embeddedsourcesettings-convert608to708
	//
	Convert608To708 *string `field:"optional" json:"convert608To708" yaml:"convert608To708"`
	// Set to "auto" to handle streams with intermittent or non-aligned SCTE-20 and embedded captions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html#cfn-medialive-channel-embeddedsourcesettings-scte20detection
	//
	Scte20Detection *string `field:"optional" json:"scte20Detection" yaml:"scte20Detection"`
	// Specifies the 608/708 channel number within the video track from which to extract captions.
	//
	// This is unused for passthrough.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html#cfn-medialive-channel-embeddedsourcesettings-source608channelnumber
	//
	Source608ChannelNumber *float64 `field:"optional" json:"source608ChannelNumber" yaml:"source608ChannelNumber"`
	// This field is unused and deprecated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html#cfn-medialive-channel-embeddedsourcesettings-source608tracknumber
	//
	Source608TrackNumber *float64 `field:"optional" json:"source608TrackNumber" yaml:"source608TrackNumber"`
}

