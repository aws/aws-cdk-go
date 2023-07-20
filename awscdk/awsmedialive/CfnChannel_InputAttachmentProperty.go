package awsmedialive


// An input to attach to this channel.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputAttachmentProperty := &InputAttachmentProperty{
//   	AutomaticInputFailoverSettings: &AutomaticInputFailoverSettingsProperty{
//   		ErrorClearTimeMsec: jsii.Number(123),
//   		FailoverConditions: []interface{}{
//   			&FailoverConditionProperty{
//   				FailoverConditionSettings: &FailoverConditionSettingsProperty{
//   					AudioSilenceSettings: &AudioSilenceFailoverSettingsProperty{
//   						AudioSelectorName: jsii.String("audioSelectorName"),
//   						AudioSilenceThresholdMsec: jsii.Number(123),
//   					},
//   					InputLossSettings: &InputLossFailoverSettingsProperty{
//   						InputLossThresholdMsec: jsii.Number(123),
//   					},
//   					VideoBlackSettings: &VideoBlackFailoverSettingsProperty{
//   						BlackDetectThreshold: jsii.Number(123),
//   						VideoBlackThresholdMsec: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		InputPreference: jsii.String("inputPreference"),
//   		SecondaryInputId: jsii.String("secondaryInputId"),
//   	},
//   	InputAttachmentName: jsii.String("inputAttachmentName"),
//   	InputId: jsii.String("inputId"),
//   	InputSettings: &InputSettingsProperty{
//   		AudioSelectors: []interface{}{
//   			&AudioSelectorProperty{
//   				Name: jsii.String("name"),
//   				SelectorSettings: &AudioSelectorSettingsProperty{
//   					AudioHlsRenditionSelection: &AudioHlsRenditionSelectionProperty{
//   						GroupId: jsii.String("groupId"),
//   						Name: jsii.String("name"),
//   					},
//   					AudioLanguageSelection: &AudioLanguageSelectionProperty{
//   						LanguageCode: jsii.String("languageCode"),
//   						LanguageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   					},
//   					AudioPidSelection: &AudioPidSelectionProperty{
//   						Pid: jsii.Number(123),
//   					},
//   					AudioTrackSelection: &AudioTrackSelectionProperty{
//   						DolbyEDecode: &AudioDolbyEDecodeProperty{
//   							ProgramSelection: jsii.String("programSelection"),
//   						},
//   						Tracks: []interface{}{
//   							&AudioTrackProperty{
//   								Track: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		CaptionSelectors: []interface{}{
//   			&CaptionSelectorProperty{
//   				LanguageCode: jsii.String("languageCode"),
//   				Name: jsii.String("name"),
//   				SelectorSettings: &CaptionSelectorSettingsProperty{
//   					AncillarySourceSettings: &AncillarySourceSettingsProperty{
//   						SourceAncillaryChannelNumber: jsii.Number(123),
//   					},
//   					AribSourceSettings: &AribSourceSettingsProperty{
//   					},
//   					DvbSubSourceSettings: &DvbSubSourceSettingsProperty{
//   						OcrLanguage: jsii.String("ocrLanguage"),
//   						Pid: jsii.Number(123),
//   					},
//   					EmbeddedSourceSettings: &EmbeddedSourceSettingsProperty{
//   						Convert608To708: jsii.String("convert608To708"),
//   						Scte20Detection: jsii.String("scte20Detection"),
//   						Source608ChannelNumber: jsii.Number(123),
//   						Source608TrackNumber: jsii.Number(123),
//   					},
//   					Scte20SourceSettings: &Scte20SourceSettingsProperty{
//   						Convert608To708: jsii.String("convert608To708"),
//   						Source608ChannelNumber: jsii.Number(123),
//   					},
//   					Scte27SourceSettings: &Scte27SourceSettingsProperty{
//   						OcrLanguage: jsii.String("ocrLanguage"),
//   						Pid: jsii.Number(123),
//   					},
//   					TeletextSourceSettings: &TeletextSourceSettingsProperty{
//   						OutputRectangle: &CaptionRectangleProperty{
//   							Height: jsii.Number(123),
//   							LeftOffset: jsii.Number(123),
//   							TopOffset: jsii.Number(123),
//   							Width: jsii.Number(123),
//   						},
//   						PageNumber: jsii.String("pageNumber"),
//   					},
//   				},
//   			},
//   		},
//   		DeblockFilter: jsii.String("deblockFilter"),
//   		DenoiseFilter: jsii.String("denoiseFilter"),
//   		FilterStrength: jsii.Number(123),
//   		InputFilter: jsii.String("inputFilter"),
//   		NetworkInputSettings: &NetworkInputSettingsProperty{
//   			HlsInputSettings: &HlsInputSettingsProperty{
//   				Bandwidth: jsii.Number(123),
//   				BufferSegments: jsii.Number(123),
//   				Retries: jsii.Number(123),
//   				RetryInterval: jsii.Number(123),
//   				Scte35Source: jsii.String("scte35Source"),
//   			},
//   			ServerValidation: jsii.String("serverValidation"),
//   		},
//   		Scte35Pid: jsii.Number(123),
//   		Smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   		SourceEndBehavior: jsii.String("sourceEndBehavior"),
//   		VideoSelector: &VideoSelectorProperty{
//   			ColorSpace: jsii.String("colorSpace"),
//   			ColorSpaceSettings: &VideoSelectorColorSpaceSettingsProperty{
//   				Hdr10Settings: &Hdr10SettingsProperty{
//   					MaxCll: jsii.Number(123),
//   					MaxFall: jsii.Number(123),
//   				},
//   			},
//   			ColorSpaceUsage: jsii.String("colorSpaceUsage"),
//   			SelectorSettings: &VideoSelectorSettingsProperty{
//   				VideoSelectorPid: &VideoSelectorPidProperty{
//   					Pid: jsii.Number(123),
//   				},
//   				VideoSelectorProgramId: &VideoSelectorProgramIdProperty{
//   					ProgramId: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html
//
type CfnChannel_InputAttachmentProperty struct {
	// Settings to implement automatic input failover in this input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html#cfn-medialive-channel-inputattachment-automaticinputfailoversettings
	//
	AutomaticInputFailoverSettings interface{} `field:"optional" json:"automaticInputFailoverSettings" yaml:"automaticInputFailoverSettings"`
	// A name for the attachment.
	//
	// This is required if you want to use this input in an input switch action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html#cfn-medialive-channel-inputattachment-inputattachmentname
	//
	InputAttachmentName *string `field:"optional" json:"inputAttachmentName" yaml:"inputAttachmentName"`
	// The ID of the input to attach.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html#cfn-medialive-channel-inputattachment-inputid
	//
	InputId *string `field:"optional" json:"inputId" yaml:"inputId"`
	// Information about the content to extract from the input and about the general handling of the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputattachment.html#cfn-medialive-channel-inputattachment-inputsettings
	//
	InputSettings interface{} `field:"optional" json:"inputSettings" yaml:"inputSettings"`
}

