package awsmedialive


// Captions Selector Settings.
//
// The parent of this entity is CaptionSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionSelectorSettingsProperty := &captionSelectorSettingsProperty{
//   	ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   		sourceAncillaryChannelNumber: jsii.Number(123),
//   	},
//   	aribSourceSettings: &aribSourceSettingsProperty{
//   	},
//   	dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   		ocrLanguage: jsii.String("ocrLanguage"),
//   		pid: jsii.Number(123),
//   	},
//   	embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   		convert608To708: jsii.String("convert608To708"),
//   		scte20Detection: jsii.String("scte20Detection"),
//   		source608ChannelNumber: jsii.Number(123),
//   		source608TrackNumber: jsii.Number(123),
//   	},
//   	scte20SourceSettings: &scte20SourceSettingsProperty{
//   		convert608To708: jsii.String("convert608To708"),
//   		source608ChannelNumber: jsii.Number(123),
//   	},
//   	scte27SourceSettings: &scte27SourceSettingsProperty{
//   		ocrLanguage: jsii.String("ocrLanguage"),
//   		pid: jsii.Number(123),
//   	},
//   	teletextSourceSettings: &teletextSourceSettingsProperty{
//   		outputRectangle: &captionRectangleProperty{
//   			height: jsii.Number(123),
//   			leftOffset: jsii.Number(123),
//   			topOffset: jsii.Number(123),
//   			width: jsii.Number(123),
//   		},
//   		pageNumber: jsii.String("pageNumber"),
//   	},
//   }
//
type CfnChannel_CaptionSelectorSettingsProperty struct {
	// Information about the ancillary captions to extract from the input.
	AncillarySourceSettings interface{} `field:"optional" json:"ancillarySourceSettings" yaml:"ancillarySourceSettings"`
	// Information about the ARIB captions to extract from the input.
	AribSourceSettings interface{} `field:"optional" json:"aribSourceSettings" yaml:"aribSourceSettings"`
	// Information about the DVB Sub captions to extract from the input.
	DvbSubSourceSettings interface{} `field:"optional" json:"dvbSubSourceSettings" yaml:"dvbSubSourceSettings"`
	// Information about the embedded captions to extract from the input.
	EmbeddedSourceSettings interface{} `field:"optional" json:"embeddedSourceSettings" yaml:"embeddedSourceSettings"`
	// Information about the SCTE-20 captions to extract from the input.
	Scte20SourceSettings interface{} `field:"optional" json:"scte20SourceSettings" yaml:"scte20SourceSettings"`
	// Information about the SCTE-27 captions to extract from the input.
	Scte27SourceSettings interface{} `field:"optional" json:"scte27SourceSettings" yaml:"scte27SourceSettings"`
	// Information about the Teletext captions to extract from the input.
	TeletextSourceSettings interface{} `field:"optional" json:"teletextSourceSettings" yaml:"teletextSourceSettings"`
}

