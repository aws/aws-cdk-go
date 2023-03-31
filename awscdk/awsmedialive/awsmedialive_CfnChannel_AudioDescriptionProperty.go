package awsmedialive


// The encoding information for one output audio.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioDescriptionProperty := &audioDescriptionProperty{
//   	audioNormalizationSettings: &audioNormalizationSettingsProperty{
//   		algorithm: jsii.String("algorithm"),
//   		algorithmControl: jsii.String("algorithmControl"),
//   		targetLkfs: jsii.Number(123),
//   	},
//   	audioSelectorName: jsii.String("audioSelectorName"),
//   	audioType: jsii.String("audioType"),
//   	audioTypeControl: jsii.String("audioTypeControl"),
//   	audioWatermarkingSettings: &audioWatermarkSettingsProperty{
//   		nielsenWatermarksSettings: &nielsenWatermarksSettingsProperty{
//   			nielsenCbetSettings: &nielsenCBETProperty{
//   				cbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   				cbetStepaside: jsii.String("cbetStepaside"),
//   				csid: jsii.String("csid"),
//   			},
//   			nielsenDistributionType: jsii.String("nielsenDistributionType"),
//   			nielsenNaesIiNwSettings: &nielsenNaesIiNwProperty{
//   				checkDigitString: jsii.String("checkDigitString"),
//   				sid: jsii.Number(123),
//   			},
//   		},
//   	},
//   	codecSettings: &audioCodecSettingsProperty{
//   		aacSettings: &aacSettingsProperty{
//   			bitrate: jsii.Number(123),
//   			codingMode: jsii.String("codingMode"),
//   			inputType: jsii.String("inputType"),
//   			profile: jsii.String("profile"),
//   			rateControlMode: jsii.String("rateControlMode"),
//   			rawFormat: jsii.String("rawFormat"),
//   			sampleRate: jsii.Number(123),
//   			spec: jsii.String("spec"),
//   			vbrQuality: jsii.String("vbrQuality"),
//   		},
//   		ac3Settings: &ac3SettingsProperty{
//   			bitrate: jsii.Number(123),
//   			bitstreamMode: jsii.String("bitstreamMode"),
//   			codingMode: jsii.String("codingMode"),
//   			dialnorm: jsii.Number(123),
//   			drcProfile: jsii.String("drcProfile"),
//   			lfeFilter: jsii.String("lfeFilter"),
//   			metadataControl: jsii.String("metadataControl"),
//   		},
//   		eac3Settings: &eac3SettingsProperty{
//   			attenuationControl: jsii.String("attenuationControl"),
//   			bitrate: jsii.Number(123),
//   			bitstreamMode: jsii.String("bitstreamMode"),
//   			codingMode: jsii.String("codingMode"),
//   			dcFilter: jsii.String("dcFilter"),
//   			dialnorm: jsii.Number(123),
//   			drcLine: jsii.String("drcLine"),
//   			drcRf: jsii.String("drcRf"),
//   			lfeControl: jsii.String("lfeControl"),
//   			lfeFilter: jsii.String("lfeFilter"),
//   			loRoCenterMixLevel: jsii.Number(123),
//   			loRoSurroundMixLevel: jsii.Number(123),
//   			ltRtCenterMixLevel: jsii.Number(123),
//   			ltRtSurroundMixLevel: jsii.Number(123),
//   			metadataControl: jsii.String("metadataControl"),
//   			passthroughControl: jsii.String("passthroughControl"),
//   			phaseControl: jsii.String("phaseControl"),
//   			stereoDownmix: jsii.String("stereoDownmix"),
//   			surroundExMode: jsii.String("surroundExMode"),
//   			surroundMode: jsii.String("surroundMode"),
//   		},
//   		mp2Settings: &mp2SettingsProperty{
//   			bitrate: jsii.Number(123),
//   			codingMode: jsii.String("codingMode"),
//   			sampleRate: jsii.Number(123),
//   		},
//   		passThroughSettings: &passThroughSettingsProperty{
//   		},
//   		wavSettings: &wavSettingsProperty{
//   			bitDepth: jsii.Number(123),
//   			codingMode: jsii.String("codingMode"),
//   			sampleRate: jsii.Number(123),
//   		},
//   	},
//   	languageCode: jsii.String("languageCode"),
//   	languageCodeControl: jsii.String("languageCodeControl"),
//   	name: jsii.String("name"),
//   	remixSettings: &remixSettingsProperty{
//   		channelMappings: []interface{}{
//   			&audioChannelMappingProperty{
//   				inputChannelLevels: []interface{}{
//   					&inputChannelLevelProperty{
//   						gain: jsii.Number(123),
//   						inputChannel: jsii.Number(123),
//   					},
//   				},
//   				outputChannel: jsii.Number(123),
//   			},
//   		},
//   		channelsIn: jsii.Number(123),
//   		channelsOut: jsii.Number(123),
//   	},
//   	streamName: jsii.String("streamName"),
//   }
//
type CfnChannel_AudioDescriptionProperty struct {
	// The advanced audio normalization settings.
	AudioNormalizationSettings interface{} `field:"optional" json:"audioNormalizationSettings" yaml:"audioNormalizationSettings"`
	// The name of the AudioSelector that is used as the source for this AudioDescription.
	AudioSelectorName *string `field:"optional" json:"audioSelectorName" yaml:"audioSelectorName"`
	// Applies only if audioTypeControl is useConfigured.
	//
	// The values for audioType are defined in ISO-IEC 13818-1.
	AudioType *string `field:"optional" json:"audioType" yaml:"audioType"`
	// Determines how audio type is determined.
	//
	// followInput: If the input contains an ISO 639 audioType, then that value is passed through to the output. If the input contains no ISO 639 audioType, the value in Audio Type is included in the output. useConfigured: The value in Audio Type is included in the output. Note that this field and audioType are both ignored if inputType is broadcasterMixedAd.
	AudioTypeControl *string `field:"optional" json:"audioTypeControl" yaml:"audioTypeControl"`
	// Settings to configure one or more solutions that insert audio watermarks in the audio encode.
	AudioWatermarkingSettings interface{} `field:"optional" json:"audioWatermarkingSettings" yaml:"audioWatermarkingSettings"`
	// The audio codec settings.
	CodecSettings interface{} `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Indicates the language of the audio output track.
	//
	// Used only if languageControlMode is useConfigured, or there is no ISO 639 language code specified in the input.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Choosing followInput causes the ISO 639 language code of the output to follow the ISO 639 language code of the input.
	//
	// The languageCode setting is used when useConfigured is set, or when followInput is selected but there is no ISO 639 language code specified by the input.
	LanguageCodeControl *string `field:"optional" json:"languageCodeControl" yaml:"languageCodeControl"`
	// The name of this AudioDescription.
	//
	// Outputs use this name to uniquely identify this AudioDescription. Description names should be unique within this channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The settings that control how input audio channels are remixed into the output audio channels.
	RemixSettings interface{} `field:"optional" json:"remixSettings" yaml:"remixSettings"`
	// Used for Microsoft Smooth and Apple HLS outputs.
	//
	// Indicates the name displayed by the player (for example, English or Director Commentary).
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

