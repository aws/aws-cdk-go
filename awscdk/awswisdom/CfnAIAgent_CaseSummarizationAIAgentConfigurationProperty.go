package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   caseSummarizationAIAgentConfigurationProperty := &CaseSummarizationAIAgentConfigurationProperty{
//   	CaseSummarizationAiGuardrailId: jsii.String("caseSummarizationAiGuardrailId"),
//   	CaseSummarizationAiPromptId: jsii.String("caseSummarizationAiPromptId"),
//   	Locale: jsii.String("locale"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-casesummarizationaiagentconfiguration.html
//
type CfnAIAgent_CaseSummarizationAIAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-casesummarizationaiagentconfiguration.html#cfn-wisdom-aiagent-casesummarizationaiagentconfiguration-casesummarizationaiguardrailid
	//
	CaseSummarizationAiGuardrailId *string `field:"optional" json:"caseSummarizationAiGuardrailId" yaml:"caseSummarizationAiGuardrailId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-casesummarizationaiagentconfiguration.html#cfn-wisdom-aiagent-casesummarizationaiagentconfiguration-casesummarizationaipromptid
	//
	CaseSummarizationAiPromptId *string `field:"optional" json:"caseSummarizationAiPromptId" yaml:"caseSummarizationAiPromptId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-casesummarizationaiagentconfiguration.html#cfn-wisdom-aiagent-casesummarizationaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

