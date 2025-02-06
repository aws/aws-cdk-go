package awsbedrock


// Contains the message for a prompt.
//
// For more information, see [Construct and store reusable prompts with Prompt management in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var any interface{}
//   var auto interface{}
//   var json interface{}
//
//   promptTemplateConfigurationProperty := &PromptTemplateConfigurationProperty{
//   	Chat: &ChatPromptTemplateConfigurationProperty{
//   		Messages: []interface{}{
//   			&MessageProperty{
//   				Content: []interface{}{
//   					&ContentBlockProperty{
//   						Text: jsii.String("text"),
//   					},
//   				},
//   				Role: jsii.String("role"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		InputVariables: []interface{}{
//   			&PromptInputVariableProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		System: []interface{}{
//   			&SystemContentBlockProperty{
//   				Text: jsii.String("text"),
//   			},
//   		},
//   		ToolConfiguration: &ToolConfigurationProperty{
//   			Tools: []interface{}{
//   				&ToolProperty{
//   					ToolSpec: &ToolSpecificationProperty{
//   						InputSchema: &ToolInputSchemaProperty{
//   							Json: json,
//   						},
//   						Name: jsii.String("name"),
//
//   						// the properties below are optional
//   						Description: jsii.String("description"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			ToolChoice: &ToolChoiceProperty{
//   				Any: any,
//   				Auto: auto,
//   				Tool: &SpecificToolChoiceProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-prompttemplateconfiguration.html
//
type CfnPromptVersion_PromptTemplateConfigurationProperty struct {
	// Contains configurations to use the prompt in a conversational format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-prompttemplateconfiguration.html#cfn-bedrock-promptversion-prompttemplateconfiguration-chat
	//
	Chat interface{} `field:"optional" json:"chat" yaml:"chat"`
	// Contains configurations for the text in a message for a prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-prompttemplateconfiguration.html#cfn-bedrock-promptversion-prompttemplateconfiguration-text
	//
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

