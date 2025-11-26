package previewawsbedrockmixins


// Contains the message for a prompt.
//
// For more information, see [Construct and store reusable prompts with Prompt management in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var any interface{}
//   var auto interface{}
//   var json interface{}
//
//   promptTemplateConfigurationProperty := &PromptTemplateConfigurationProperty{
//   	Chat: &ChatPromptTemplateConfigurationProperty{
//   		InputVariables: []interface{}{
//   			&PromptInputVariableProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Messages: []interface{}{
//   			&MessageProperty{
//   				Content: []interface{}{
//   					&ContentBlockProperty{
//   						CachePoint: &CachePointBlockProperty{
//   							Type: jsii.String("type"),
//   						},
//   						Text: jsii.String("text"),
//   					},
//   				},
//   				Role: jsii.String("role"),
//   			},
//   		},
//   		System: []interface{}{
//   			&SystemContentBlockProperty{
//   				CachePoint: &CachePointBlockProperty{
//   					Type: jsii.String("type"),
//   				},
//   				Text: jsii.String("text"),
//   			},
//   		},
//   		ToolConfiguration: &ToolConfigurationProperty{
//   			ToolChoice: &ToolChoiceProperty{
//   				Any: any,
//   				Auto: auto,
//   				Tool: &SpecificToolChoiceProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Tools: []interface{}{
//   				&ToolProperty{
//   					CachePoint: &CachePointBlockProperty{
//   						Type: jsii.String("type"),
//   					},
//   					ToolSpec: &ToolSpecificationProperty{
//   						Description: jsii.String("description"),
//   						InputSchema: &ToolInputSchemaProperty{
//   							Json: json,
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Text: &TextPromptTemplateConfigurationProperty{
//   		CachePoint: &CachePointBlockProperty{
//   			Type: jsii.String("type"),
//   		},
//   		InputVariables: []interface{}{
//   			&PromptInputVariableProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Text: jsii.String("text"),
//   		TextS3Location: &TextS3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Key: jsii.String("key"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-prompttemplateconfiguration.html
//
type CfnPromptPropsMixin_PromptTemplateConfigurationProperty struct {
	// Contains configurations to use the prompt in a conversational format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-prompttemplateconfiguration.html#cfn-bedrock-prompt-prompttemplateconfiguration-chat
	//
	Chat interface{} `field:"optional" json:"chat" yaml:"chat"`
	// Contains configurations for the text in a message for a prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-prompttemplateconfiguration.html#cfn-bedrock-prompt-prompttemplateconfiguration-text
	//
	Text interface{} `field:"optional" json:"text" yaml:"text"`
}

