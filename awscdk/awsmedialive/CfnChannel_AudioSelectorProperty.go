package awsmedialive


// Information about one audio to extract from the input.
//
// The parent of this entity is InputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioSelectorProperty := &AudioSelectorProperty{
//   	Name: jsii.String("name"),
//   	SelectorSettings: &AudioSelectorSettingsProperty{
//   		AudioHlsRenditionSelection: &AudioHlsRenditionSelectionProperty{
//   			GroupId: jsii.String("groupId"),
//   			Name: jsii.String("name"),
//   		},
//   		AudioLanguageSelection: &AudioLanguageSelectionProperty{
//   			LanguageCode: jsii.String("languageCode"),
//   			LanguageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   		},
//   		AudioPidSelection: &AudioPidSelectionProperty{
//   			Pid: jsii.Number(123),
//   		},
//   		AudioTrackSelection: &AudioTrackSelectionProperty{
//   			Tracks: []interface{}{
//   				&AudioTrackProperty{
//   					Track: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnChannel_AudioSelectorProperty struct {
	// A name for this AudioSelector.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the specific audio to extract from the input.
	SelectorSettings interface{} `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

