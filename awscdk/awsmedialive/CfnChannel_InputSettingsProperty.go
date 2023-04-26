package awsmedialive


// Information about extracting content from the input and about handling the content.
//
// The parent of this entity is InputAttachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSettingsProperty := &InputSettingsProperty{
//   	AudioSelectors: []interface{}{
//   		&AudioSelectorProperty{
//   			Name: jsii.String("name"),
//   			SelectorSettings: &AudioSelectorSettingsProperty{
//   				AudioHlsRenditionSelection: &AudioHlsRenditionSelectionProperty{
//   					GroupId: jsii.String("groupId"),
//   					Name: jsii.String("name"),
//   				},
//   				AudioLanguageSelection: &AudioLanguageSelectionProperty{
//   					LanguageCode: jsii.String("languageCode"),
//   					LanguageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   				},
//   				AudioPidSelection: &AudioPidSelectionProperty{
//   					Pid: jsii.Number(123),
//   				},
//   				AudioTrackSelection: &AudioTrackSelectionProperty{
//   					Tracks: []interface{}{
//   						&AudioTrackProperty{
//   							Track: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	CaptionSelectors: []interface{}{
//   		&CaptionSelectorProperty{
//   			LanguageCode: jsii.String("languageCode"),
//   			Name: jsii.String("name"),
//   			SelectorSettings: &CaptionSelectorSettingsProperty{
//   				AncillarySourceSettings: &AncillarySourceSettingsProperty{
//   					SourceAncillaryChannelNumber: jsii.Number(123),
//   				},
//   				AribSourceSettings: &AribSourceSettingsProperty{
//   				},
//   				DvbSubSourceSettings: &DvbSubSourceSettingsProperty{
//   					OcrLanguage: jsii.String("ocrLanguage"),
//   					Pid: jsii.Number(123),
//   				},
//   				EmbeddedSourceSettings: &EmbeddedSourceSettingsProperty{
//   					Convert608To708: jsii.String("convert608To708"),
//   					Scte20Detection: jsii.String("scte20Detection"),
//   					Source608ChannelNumber: jsii.Number(123),
//   					Source608TrackNumber: jsii.Number(123),
//   				},
//   				Scte20SourceSettings: &Scte20SourceSettingsProperty{
//   					Convert608To708: jsii.String("convert608To708"),
//   					Source608ChannelNumber: jsii.Number(123),
//   				},
//   				Scte27SourceSettings: &Scte27SourceSettingsProperty{
//   					OcrLanguage: jsii.String("ocrLanguage"),
//   					Pid: jsii.Number(123),
//   				},
//   				TeletextSourceSettings: &TeletextSourceSettingsProperty{
//   					OutputRectangle: &CaptionRectangleProperty{
//   						Height: jsii.Number(123),
//   						LeftOffset: jsii.Number(123),
//   						TopOffset: jsii.Number(123),
//   						Width: jsii.Number(123),
//   					},
//   					PageNumber: jsii.String("pageNumber"),
//   				},
//   			},
//   		},
//   	},
//   	DeblockFilter: jsii.String("deblockFilter"),
//   	DenoiseFilter: jsii.String("denoiseFilter"),
//   	FilterStrength: jsii.Number(123),
//   	InputFilter: jsii.String("inputFilter"),
//   	NetworkInputSettings: &NetworkInputSettingsProperty{
//   		HlsInputSettings: &HlsInputSettingsProperty{
//   			Bandwidth: jsii.Number(123),
//   			BufferSegments: jsii.Number(123),
//   			Retries: jsii.Number(123),
//   			RetryInterval: jsii.Number(123),
//   			Scte35Source: jsii.String("scte35Source"),
//   		},
//   		ServerValidation: jsii.String("serverValidation"),
//   	},
//   	Scte35Pid: jsii.Number(123),
//   	Smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   	SourceEndBehavior: jsii.String("sourceEndBehavior"),
//   	VideoSelector: &VideoSelectorProperty{
//   		ColorSpace: jsii.String("colorSpace"),
//   		ColorSpaceSettings: &VideoSelectorColorSpaceSettingsProperty{
//   			Hdr10Settings: &Hdr10SettingsProperty{
//   				MaxCll: jsii.Number(123),
//   				MaxFall: jsii.Number(123),
//   			},
//   		},
//   		ColorSpaceUsage: jsii.String("colorSpaceUsage"),
//   		SelectorSettings: &VideoSelectorSettingsProperty{
//   			VideoSelectorPid: &VideoSelectorPidProperty{
//   				Pid: jsii.Number(123),
//   			},
//   			VideoSelectorProgramId: &VideoSelectorProgramIdProperty{
//   				ProgramId: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnChannel_InputSettingsProperty struct {
	// Information about the specific audio to extract from the input.
	//
	// The parent of this entity is InputSettings.
	AudioSelectors interface{} `field:"optional" json:"audioSelectors" yaml:"audioSelectors"`
	// Information about the specific captions to extract from the input.
	CaptionSelectors interface{} `field:"optional" json:"captionSelectors" yaml:"captionSelectors"`
	// Enables or disables the deblock filter when filtering.
	DeblockFilter *string `field:"optional" json:"deblockFilter" yaml:"deblockFilter"`
	// Enables or disables the denoise filter when filtering.
	DenoiseFilter *string `field:"optional" json:"denoiseFilter" yaml:"denoiseFilter"`
	// Adjusts the magnitude of filtering from 1 (minimal) to 5 (strongest).
	FilterStrength *float64 `field:"optional" json:"filterStrength" yaml:"filterStrength"`
	// Turns on the filter for this input.
	//
	// MPEG-2 inputs have the deblocking filter enabled by default. 1) auto - filtering is applied depending on input type/quality 2) disabled - no filtering is applied to the input 3) forced - filtering is applied regardless of the input type.
	InputFilter *string `field:"optional" json:"inputFilter" yaml:"inputFilter"`
	// Information about how to connect to the upstream system.
	NetworkInputSettings interface{} `field:"optional" json:"networkInputSettings" yaml:"networkInputSettings"`
	// `CfnChannel.InputSettingsProperty.Scte35Pid`.
	Scte35Pid *float64 `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Specifies whether to extract applicable ancillary data from a SMPTE-2038 source in this input.
	//
	// Applicable data types are captions, timecode, AFD, and SCTE-104 messages.
	// - PREFER: Extract from SMPTE-2038 if present in this input, otherwise extract from another source (if any).
	// - IGNORE: Never extract any ancillary data from SMPTE-2038.
	Smpte2038DataPreference *string `field:"optional" json:"smpte2038DataPreference" yaml:"smpte2038DataPreference"`
	// The loop input if it is a file.
	SourceEndBehavior *string `field:"optional" json:"sourceEndBehavior" yaml:"sourceEndBehavior"`
	// Information about one video to extract from the input.
	VideoSelector interface{} `field:"optional" json:"videoSelector" yaml:"videoSelector"`
}

