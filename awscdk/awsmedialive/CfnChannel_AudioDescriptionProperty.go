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
//   audioDescriptionProperty := &AudioDescriptionProperty{
//   	AudioNormalizationSettings: &AudioNormalizationSettingsProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		AlgorithmControl: jsii.String("algorithmControl"),
//   		TargetLkfs: jsii.Number(123),
//   	},
//   	AudioSelectorName: jsii.String("audioSelectorName"),
//   	AudioType: jsii.String("audioType"),
//   	AudioTypeControl: jsii.String("audioTypeControl"),
//   	AudioWatermarkingSettings: &AudioWatermarkSettingsProperty{
//   		NielsenWatermarksSettings: &NielsenWatermarksSettingsProperty{
//   			NielsenCbetSettings: &NielsenCBETProperty{
//   				CbetCheckDigitString: jsii.String("cbetCheckDigitString"),
//   				CbetStepaside: jsii.String("cbetStepaside"),
//   				Csid: jsii.String("csid"),
//   			},
//   			NielsenDistributionType: jsii.String("nielsenDistributionType"),
//   			NielsenNaesIiNwSettings: &NielsenNaesIiNwProperty{
//   				CheckDigitString: jsii.String("checkDigitString"),
//   				Sid: jsii.Number(123),
//   				Timezone: jsii.String("timezone"),
//   			},
//   		},
//   	},
//   	CodecSettings: &AudioCodecSettingsProperty{
//   		AacSettings: &AacSettingsProperty{
//   			Bitrate: jsii.Number(123),
//   			CodingMode: jsii.String("codingMode"),
//   			InputType: jsii.String("inputType"),
//   			Profile: jsii.String("profile"),
//   			RateControlMode: jsii.String("rateControlMode"),
//   			RawFormat: jsii.String("rawFormat"),
//   			SampleRate: jsii.Number(123),
//   			Spec: jsii.String("spec"),
//   			VbrQuality: jsii.String("vbrQuality"),
//   		},
//   		Ac3Settings: &Ac3SettingsProperty{
//   			AttenuationControl: jsii.String("attenuationControl"),
//   			Bitrate: jsii.Number(123),
//   			BitstreamMode: jsii.String("bitstreamMode"),
//   			CodingMode: jsii.String("codingMode"),
//   			Dialnorm: jsii.Number(123),
//   			DrcProfile: jsii.String("drcProfile"),
//   			LfeFilter: jsii.String("lfeFilter"),
//   			MetadataControl: jsii.String("metadataControl"),
//   		},
//   		Eac3AtmosSettings: &Eac3AtmosSettingsProperty{
//   			Bitrate: jsii.Number(123),
//   			CodingMode: jsii.String("codingMode"),
//   			Dialnorm: jsii.Number(123),
//   			DrcLine: jsii.String("drcLine"),
//   			DrcRf: jsii.String("drcRf"),
//   			HeightTrim: jsii.Number(123),
//   			SurroundTrim: jsii.Number(123),
//   		},
//   		Eac3Settings: &Eac3SettingsProperty{
//   			AttenuationControl: jsii.String("attenuationControl"),
//   			Bitrate: jsii.Number(123),
//   			BitstreamMode: jsii.String("bitstreamMode"),
//   			CodingMode: jsii.String("codingMode"),
//   			DcFilter: jsii.String("dcFilter"),
//   			Dialnorm: jsii.Number(123),
//   			DrcLine: jsii.String("drcLine"),
//   			DrcRf: jsii.String("drcRf"),
//   			LfeControl: jsii.String("lfeControl"),
//   			LfeFilter: jsii.String("lfeFilter"),
//   			LoRoCenterMixLevel: jsii.Number(123),
//   			LoRoSurroundMixLevel: jsii.Number(123),
//   			LtRtCenterMixLevel: jsii.Number(123),
//   			LtRtSurroundMixLevel: jsii.Number(123),
//   			MetadataControl: jsii.String("metadataControl"),
//   			PassthroughControl: jsii.String("passthroughControl"),
//   			PhaseControl: jsii.String("phaseControl"),
//   			StereoDownmix: jsii.String("stereoDownmix"),
//   			SurroundExMode: jsii.String("surroundExMode"),
//   			SurroundMode: jsii.String("surroundMode"),
//   		},
//   		Mp2Settings: &Mp2SettingsProperty{
//   			Bitrate: jsii.Number(123),
//   			CodingMode: jsii.String("codingMode"),
//   			SampleRate: jsii.Number(123),
//   		},
//   		PassThroughSettings: &PassThroughSettingsProperty{
//   		},
//   		WavSettings: &WavSettingsProperty{
//   			BitDepth: jsii.Number(123),
//   			CodingMode: jsii.String("codingMode"),
//   			SampleRate: jsii.Number(123),
//   		},
//   	},
//   	LanguageCode: jsii.String("languageCode"),
//   	LanguageCodeControl: jsii.String("languageCodeControl"),
//   	Name: jsii.String("name"),
//   	RemixSettings: &RemixSettingsProperty{
//   		ChannelMappings: []interface{}{
//   			&AudioChannelMappingProperty{
//   				InputChannelLevels: []interface{}{
//   					&InputChannelLevelProperty{
//   						Gain: jsii.Number(123),
//   						InputChannel: jsii.Number(123),
//   					},
//   				},
//   				OutputChannel: jsii.Number(123),
//   			},
//   		},
//   		ChannelsIn: jsii.Number(123),
//   		ChannelsOut: jsii.Number(123),
//   	},
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html
//
type CfnChannel_AudioDescriptionProperty struct {
	// The advanced audio normalization settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-audionormalizationsettings
	//
	AudioNormalizationSettings interface{} `field:"optional" json:"audioNormalizationSettings" yaml:"audioNormalizationSettings"`
	// The name of the AudioSelector that is used as the source for this AudioDescription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-audioselectorname
	//
	AudioSelectorName *string `field:"optional" json:"audioSelectorName" yaml:"audioSelectorName"`
	// Applies only if audioTypeControl is useConfigured.
	//
	// The values for audioType are defined in ISO-IEC 13818-1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-audiotype
	//
	AudioType *string `field:"optional" json:"audioType" yaml:"audioType"`
	// Determines how audio type is determined.
	//
	// followInput: If the input contains an ISO 639 audioType, then that value is passed through to the output. If the input contains no ISO 639 audioType, the value in Audio Type is included in the output. useConfigured: The value in Audio Type is included in the output. Note that this field and audioType are both ignored if inputType is broadcasterMixedAd.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-audiotypecontrol
	//
	AudioTypeControl *string `field:"optional" json:"audioTypeControl" yaml:"audioTypeControl"`
	// Settings to configure one or more solutions that insert audio watermarks in the audio encode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-audiowatermarkingsettings
	//
	AudioWatermarkingSettings interface{} `field:"optional" json:"audioWatermarkingSettings" yaml:"audioWatermarkingSettings"`
	// The audio codec settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-codecsettings
	//
	CodecSettings interface{} `field:"optional" json:"codecSettings" yaml:"codecSettings"`
	// Indicates the language of the audio output track.
	//
	// Used only if languageControlMode is useConfigured, or there is no ISO 639 language code specified in the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Choosing followInput causes the ISO 639 language code of the output to follow the ISO 639 language code of the input.
	//
	// The languageCode setting is used when useConfigured is set, or when followInput is selected but there is no ISO 639 language code specified by the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-languagecodecontrol
	//
	LanguageCodeControl *string `field:"optional" json:"languageCodeControl" yaml:"languageCodeControl"`
	// The name of this AudioDescription.
	//
	// Outputs use this name to uniquely identify this AudioDescription. Description names should be unique within this channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The settings that control how input audio channels are remixed into the output audio channels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-remixsettings
	//
	RemixSettings interface{} `field:"optional" json:"remixSettings" yaml:"remixSettings"`
	// Used for Microsoft Smooth and Apple HLS outputs.
	//
	// Indicates the name displayed by the player (for example, English or Director Commentary).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiodescription.html#cfn-medialive-channel-audiodescription-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

