package mixinsawswisdom


// A typed union that specifies the configuration for a prompt template based on its type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aIPromptTemplateConfigurationProperty := &AIPromptTemplateConfigurationProperty{
//   	TextFullAiPromptEditTemplateConfiguration: &TextFullAIPromptEditTemplateConfigurationProperty{
//   		Text: jsii.String("text"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-aiprompttemplateconfiguration.html
//
type CfnAIPromptPropsMixin_AIPromptTemplateConfigurationProperty struct {
	// The configuration for a prompt template that supports full textual prompt configuration using a YAML prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-aiprompttemplateconfiguration.html#cfn-wisdom-aiprompt-aiprompttemplateconfiguration-textfullaipromptedittemplateconfiguration
	//
	TextFullAiPromptEditTemplateConfiguration interface{} `field:"optional" json:"textFullAiPromptEditTemplateConfiguration" yaml:"textFullAiPromptEditTemplateConfiguration"`
}

