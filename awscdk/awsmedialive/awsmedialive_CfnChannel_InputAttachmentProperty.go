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
//   inputAttachmentProperty := &inputAttachmentProperty{
//   	automaticInputFailoverSettings: &automaticInputFailoverSettingsProperty{
//   		errorClearTimeMsec: jsii.Number(123),
//   		failoverConditions: []interface{}{
//   			&failoverConditionProperty{
//   				failoverConditionSettings: &failoverConditionSettingsProperty{
//   					audioSilenceSettings: &audioSilenceFailoverSettingsProperty{
//   						audioSelectorName: jsii.String("audioSelectorName"),
//   						audioSilenceThresholdMsec: jsii.Number(123),
//   					},
//   					inputLossSettings: &inputLossFailoverSettingsProperty{
//   						inputLossThresholdMsec: jsii.Number(123),
//   					},
//   					videoBlackSettings: &videoBlackFailoverSettingsProperty{
//   						blackDetectThreshold: jsii.Number(123),
//   						videoBlackThresholdMsec: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		inputPreference: jsii.String("inputPreference"),
//   		secondaryInputId: jsii.String("secondaryInputId"),
//   	},
//   	inputAttachmentName: jsii.String("inputAttachmentName"),
//   	inputId: jsii.String("inputId"),
//   	inputSettings: &inputSettingsProperty{
//   		audioSelectors: []interface{}{
//   			&audioSelectorProperty{
//   				name: jsii.String("name"),
//   				selectorSettings: &audioSelectorSettingsProperty{
//   					audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   						groupId: jsii.String("groupId"),
//   						name: jsii.String("name"),
//   					},
//   					audioLanguageSelection: &audioLanguageSelectionProperty{
//   						languageCode: jsii.String("languageCode"),
//   						languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   					},
//   					audioPidSelection: &audioPidSelectionProperty{
//   						pid: jsii.Number(123),
//   					},
//   					audioTrackSelection: &audioTrackSelectionProperty{
//   						tracks: []interface{}{
//   							&audioTrackProperty{
//   								track: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		captionSelectors: []interface{}{
//   			&captionSelectorProperty{
//   				languageCode: jsii.String("languageCode"),
//   				name: jsii.String("name"),
//   				selectorSettings: &captionSelectorSettingsProperty{
//   					ancillarySourceSettings: &ancillarySourceSettingsProperty{
//   						sourceAncillaryChannelNumber: jsii.Number(123),
//   					},
//   					aribSourceSettings: &aribSourceSettingsProperty{
//   					},
//   					dvbSubSourceSettings: &dvbSubSourceSettingsProperty{
//   						ocrLanguage: jsii.String("ocrLanguage"),
//   						pid: jsii.Number(123),
//   					},
//   					embeddedSourceSettings: &embeddedSourceSettingsProperty{
//   						convert608To708: jsii.String("convert608To708"),
//   						scte20Detection: jsii.String("scte20Detection"),
//   						source608ChannelNumber: jsii.Number(123),
//   						source608TrackNumber: jsii.Number(123),
//   					},
//   					scte20SourceSettings: &scte20SourceSettingsProperty{
//   						convert608To708: jsii.String("convert608To708"),
//   						source608ChannelNumber: jsii.Number(123),
//   					},
//   					scte27SourceSettings: &scte27SourceSettingsProperty{
//   						ocrLanguage: jsii.String("ocrLanguage"),
//   						pid: jsii.Number(123),
//   					},
//   					teletextSourceSettings: &teletextSourceSettingsProperty{
//   						outputRectangle: &captionRectangleProperty{
//   							height: jsii.Number(123),
//   							leftOffset: jsii.Number(123),
//   							topOffset: jsii.Number(123),
//   							width: jsii.Number(123),
//   						},
//   						pageNumber: jsii.String("pageNumber"),
//   					},
//   				},
//   			},
//   		},
//   		deblockFilter: jsii.String("deblockFilter"),
//   		denoiseFilter: jsii.String("denoiseFilter"),
//   		filterStrength: jsii.Number(123),
//   		inputFilter: jsii.String("inputFilter"),
//   		networkInputSettings: &networkInputSettingsProperty{
//   			hlsInputSettings: &hlsInputSettingsProperty{
//   				bandwidth: jsii.Number(123),
//   				bufferSegments: jsii.Number(123),
//   				retries: jsii.Number(123),
//   				retryInterval: jsii.Number(123),
//   				scte35Source: jsii.String("scte35Source"),
//   			},
//   			serverValidation: jsii.String("serverValidation"),
//   		},
//   		scte35Pid: jsii.Number(123),
//   		smpte2038DataPreference: jsii.String("smpte2038DataPreference"),
//   		sourceEndBehavior: jsii.String("sourceEndBehavior"),
//   		videoSelector: &videoSelectorProperty{
//   			colorSpace: jsii.String("colorSpace"),
//   			colorSpaceSettings: &videoSelectorColorSpaceSettingsProperty{
//   				hdr10Settings: &hdr10SettingsProperty{
//   					maxCll: jsii.Number(123),
//   					maxFall: jsii.Number(123),
//   				},
//   			},
//   			colorSpaceUsage: jsii.String("colorSpaceUsage"),
//   			selectorSettings: &videoSelectorSettingsProperty{
//   				videoSelectorPid: &videoSelectorPidProperty{
//   					pid: jsii.Number(123),
//   				},
//   				videoSelectorProgramId: &videoSelectorProgramIdProperty{
//   					programId: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnChannel_InputAttachmentProperty struct {
	// Settings to implement automatic input failover in this input.
	AutomaticInputFailoverSettings interface{} `field:"optional" json:"automaticInputFailoverSettings" yaml:"automaticInputFailoverSettings"`
	// A name for the attachment.
	//
	// This is required if you want to use this input in an input switch action.
	InputAttachmentName *string `field:"optional" json:"inputAttachmentName" yaml:"inputAttachmentName"`
	// The ID of the input to attach.
	InputId *string `field:"optional" json:"inputId" yaml:"inputId"`
	// Information about the content to extract from the input and about the general handling of the content.
	InputSettings interface{} `field:"optional" json:"inputSettings" yaml:"inputSettings"`
}

