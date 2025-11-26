package previewawsmedialivemixins


// The configuration of one captions encode in the output.
//
// The parent of this entity is CaptionDescription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   captionDestinationSettingsProperty := &CaptionDestinationSettingsProperty{
//   	AribDestinationSettings: &AribDestinationSettingsProperty{
//   	},
//   	BurnInDestinationSettings: &BurnInDestinationSettingsProperty{
//   		Alignment: jsii.String("alignment"),
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		BackgroundOpacity: jsii.Number(123),
//   		Font: &InputLocationProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			Uri: jsii.String("uri"),
//   			Username: jsii.String("username"),
//   		},
//   		FontColor: jsii.String("fontColor"),
//   		FontOpacity: jsii.Number(123),
//   		FontResolution: jsii.Number(123),
//   		FontSize: jsii.String("fontSize"),
//   		OutlineColor: jsii.String("outlineColor"),
//   		OutlineSize: jsii.Number(123),
//   		ShadowColor: jsii.String("shadowColor"),
//   		ShadowOpacity: jsii.Number(123),
//   		ShadowXOffset: jsii.Number(123),
//   		ShadowYOffset: jsii.Number(123),
//   		SubtitleRows: jsii.String("subtitleRows"),
//   		TeletextGridControl: jsii.String("teletextGridControl"),
//   		XPosition: jsii.Number(123),
//   		YPosition: jsii.Number(123),
//   	},
//   	DvbSubDestinationSettings: &DvbSubDestinationSettingsProperty{
//   		Alignment: jsii.String("alignment"),
//   		BackgroundColor: jsii.String("backgroundColor"),
//   		BackgroundOpacity: jsii.Number(123),
//   		Font: &InputLocationProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			Uri: jsii.String("uri"),
//   			Username: jsii.String("username"),
//   		},
//   		FontColor: jsii.String("fontColor"),
//   		FontOpacity: jsii.Number(123),
//   		FontResolution: jsii.Number(123),
//   		FontSize: jsii.String("fontSize"),
//   		OutlineColor: jsii.String("outlineColor"),
//   		OutlineSize: jsii.Number(123),
//   		ShadowColor: jsii.String("shadowColor"),
//   		ShadowOpacity: jsii.Number(123),
//   		ShadowXOffset: jsii.Number(123),
//   		ShadowYOffset: jsii.Number(123),
//   		SubtitleRows: jsii.String("subtitleRows"),
//   		TeletextGridControl: jsii.String("teletextGridControl"),
//   		XPosition: jsii.Number(123),
//   		YPosition: jsii.Number(123),
//   	},
//   	EbuTtDDestinationSettings: &EbuTtDDestinationSettingsProperty{
//   		CopyrightHolder: jsii.String("copyrightHolder"),
//   		DefaultFontSize: jsii.Number(123),
//   		DefaultLineHeight: jsii.Number(123),
//   		FillLineGap: jsii.String("fillLineGap"),
//   		FontFamily: jsii.String("fontFamily"),
//   		StyleControl: jsii.String("styleControl"),
//   	},
//   	EmbeddedDestinationSettings: &EmbeddedDestinationSettingsProperty{
//   	},
//   	EmbeddedPlusScte20DestinationSettings: &EmbeddedPlusScte20DestinationSettingsProperty{
//   	},
//   	RtmpCaptionInfoDestinationSettings: &RtmpCaptionInfoDestinationSettingsProperty{
//   	},
//   	Scte20PlusEmbeddedDestinationSettings: &Scte20PlusEmbeddedDestinationSettingsProperty{
//   	},
//   	Scte27DestinationSettings: &Scte27DestinationSettingsProperty{
//   	},
//   	SmpteTtDestinationSettings: &SmpteTtDestinationSettingsProperty{
//   	},
//   	TeletextDestinationSettings: &TeletextDestinationSettingsProperty{
//   	},
//   	TtmlDestinationSettings: &TtmlDestinationSettingsProperty{
//   		StyleControl: jsii.String("styleControl"),
//   	},
//   	WebvttDestinationSettings: &WebvttDestinationSettingsProperty{
//   		StyleControl: jsii.String("styleControl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html
//
type CfnChannelPropsMixin_CaptionDestinationSettingsProperty struct {
	// The configuration of one ARIB captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-aribdestinationsettings
	//
	AribDestinationSettings interface{} `field:"optional" json:"aribDestinationSettings" yaml:"aribDestinationSettings"`
	// The configuration of one burn-in captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-burnindestinationsettings
	//
	BurnInDestinationSettings interface{} `field:"optional" json:"burnInDestinationSettings" yaml:"burnInDestinationSettings"`
	// The configuration of one DVB Sub captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-dvbsubdestinationsettings
	//
	DvbSubDestinationSettings interface{} `field:"optional" json:"dvbSubDestinationSettings" yaml:"dvbSubDestinationSettings"`
	// Settings for EBU-TT captions in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-ebuttddestinationsettings
	//
	EbuTtDDestinationSettings interface{} `field:"optional" json:"ebuTtDDestinationSettings" yaml:"ebuTtDDestinationSettings"`
	// The configuration of one embedded captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-embeddeddestinationsettings
	//
	EmbeddedDestinationSettings interface{} `field:"optional" json:"embeddedDestinationSettings" yaml:"embeddedDestinationSettings"`
	// The configuration of one embedded plus SCTE-20 captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-embeddedplusscte20destinationsettings
	//
	EmbeddedPlusScte20DestinationSettings interface{} `field:"optional" json:"embeddedPlusScte20DestinationSettings" yaml:"embeddedPlusScte20DestinationSettings"`
	// The configuration of one RTMPCaptionInfo captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-rtmpcaptioninfodestinationsettings
	//
	RtmpCaptionInfoDestinationSettings interface{} `field:"optional" json:"rtmpCaptionInfoDestinationSettings" yaml:"rtmpCaptionInfoDestinationSettings"`
	// The configuration of one SCTE20 plus embedded captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-scte20plusembeddeddestinationsettings
	//
	Scte20PlusEmbeddedDestinationSettings interface{} `field:"optional" json:"scte20PlusEmbeddedDestinationSettings" yaml:"scte20PlusEmbeddedDestinationSettings"`
	// The configuration of one SCTE-27 captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-scte27destinationsettings
	//
	Scte27DestinationSettings interface{} `field:"optional" json:"scte27DestinationSettings" yaml:"scte27DestinationSettings"`
	// The configuration of one SMPTE-TT captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-smptettdestinationsettings
	//
	SmpteTtDestinationSettings interface{} `field:"optional" json:"smpteTtDestinationSettings" yaml:"smpteTtDestinationSettings"`
	// The configuration of one Teletext captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-teletextdestinationsettings
	//
	TeletextDestinationSettings interface{} `field:"optional" json:"teletextDestinationSettings" yaml:"teletextDestinationSettings"`
	// The configuration of one TTML captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-ttmldestinationsettings
	//
	TtmlDestinationSettings interface{} `field:"optional" json:"ttmlDestinationSettings" yaml:"ttmlDestinationSettings"`
	// The configuration of one WebVTT captions encode in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captiondestinationsettings.html#cfn-medialive-channel-captiondestinationsettings-webvttdestinationsettings
	//
	WebvttDestinationSettings interface{} `field:"optional" json:"webvttDestinationSettings" yaml:"webvttDestinationSettings"`
}

