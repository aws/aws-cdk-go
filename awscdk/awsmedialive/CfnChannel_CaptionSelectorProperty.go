package awsmedialive


// Information about one caption to extract from the input.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionSelectorProperty := &CaptionSelectorProperty{
//   	LanguageCode: jsii.String("languageCode"),
//   	Name: jsii.String("name"),
//   	SelectorSettings: &CaptionSelectorSettingsProperty{
//   		AncillarySourceSettings: &AncillarySourceSettingsProperty{
//   			SourceAncillaryChannelNumber: jsii.Number(123),
//   		},
//   		AribSourceSettings: &AribSourceSettingsProperty{
//   		},
//   		DvbSubSourceSettings: &DvbSubSourceSettingsProperty{
//   			OcrLanguage: jsii.String("ocrLanguage"),
//   			Pid: jsii.Number(123),
//   		},
//   		EmbeddedSourceSettings: &EmbeddedSourceSettingsProperty{
//   			Convert608To708: jsii.String("convert608To708"),
//   			Scte20Detection: jsii.String("scte20Detection"),
//   			Source608ChannelNumber: jsii.Number(123),
//   			Source608TrackNumber: jsii.Number(123),
//   		},
//   		Scte20SourceSettings: &Scte20SourceSettingsProperty{
//   			Convert608To708: jsii.String("convert608To708"),
//   			Source608ChannelNumber: jsii.Number(123),
//   		},
//   		Scte27SourceSettings: &Scte27SourceSettingsProperty{
//   			OcrLanguage: jsii.String("ocrLanguage"),
//   			Pid: jsii.Number(123),
//   		},
//   		TeletextSourceSettings: &TeletextSourceSettingsProperty{
//   			OutputRectangle: &CaptionRectangleProperty{
//   				Height: jsii.Number(123),
//   				LeftOffset: jsii.Number(123),
//   				TopOffset: jsii.Number(123),
//   				Width: jsii.Number(123),
//   			},
//   			PageNumber: jsii.String("pageNumber"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselector.html
//
type CfnChannel_CaptionSelectorProperty struct {
	// When specified, this field indicates the three-letter language code of the captions track to extract from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselector.html#cfn-medialive-channel-captionselector-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The name identifier for a captions selector.
	//
	// This name is used to associate this captions selector with one or more captions descriptions. Names must be unique within a channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselector.html#cfn-medialive-channel-captionselector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the specific audio to extract from the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselector.html#cfn-medialive-channel-captionselector-selectorsettings
	//
	SelectorSettings interface{} `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

