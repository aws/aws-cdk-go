package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract base class for tools that can be used by the model.
//
// Example:
//   cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
//   })
//
//   variantChat := bedrock.PromptVariant_Chat(&ChatPromptVariantProps{
//   	VariantName: jsii.String("variant1"),
//   	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
//   	Messages: []ChatMessage{
//   		bedrock.ChatMessage_User(jsii.String("From now on, you speak Japanese!")),
//   		bedrock.ChatMessage_Assistant(jsii.String("Konnichiwa!")),
//   		bedrock.ChatMessage_*User(jsii.String("From now on, you speak {{language}}!")),
//   	},
//   	System: jsii.String("You are a helpful assistant that only speaks the language you`re told."),
//   	PromptVariables: []*string{
//   		jsii.String("language"),
//   	},
//   	ToolConfiguration: &ToolConfiguration{
//   		ToolChoice: bedrock.ToolChoice_AUTO(),
//   		Tools: []Tool{
//   			bedrock.Tool_Function(&FunctionToolProps{
//   				Name: jsii.String("top_song"),
//   				Description: jsii.String("Get the most popular song played on a radio station."),
//   				InputSchema: map[string]interface{}{
//   					"type": jsii.String("object"),
//   					"properties": map[string]map[string]*string{
//   						"sign": map[string]*string{
//   							"type": jsii.String("string"),
//   							"description": jsii.String("The call sign for the radio station for which you want the most popular song. Example calls signs are WZPZ and WKR."),
//   						},
//   					},
//   					"required": []*string{
//   						jsii.String("sign"),
//   					},
//   				},
//   			}),
//   		},
//   	},
//   })
//
//   bedrock.NewPrompt(this, jsii.String("prompt1"), &PromptProps{
//   	PromptName: jsii.String("prompt-chat"),
//   	Description: jsii.String("my first chat prompt"),
//   	DefaultVariant: variantChat,
//   	Variants: []IPromptVariant{
//   		variantChat,
//   	},
//   	KmsKey: cmk,
//   })
//
// Experimental.
type Tool interface {
}

// The jsii proxy struct for Tool
type jsiiProxy_Tool struct {
	_ byte // padding
}

// Experimental.
func NewTool_Override(t Tool) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.Tool",
		nil, // no parameters
		t,
	)
}

// Creates a function tool.
// Experimental.
func Tool_Function(props *FunctionToolProps) Tool {
	_init_.Initialize()

	if err := validateTool_FunctionParameters(props); err != nil {
		panic(err)
	}
	var returns Tool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.Tool",
		"function",
		[]interface{}{props},
		&returns,
	)

	return returns
}

