package mixinsawsbedrock


// Settings for generating data from audio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioStandardExtractionProperty := &AudioStandardExtractionProperty{
//   	Category: &AudioExtractionCategoryProperty{
//   		State: jsii.String("state"),
//   		TypeConfiguration: &AudioExtractionCategoryTypeConfigurationProperty{
//   			Transcript: &TranscriptConfigurationProperty{
//   				ChannelLabeling: &ChannelLabelingConfigurationProperty{
//   					State: jsii.String("state"),
//   				},
//   				SpeakerLabeling: &SpeakerLabelingConfigurationProperty{
//   					State: jsii.String("state"),
//   				},
//   			},
//   		},
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardextraction.html
//
type CfnDataAutomationProjectPropsMixin_AudioStandardExtractionProperty struct {
	// Settings for generating data from audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardextraction.html#cfn-bedrock-dataautomationproject-audiostandardextraction-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
}

