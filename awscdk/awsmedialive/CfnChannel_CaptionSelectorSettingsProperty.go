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
//   captionSelectorSettingsProperty := &CaptionSelectorSettingsProperty{
//   	AncillarySourceSettings: &AncillarySourceSettingsProperty{
//   		SourceAncillaryChannelNumber: jsii.Number(123),
//   	},
//   	AribSourceSettings: &AribSourceSettingsProperty{
//   	},
//   	DvbSubSourceSettings: &DvbSubSourceSettingsProperty{
//   		OcrLanguage: jsii.String("ocrLanguage"),
//   		Pid: jsii.Number(123),
//   	},
//   	EmbeddedSourceSettings: &EmbeddedSourceSettingsProperty{
//   		Convert608To708: jsii.String("convert608To708"),
//   		Scte20Detection: jsii.String("scte20Detection"),
//   		Source608ChannelNumber: jsii.Number(123),
//   		Source608TrackNumber: jsii.Number(123),
//   	},
//   	Scte20SourceSettings: &Scte20SourceSettingsProperty{
//   		Convert608To708: jsii.String("convert608To708"),
//   		Source608ChannelNumber: jsii.Number(123),
//   	},
//   	Scte27SourceSettings: &Scte27SourceSettingsProperty{
//   		OcrLanguage: jsii.String("ocrLanguage"),
//   		Pid: jsii.Number(123),
//   	},
//   	TeletextSourceSettings: &TeletextSourceSettingsProperty{
//   		OutputRectangle: &CaptionRectangleProperty{
//   			Height: jsii.Number(123),
//   			LeftOffset: jsii.Number(123),
//   			TopOffset: jsii.Number(123),
//   			Width: jsii.Number(123),
//   		},
//   		PageNumber: jsii.String("pageNumber"),
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

