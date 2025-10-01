package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transcriptConfigurationProperty := &TranscriptConfigurationProperty{
//   	ChannelLabeling: &ChannelLabelingConfigurationProperty{
//   		State: jsii.String("state"),
//   	},
//   	SpeakerLabeling: &SpeakerLabelingConfigurationProperty{
//   		State: jsii.String("state"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-transcriptconfiguration.html
//
type CfnDataAutomationProject_TranscriptConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-transcriptconfiguration.html#cfn-bedrock-dataautomationproject-transcriptconfiguration-channellabeling
	//
	ChannelLabeling interface{} `field:"optional" json:"channelLabeling" yaml:"channelLabeling"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-transcriptconfiguration.html#cfn-bedrock-dataautomationproject-transcriptconfiguration-speakerlabeling
	//
	SpeakerLabeling interface{} `field:"optional" json:"speakerLabeling" yaml:"speakerLabeling"`
}

