package awsbedrock


// Contains the message for a prompt.
//
// For more information, see [Prompt management in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptTemplateConfigurationProperty := &PromptTemplateConfigurationProperty{
//   	Text: &TextPromptTemplateConfigurationProperty{
//   		Text: jsii.String("text"),
//
//   		// the properties below are optional
//   		InputVariables: []interface{}{
//   			&PromptInputVariableProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-prompttemplateconfiguration.html
//
type CfnFlow_PromptTemplateConfigurationProperty struct {
	// Contains configurations for the text in a message for a prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-prompttemplateconfiguration.html#cfn-bedrock-flow-prompttemplateconfiguration-text
	//
	Text interface{} `field:"required" json:"text" yaml:"text"`
}

