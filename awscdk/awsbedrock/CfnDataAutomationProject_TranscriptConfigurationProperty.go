package awsbedrock


// Configuration for transcript options.
//
// This option allows you to enable speaker labeling and channel labeling.
//
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
	// Enables channel labeling.
	//
	// Each audio channel will be labeled with a number, and the transcript will indicate which channel is being used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-transcriptconfiguration.html#cfn-bedrock-dataautomationproject-transcriptconfiguration-channellabeling
	//
	ChannelLabeling interface{} `field:"optional" json:"channelLabeling" yaml:"channelLabeling"`
	// Enables speaker labeling.
	//
	// Each speaker within a transcript will recieve a number, and the transcript will note which speaker is talking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-transcriptconfiguration.html#cfn-bedrock-dataautomationproject-transcriptconfiguration-speakerlabeling
	//
	SpeakerLabeling interface{} `field:"optional" json:"speakerLabeling" yaml:"speakerLabeling"`
}

