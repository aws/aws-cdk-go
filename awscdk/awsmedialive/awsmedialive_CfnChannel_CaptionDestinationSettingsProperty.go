package awsmedialive


// The configuration of one captions encode in the output.
//
// The parent of this entity is CaptionDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionDestinationSettingsProperty := &captionDestinationSettingsProperty{
//   	aribDestinationSettings: &aribDestinationSettingsProperty{
//   	},
//   	burnInDestinationSettings: &burnInDestinationSettingsProperty{
//   		alignment: jsii.String("alignment"),
//   		backgroundColor: jsii.String("backgroundColor"),
//   		backgroundOpacity: jsii.Number(123),
//   		font: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		fontColor: jsii.String("fontColor"),
//   		fontOpacity: jsii.Number(123),
//   		fontResolution: jsii.Number(123),
//   		fontSize: jsii.String("fontSize"),
//   		outlineColor: jsii.String("outlineColor"),
//   		outlineSize: jsii.Number(123),
//   		shadowColor: jsii.String("shadowColor"),
//   		shadowOpacity: jsii.Number(123),
//   		shadowXOffset: jsii.Number(123),
//   		shadowYOffset: jsii.Number(123),
//   		teletextGridControl: jsii.String("teletextGridControl"),
//   		xPosition: jsii.Number(123),
//   		yPosition: jsii.Number(123),
//   	},
//   	dvbSubDestinationSettings: &dvbSubDestinationSettingsProperty{
//   		alignment: jsii.String("alignment"),
//   		backgroundColor: jsii.String("backgroundColor"),
//   		backgroundOpacity: jsii.Number(123),
//   		font: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		fontColor: jsii.String("fontColor"),
//   		fontOpacity: jsii.Number(123),
//   		fontResolution: jsii.Number(123),
//   		fontSize: jsii.String("fontSize"),
//   		outlineColor: jsii.String("outlineColor"),
//   		outlineSize: jsii.Number(123),
//   		shadowColor: jsii.String("shadowColor"),
//   		shadowOpacity: jsii.Number(123),
//   		shadowXOffset: jsii.Number(123),
//   		shadowYOffset: jsii.Number(123),
//   		teletextGridControl: jsii.String("teletextGridControl"),
//   		xPosition: jsii.Number(123),
//   		yPosition: jsii.Number(123),
//   	},
//   	ebuTtDDestinationSettings: &ebuTtDDestinationSettingsProperty{
//   		copyrightHolder: jsii.String("copyrightHolder"),
//   		fillLineGap: jsii.String("fillLineGap"),
//   		fontFamily: jsii.String("fontFamily"),
//   		styleControl: jsii.String("styleControl"),
//   	},
//   	embeddedDestinationSettings: &embeddedDestinationSettingsProperty{
//   	},
//   	embeddedPlusScte20DestinationSettings: &embeddedPlusScte20DestinationSettingsProperty{
//   	},
//   	rtmpCaptionInfoDestinationSettings: &rtmpCaptionInfoDestinationSettingsProperty{
//   	},
//   	scte20PlusEmbeddedDestinationSettings: &scte20PlusEmbeddedDestinationSettingsProperty{
//   	},
//   	scte27DestinationSettings: &scte27DestinationSettingsProperty{
//   	},
//   	smpteTtDestinationSettings: &smpteTtDestinationSettingsProperty{
//   	},
//   	teletextDestinationSettings: &teletextDestinationSettingsProperty{
//   	},
//   	ttmlDestinationSettings: &ttmlDestinationSettingsProperty{
//   		styleControl: jsii.String("styleControl"),
//   	},
//   	webvttDestinationSettings: &webvttDestinationSettingsProperty{
//   		styleControl: jsii.String("styleControl"),
//   	},
//   }
//
type CfnChannel_CaptionDestinationSettingsProperty struct {
	// The configuration of one ARIB captions encode in the output.
	AribDestinationSettings interface{} `field:"optional" json:"aribDestinationSettings" yaml:"aribDestinationSettings"`
	// The configuration of one burn-in captions encode in the output.
	BurnInDestinationSettings interface{} `field:"optional" json:"burnInDestinationSettings" yaml:"burnInDestinationSettings"`
	// The configuration of one DVB Sub captions encode in the output.
	DvbSubDestinationSettings interface{} `field:"optional" json:"dvbSubDestinationSettings" yaml:"dvbSubDestinationSettings"`
	// Settings for EBU-TT captions in the output.
	EbuTtDDestinationSettings interface{} `field:"optional" json:"ebuTtDDestinationSettings" yaml:"ebuTtDDestinationSettings"`
	// The configuration of one embedded captions encode in the output.
	EmbeddedDestinationSettings interface{} `field:"optional" json:"embeddedDestinationSettings" yaml:"embeddedDestinationSettings"`
	// The configuration of one embedded plus SCTE-20 captions encode in the output.
	EmbeddedPlusScte20DestinationSettings interface{} `field:"optional" json:"embeddedPlusScte20DestinationSettings" yaml:"embeddedPlusScte20DestinationSettings"`
	// The configuration of one RTMPCaptionInfo captions encode in the output.
	RtmpCaptionInfoDestinationSettings interface{} `field:"optional" json:"rtmpCaptionInfoDestinationSettings" yaml:"rtmpCaptionInfoDestinationSettings"`
	// The configuration of one SCTE20 plus embedded captions encode in the output.
	Scte20PlusEmbeddedDestinationSettings interface{} `field:"optional" json:"scte20PlusEmbeddedDestinationSettings" yaml:"scte20PlusEmbeddedDestinationSettings"`
	// The configuration of one SCTE-27 captions encode in the output.
	Scte27DestinationSettings interface{} `field:"optional" json:"scte27DestinationSettings" yaml:"scte27DestinationSettings"`
	// The configuration of one SMPTE-TT captions encode in the output.
	SmpteTtDestinationSettings interface{} `field:"optional" json:"smpteTtDestinationSettings" yaml:"smpteTtDestinationSettings"`
	// The configuration of one Teletext captions encode in the output.
	TeletextDestinationSettings interface{} `field:"optional" json:"teletextDestinationSettings" yaml:"teletextDestinationSettings"`
	// The configuration of one TTML captions encode in the output.
	TtmlDestinationSettings interface{} `field:"optional" json:"ttmlDestinationSettings" yaml:"ttmlDestinationSettings"`
	// The configuration of one WebVTT captions encode in the output.
	WebvttDestinationSettings interface{} `field:"optional" json:"webvttDestinationSettings" yaml:"webvttDestinationSettings"`
}

