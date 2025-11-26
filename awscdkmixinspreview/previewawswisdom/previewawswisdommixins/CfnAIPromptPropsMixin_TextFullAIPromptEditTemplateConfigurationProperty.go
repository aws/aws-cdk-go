package previewawswisdommixins


// The configuration for a prompt template that supports full textual prompt configuration using a YAML prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   textFullAIPromptEditTemplateConfigurationProperty := &TextFullAIPromptEditTemplateConfigurationProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-textfullaipromptedittemplateconfiguration.html
//
type CfnAIPromptPropsMixin_TextFullAIPromptEditTemplateConfigurationProperty struct {
	// The YAML text for the AI Prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiprompt-textfullaipromptedittemplateconfiguration.html#cfn-wisdom-aiprompt-textfullaipromptedittemplateconfiguration-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

