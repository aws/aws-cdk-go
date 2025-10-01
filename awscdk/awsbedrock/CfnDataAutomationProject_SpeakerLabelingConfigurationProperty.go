package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   speakerLabelingConfigurationProperty := &SpeakerLabelingConfigurationProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-speakerlabelingconfiguration.html
//
type CfnDataAutomationProject_SpeakerLabelingConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-speakerlabelingconfiguration.html#cfn-bedrock-dataautomationproject-speakerlabelingconfiguration-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

