package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract base class for prompt template configurations.
//
// This provides a high-level abstraction over the underlying CloudFormation
// template configuration properties, offering a more developer-friendly interface
// while maintaining full compatibility with the underlying AWS Bedrock service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var chatMessage ChatMessage
//   var tool Tool
//   var toolChoice ToolChoice
//
//   promptTemplateConfiguration := bedrock_alpha.PromptTemplateConfiguration_Chat(&ChatTemplateConfigurationProps{
//   	Messages: []ChatMessage{
//   		chatMessage,
//   	},
//
//   	// the properties below are optional
//   	InputVariables: []*string{
//   		jsii.String("inputVariables"),
//   	},
//   	System: jsii.String("system"),
//   	ToolConfiguration: &ToolConfiguration{
//   		ToolChoice: toolChoice,
//   		Tools: []Tool{
//   			tool,
//   		},
//   	},
//   })
//
// Experimental.
type PromptTemplateConfiguration interface {
}

// The jsii proxy struct for PromptTemplateConfiguration
type jsiiProxy_PromptTemplateConfiguration struct {
	_ byte // padding
}

// Experimental.
func NewPromptTemplateConfiguration_Override(p PromptTemplateConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.PromptTemplateConfiguration",
		nil, // no parameters
		p,
	)
}

// Creates a chat template configuration.
// Experimental.
func PromptTemplateConfiguration_Chat(props *ChatTemplateConfigurationProps) PromptTemplateConfiguration {
	_init_.Initialize()

	if err := validatePromptTemplateConfiguration_ChatParameters(props); err != nil {
		panic(err)
	}
	var returns PromptTemplateConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptTemplateConfiguration",
		"chat",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a text template configuration.
// Experimental.
func PromptTemplateConfiguration_Text(props *TextTemplateConfigurationProps) PromptTemplateConfiguration {
	_init_.Initialize()

	if err := validatePromptTemplateConfiguration_TextParameters(props); err != nil {
		panic(err)
	}
	var returns PromptTemplateConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptTemplateConfiguration",
		"text",
		[]interface{}{props},
		&returns,
	)

	return returns
}

