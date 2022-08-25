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
//   audioSelectorProperty := &audioSelectorProperty{
//   	name: jsii.String("name"),
//   	selectorSettings: &audioSelectorSettingsProperty{
//   		audioHlsRenditionSelection: &audioHlsRenditionSelectionProperty{
//   			groupId: jsii.String("groupId"),
//   			name: jsii.String("name"),
//   		},
//   		audioLanguageSelection: &audioLanguageSelectionProperty{
//   			languageCode: jsii.String("languageCode"),
//   			languageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   		},
//   		audioPidSelection: &audioPidSelectionProperty{
//   			pid: jsii.Number(123),
//   		},
//   		audioTrackSelection: &audioTrackSelectionProperty{
//   			tracks: []interface{}{
//   				&audioTrackProperty{
//   					track: jsii.Number(123),
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

