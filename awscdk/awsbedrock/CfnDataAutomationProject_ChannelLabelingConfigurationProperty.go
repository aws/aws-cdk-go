package awsbedrock


// Enables or disables channel labeling.
//
// Channel labeling, when enabled will assign a number to each audio channel, and indicate which channel is being used in each portion of the transcript. This appears in the response as "ch_0" for the first channel, and "ch_1" for the second.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelLabelingConfigurationProperty := &ChannelLabelingConfigurationProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-channellabelingconfiguration.html
//
type CfnDataAutomationProject_ChannelLabelingConfigurationProperty struct {
	// State of channel labeling, either enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-channellabelingconfiguration.html#cfn-bedrock-dataautomationproject-channellabelingconfiguration-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

