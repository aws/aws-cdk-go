package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionSummarizationAIAgentConfigurationProperty := &SessionSummarizationAIAgentConfigurationProperty{
//   	Locale: jsii.String("locale"),
//   	SessionSummarizationAiPromptId: jsii.String("sessionSummarizationAiPromptId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-sessionsummarizationaiagentconfiguration.html
//
type CfnAIAgent_SessionSummarizationAIAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-sessionsummarizationaiagentconfiguration.html#cfn-wisdom-aiagent-sessionsummarizationaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-sessionsummarizationaiagentconfiguration.html#cfn-wisdom-aiagent-sessionsummarizationaiagentconfiguration-sessionsummarizationaipromptid
	//
	SessionSummarizationAiPromptId *string `field:"optional" json:"sessionSummarizationAiPromptId" yaml:"sessionSummarizationAiPromptId"`
}

