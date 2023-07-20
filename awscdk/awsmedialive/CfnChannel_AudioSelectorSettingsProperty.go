package awsmedialive


// Information about the audio to extract from the input.
//
// The parent of this entity is AudioSelector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSelectorSettingsProperty := &AudioSelectorSettingsProperty{
//   	AudioHlsRenditionSelection: &AudioHlsRenditionSelectionProperty{
//   		GroupId: jsii.String("groupId"),
//   		Name: jsii.String("name"),
//   	},
//   	AudioLanguageSelection: &AudioLanguageSelectionProperty{
//   		LanguageCode: jsii.String("languageCode"),
//   		LanguageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   	},
//   	AudioPidSelection: &AudioPidSelectionProperty{
//   		Pid: jsii.Number(123),
//   	},
//   	AudioTrackSelection: &AudioTrackSelectionProperty{
//   		DolbyEDecode: &AudioDolbyEDecodeProperty{
//   			ProgramSelection: jsii.String("programSelection"),
//   		},
//   		Tracks: []interface{}{
//   			&AudioTrackProperty{
//   				Track: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html
//
type CfnChannel_AudioSelectorSettingsProperty struct {
	// Selector for HLS audio rendition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html#cfn-medialive-channel-audioselectorsettings-audiohlsrenditionselection
	//
	AudioHlsRenditionSelection interface{} `field:"optional" json:"audioHlsRenditionSelection" yaml:"audioHlsRenditionSelection"`
	// The language code of the audio to select.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html#cfn-medialive-channel-audioselectorsettings-audiolanguageselection
	//
	AudioLanguageSelection interface{} `field:"optional" json:"audioLanguageSelection" yaml:"audioLanguageSelection"`
	// The PID of the audio to select.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html#cfn-medialive-channel-audioselectorsettings-audiopidselection
	//
	AudioPidSelection interface{} `field:"optional" json:"audioPidSelection" yaml:"audioPidSelection"`
	// Information about the audio track to extract.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html#cfn-medialive-channel-audioselectorsettings-audiotrackselection
	//
	AudioTrackSelection interface{} `field:"optional" json:"audioTrackSelection" yaml:"audioTrackSelection"`
}

