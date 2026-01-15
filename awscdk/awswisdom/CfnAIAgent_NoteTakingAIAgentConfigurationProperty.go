package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noteTakingAIAgentConfigurationProperty := &NoteTakingAIAgentConfigurationProperty{
//   	Locale: jsii.String("locale"),
//   	NoteTakingAiGuardrailId: jsii.String("noteTakingAiGuardrailId"),
//   	NoteTakingAiPromptId: jsii.String("noteTakingAiPromptId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-notetakingaiagentconfiguration.html
//
type CfnAIAgent_NoteTakingAIAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-notetakingaiagentconfiguration.html#cfn-wisdom-aiagent-notetakingaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-notetakingaiagentconfiguration.html#cfn-wisdom-aiagent-notetakingaiagentconfiguration-notetakingaiguardrailid
	//
	NoteTakingAiGuardrailId *string `field:"optional" json:"noteTakingAiGuardrailId" yaml:"noteTakingAiGuardrailId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-notetakingaiagentconfiguration.html#cfn-wisdom-aiagent-notetakingaiagentconfiguration-notetakingaipromptid
	//
	NoteTakingAiPromptId *string `field:"optional" json:"noteTakingAiPromptId" yaml:"noteTakingAiPromptId"`
}

