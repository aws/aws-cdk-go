package awsbedrock


// Configuration for chat prompt template.
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
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	ToolConfiguration: &ToolConfigurationProperty{
//   		Tools: []interface{}{
//   			&ToolProperty{
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
	// List of messages for chat prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html#cfn-bedrock-promptversion-chatprompttemplateconfiguration-messages
	//
	Messages interface{} `field:"required" json:"messages" yaml:"messages"`
	// List of input variables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html#cfn-bedrock-promptversion-chatprompttemplateconfiguration-inputvariables
	//
	InputVariables interface{} `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// Configuration for chat prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html#cfn-bedrock-promptversion-chatprompttemplateconfiguration-system
	//
	System interface{} `field:"optional" json:"system" yaml:"system"`
	// Tool configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-chatprompttemplateconfiguration.html#cfn-bedrock-promptversion-chatprompttemplateconfiguration-toolconfiguration
	//
	ToolConfiguration interface{} `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

