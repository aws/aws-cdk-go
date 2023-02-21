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
//   captionSelectorProperty := &captionSelectorProperty{
//   	languageCode: jsii.String("languageCode"),
//   	name: jsii.String("name"),
//   	selectorSettings: &captionSelectorSettingsProperty{
//   		ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   			sourceAncillaryChannelNumber: jsii.Number(123),
//   		},
//   		aribSourceSettings: &aribSourceSettingsProperty{
//   		},
//   		dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   			ocrLanguage: jsii.String("ocrLanguage"),
//   			pid: jsii.Number(123),
//   		},
//   		embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   			convert608To708: jsii.String("convert608To708"),
//   			scte20Detection: jsii.String("scte20Detection"),
//   			source608ChannelNumber: jsii.Number(123),
//   			source608TrackNumber: jsii.Number(123),
//   		},
//   		scte20SourceSettings: &scte20SourceSettingsProperty{
//   			convert608To708: jsii.String("convert608To708"),
//   			source608ChannelNumber: jsii.Number(123),
//   		},
//   		scte27SourceSettings: &scte27SourceSettingsProperty{
//   			ocrLanguage: jsii.String("ocrLanguage"),
//   			pid: jsii.Number(123),
//   		},
//   		teletextSourceSettings: &teletextSourceSettingsProperty{
//   			outputRectangle: &captionRectangleProperty{
//   				height: jsii.Number(123),
//   				leftOffset: jsii.Number(123),
//   				topOffset: jsii.Number(123),
//   				width: jsii.Number(123),
//   			},
//   			pageNumber: jsii.String("pageNumber"),
//   		},
//   	},
//   }
//
type CfnChannel_CaptionSelectorProperty struct {
	// When specified, this field indicates the three-letter language code of the captions track to extract from the source.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The name identifier for a captions selector.
	//
	// This name is used to associate this captions selector with one or more captions descriptions. Names must be unique within a channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the specific audio to extract from the input.
	SelectorSettings interface{} `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

