package awsbedrock


// Contains configurations to use a prompt in a conversational format.
//
// For more information, see [Create a prompt using Prompt management](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-create.html) .
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
//   chatPromptTemplateConfigurationProperty := &ChatPromptTemplateConfigurationProperty{
//   	Messages: []interface{}{
//   		&MessageProperty{
//   			Content: []interface{}{
//   				&ContentBlockProperty{
//   					CachePoint: &CachePointBlockProperty{
//   						Type: jsii.String("type"),
//   					},
//   					Text: jsii.String("text"),
//   				},
//   			},
//   			Role: jsii.String("role"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	InputVariables: []interface{}{
//   		&PromptInputVariableProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	System: []interface{}{
//   		&SystemContentBlockProperty{
//   			CachePoint: &CachePointBlockProperty{
//   				Type: jsii.String("type"),
//   			},
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	ToolConfiguration: &ToolConfigurationProperty{
//   		Tools: []interface{}{
//   			&ToolProperty{
//   				CachePoint: &CachePointBlockProperty{
//   					Type: jsii.String("type"),
//   				},
//   				ToolSpec: &ToolSpecificationProperty{
//   					InputSchema: &ToolInputSchemaProperty{
//   						Json: json,
//   					},
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		ToolChoice: &ToolChoiceProperty{
//   			Any: any,
//   			Auto: auto,
//   			Tool: &SpecificToolChoiceProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html
//
type CfnPromptVersion_ChatPromptTemplateConfigurationProperty struct {
	// Contains messages in the chat for the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html#cfn-bedrock-promptversion-chatprompttemplateconfiguration-messages
	//
	Messages interface{} `field:"required" json:"messages" yaml:"messages"`
	// An array of the variables in the prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html#cfn-bedrock-promptversion-chatprompttemplateconfiguration-inputvariables
	//
	InputVariables interface{} `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// Contains system prompts to provide context to the model or to describe how it should behave.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html#cfn-bedrock-promptversion-chatprompttemplateconfiguration-system
	//
	System interface{} `field:"optional" json:"system" yaml:"system"`
	// Configuration information for the tools that the model can use when generating a response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html#cfn-bedrock-promptversion-chatprompttemplateconfiguration-toolconfiguration
	//
	ToolConfiguration interface{} `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

