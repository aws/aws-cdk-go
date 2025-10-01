package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnDataAutomationProject_AudioExtractionCategoryTypeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audioextractioncategorytypeconfiguration.html#cfn-bedrock-dataautomationproject-audioextractioncategorytypeconfiguration-transcript
	//
	Transcript interface{} `field:"optional" json:"transcript" yaml:"transcript"`
}

