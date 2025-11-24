package mixinsawsbedrock


// Output settings for processing audio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioStandardOutputConfigurationProperty := &AudioStandardOutputConfigurationProperty{
//   	Extraction: &AudioStandardExtractionProperty{
//   		Category: &AudioExtractionCategoryProperty{
//   			State: jsii.String("state"),
//   			TypeConfiguration: &AudioExtractionCategoryTypeConfigurationProperty{
//   				Transcript: &TranscriptConfigurationProperty{
//   					ChannelLabeling: &ChannelLabelingConfigurationProperty{
//   						State: jsii.String("state"),
//   					},
//   					SpeakerLabeling: &SpeakerLabelingConfigurationProperty{
//   						State: jsii.String("state"),
//   					},
//   				},
//   			},
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   	GenerativeField: &AudioStandardGenerativeFieldProperty{
//   		State: jsii.String("state"),
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardoutputconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_AudioStandardOutputConfigurationProperty struct {
	// Settings for populating data fields that describe the audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-audiostandardoutputconfiguration-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// Whether to generate descriptions of the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-audiostandardoutputconfiguration-generativefield
	//
	GenerativeField interface{} `field:"optional" json:"generativeField" yaml:"generativeField"`
}

