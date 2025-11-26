package previewawsbedrockmixins


// Contains configurations for a text prompt template.
//
// To include a variable, enclose a word in double curly braces as in `{{variable}}` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   textPromptTemplateConfigurationProperty := &TextPromptTemplateConfigurationProperty{
//   	InputVariables: []interface{}{
//   		&PromptInputVariableProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-textprompttemplateconfiguration.html
//
type CfnFlowVersionPropsMixin_TextPromptTemplateConfigurationProperty struct {
	// An array of the variables in the prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-textprompttemplateconfiguration.html#cfn-bedrock-flowversion-textprompttemplateconfiguration-inputvariables
	//
	InputVariables interface{} `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// The message for the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-textprompttemplateconfiguration.html#cfn-bedrock-flowversion-textprompttemplateconfiguration-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

