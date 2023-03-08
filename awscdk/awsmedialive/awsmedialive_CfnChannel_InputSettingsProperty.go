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
//   inputSettingsProperty := &inputSettingsProperty{
//   	audioSelectors: []interface{}{
//   		&audioSelectorProperty{
//   			name: jsii.String("name"),
//   			selectorSettings: &audioSelectorSettingsProperty{
//   				audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   					groupId: jsii.String("groupId"),
//   					name: jsii.String("name"),
//   				},
//   				audioLanguageSelection: &audioLanguageSelectionProperty{
//   					languageCode: jsii.String("languageCode"),
//   					languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   				},
//   				audioPidSelection: &audioPidSelectionProperty{
//   					pid: jsii.Number(123),
//   				},
//   				audioTrackSelection: &audioTrackSelectionProperty{
//   					tracks: []interface{}{
//   						&audioTrackProperty{
//   							track: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	captionSelectors: []interface{}{
//   		&captionSelectorProperty{
//   			languageCode: jsii.String("languageCode"),
//   			name: jsii.String("name"),
//   			selectorSettings: &captionSelectorSettingsProperty{
//   				ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   					sourceAncillaryChannelNumber: jsii.Number(123),
//   				},
//   				aribSourceSettings: &aribSourceSettingsProperty{
//   				},
//   				dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   					ocrLanguage: jsii.String("ocrLanguage"),
//   					pid: jsii.Number(123),
//   				},
//   				embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   					convert608To708: jsii.String("convert608To708"),
//   					scte20Detection: jsii.String("scte20Detection"),
//   					source608ChannelNumber: jsii.Number(123),
//   					source608TrackNumber: jsii.Number(123),
//   				},
//   				scte20SourceSettings: &scte20SourceSettingsProperty{
//   					convert608To708: jsii.String("convert608To708"),
//   					source608ChannelNumber: jsii.Number(123),
//   				},
//   				scte27SourceSettings: &scte27SourceSettingsProperty{
//   					ocrLanguage: jsii.String("ocrLanguage"),
//   					pid: jsii.Number(123),
//   				},
//   				teletextSourceSettings: &teletextSourceSettingsProperty{
//   					outputRectangle: &captionRectangleProperty{
//   						height: jsii.Number(123),
//   						leftOffset: jsii.Number(123),
//   						topOffset: jsii.Number(123),
//   						width: jsii.Number(123),
//   					},
//   					pageNumber: jsii.String("pageNumber"),
//   				},
//   			},
//   		},
//   	},
//   	deblockFilter: jsii.String("deblockFilter"),
//   	denoiseFilter: jsii.String("denoiseFilter"),
//   	filterStrength: jsii.Number(123),
//   	inputFilter: jsii.String("inputFilter"),
//   	networkInputSettings: &networkInputSettingsProperty{
//   		hlsInputSettings: &hlsInputSettingsProperty{
//   			bandwidth: jsii.Number(123),
//   			bufferSegments: jsii.Number(123),
//   			retries: jsii.Number(123),
//   			retryInterval: jsii.Number(123),
//   			scte35Source: jsii.String("scte35Source"),
//   		},
//   		serverValidation: jsii.String("serverValidation"),
//   	},
//   	scte35Pid: jsii.Number(123),
//   	smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   	sourceEndBehavior: jsii.String("sourceEndBehavior"),
//   	videoSelector: &videoSelectorProperty{
//   		colorSpace: jsii.String("colorSpace"),
//   		colorSpaceSettings: &videoSelectorColorSpaceSettingsProperty{
//   			hdr10Settings: &hdr10SettingsProperty{
//   				maxCll: jsii.Number(123),
//   				maxFall: jsii.Number(123),
//   			},
//   		},
//   		colorSpaceUsage: jsii.String("colorSpaceUsage"),
//   		selectorSettings: &videoSelectorSettingsProperty{
//   			videoSelectorPid: &videoSelectorPidProperty{
//   				pid: jsii.Number(123),
//   			},
//   			videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   				programId: jsii.Number(123),
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

