package awswisdom


// Configuration settings for the EMAIL_OVERVIEW AI agent including prompt ID and locale settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emailOverviewAIAgentConfigurationProperty := &EmailOverviewAIAgentConfigurationProperty{
//   	EmailOverviewAiPromptId: jsii.String("emailOverviewAiPromptId"),
//   	Locale: jsii.String("locale"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailoverviewaiagentconfiguration.html
//
type CfnAIAgent_EmailOverviewAIAgentConfigurationProperty struct {
	// The ID of the System AI prompt used for generating structured email conversation summaries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailoverviewaiagentconfiguration.html#cfn-wisdom-aiagent-emailoverviewaiagentconfiguration-emailoverviewaipromptid
	//
	EmailOverviewAiPromptId *string `field:"optional" json:"emailOverviewAiPromptId" yaml:"emailOverviewAiPromptId"`
	// The locale setting for language-specific email overview processing (for example, en_US, es_ES).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailoverviewaiagentconfiguration.html#cfn-wisdom-aiagent-emailoverviewaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

