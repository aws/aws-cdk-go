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
//   audioSelectorSettingsProperty := &audioSelectorSettingsProperty{
//   	audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   		groupId: jsii.String("groupId"),
//   		name: jsii.String("name"),
//   	},
//   	audioLanguageSelection: &audioLanguageSelectionProperty{
//   		languageCode: jsii.String("languageCode"),
//   		languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   	},
//   	audioPidSelection: &audioPidSelectionProperty{
//   		pid: jsii.Number(123),
//   	},
//   	audioTrackSelection: &audioTrackSelectionProperty{
//   		tracks: []interface{}{
//   			&audioTrackProperty{
//   				track: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnChannel_AudioSelectorSettingsProperty struct {
	// Selector for HLS audio rendition.
	AudioHlsRenditionSelection interface{} `field:"optional" json:"audioHlsRenditionSelection" yaml:"audioHlsRenditionSelection"`
	// The language code of the audio to select.
	AudioLanguageSelection interface{} `field:"optional" json:"audioLanguageSelection" yaml:"audioLanguageSelection"`
	// The PID of the audio to select.
	AudioPidSelection interface{} `field:"optional" json:"audioPidSelection" yaml:"audioPidSelection"`
	// Information about the audio track to extract.
	AudioTrackSelection interface{} `field:"optional" json:"audioTrackSelection" yaml:"audioTrackSelection"`
}

