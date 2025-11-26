package previewawsbedrockmixins


// Allows configuration of extractions for different types of data, such as transcript and content moderation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioExtractionCategoryTypeConfigurationProperty := &AudioExtractionCategoryTypeConfigurationProperty{
//   	Transcript: &TranscriptConfigurationProperty{
//   		ChannelLabeling: &ChannelLabelingConfigurationProperty{
//   			State: jsii.String("state"),
//   		},
//   		SpeakerLabeling: &SpeakerLabelingConfigurationProperty{
//   			State: jsii.String("state"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audioextractioncategorytypeconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_AudioExtractionCategoryTypeConfigurationProperty struct {
	// This element allows you to configure different extractions for your transcript data, such as speaker and channel labeling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audioextractioncategorytypeconfiguration.html#cfn-bedrock-dataautomationproject-audioextractioncategorytypeconfiguration-transcript
	//
	Transcript interface{} `field:"optional" json:"transcript" yaml:"transcript"`
}

