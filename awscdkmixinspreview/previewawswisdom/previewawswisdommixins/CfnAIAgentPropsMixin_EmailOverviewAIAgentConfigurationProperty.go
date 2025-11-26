package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   emailOverviewAIAgentConfigurationProperty := &EmailOverviewAIAgentConfigurationProperty{
//   	EmailOverviewAiPromptId: jsii.String("emailOverviewAiPromptId"),
//   	Locale: jsii.String("locale"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailoverviewaiagentconfiguration.html
//
type CfnAIAgentPropsMixin_EmailOverviewAIAgentConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailoverviewaiagentconfiguration.html#cfn-wisdom-aiagent-emailoverviewaiagentconfiguration-emailoverviewaipromptid
	//
	EmailOverviewAiPromptId *string `field:"optional" json:"emailOverviewAiPromptId" yaml:"emailOverviewAiPromptId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-emailoverviewaiagentconfiguration.html#cfn-wisdom-aiagent-emailoverviewaiagentconfiguration-locale
	//
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
}

