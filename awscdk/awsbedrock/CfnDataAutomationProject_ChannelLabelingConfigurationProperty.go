package awsbedrock


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-channellabelingconfiguration.html#cfn-bedrock-dataautomationproject-channellabelingconfiguration-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

