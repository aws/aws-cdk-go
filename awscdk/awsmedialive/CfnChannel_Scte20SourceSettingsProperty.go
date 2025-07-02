package awsmedialive


// Information about the SCTE-20 captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte20SourceSettingsProperty := &Scte20SourceSettingsProperty{
//   	Convert608To708: jsii.String("convert608To708"),
//   	Source608ChannelNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte20sourcesettings.html
//
type CfnChannel_Scte20SourceSettingsProperty struct {
	// If upconvert, 608 data is both passed through the "608 compatibility bytes" fields of the 708 wrapper as well as translated into 708.
	//
	// Any 708 data present in the source content is discarded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte20sourcesettings.html#cfn-medialive-channel-scte20sourcesettings-convert608to708
	//
	Convert608To708 *string `field:"optional" json:"convert608To708" yaml:"convert608To708"`
	// Specifies the 608/708 channel number within the video track from which to extract captions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte20sourcesettings.html#cfn-medialive-channel-scte20sourcesettings-source608channelnumber
	//
	Source608ChannelNumber *float64 `field:"optional" json:"source608ChannelNumber" yaml:"source608ChannelNumber"`
}

