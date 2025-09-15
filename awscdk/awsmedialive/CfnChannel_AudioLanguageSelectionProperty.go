package awsmedialive


// Information about the audio language to extract.
//
// The parent of this entity is AudioSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLanguageSelectionProperty := &AudioLanguageSelectionProperty{
//   	LanguageCode: jsii.String("languageCode"),
//   	LanguageSelectionPolicy: jsii.String("languageSelectionPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiolanguageselection.html
//
type CfnChannel_AudioLanguageSelectionProperty struct {
	// Selects a specific three-letter language code from within an audio source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiolanguageselection.html#cfn-medialive-channel-audiolanguageselection-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// When set to "strict," the transport stream demux strictly identifies audio streams by their language descriptor.
	//
	// If a PMT update occurs such that an audio stream matching the initially selected language is no longer present, then mute is encoded until the language returns. If set to "loose," then on a PMT update the demux chooses another audio stream in the program with the same stream type if it can't find one with the same language.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiolanguageselection.html#cfn-medialive-channel-audiolanguageselection-languageselectionpolicy
	//
	LanguageSelectionPolicy *string `field:"optional" json:"languageSelectionPolicy" yaml:"languageSelectionPolicy"`
}

