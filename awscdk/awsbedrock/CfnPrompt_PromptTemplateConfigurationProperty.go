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
//   		InputVariables: []interface{}{
//   			&PromptInputVariableProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Text: jsii.String("text"),
//   		TextS3Location: &TextS3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			Version: jsii.String("version"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-prompttemplateconfiguration.html
//
type CfnPrompt_PromptTemplateConfigurationProperty struct {
	// Contains configurations for the text in a message for a prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-prompttemplateconfiguration.html#cfn-bedrock-prompt-prompttemplateconfiguration-text
	//
	Text interface{} `field:"required" json:"text" yaml:"text"`
}

