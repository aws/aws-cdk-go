package mixinsawsbedrock


// Enables or disables speaker labeling.
//
// Speaker labeling, when enabled will assign a number to each speaker, and indicate which speaker is talking in each portion of the transcript. This appears in the response as "spk_0" for the first speaker, "spk_1" for the second, and so on for up to 30 speakers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   speakerLabelingConfigurationProperty := &SpeakerLabelingConfigurationProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-speakerlabelingconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_SpeakerLabelingConfigurationProperty struct {
	// State of speaker labeling, either enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-speakerlabelingconfiguration.html#cfn-bedrock-dataautomationproject-speakerlabelingconfiguration-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

