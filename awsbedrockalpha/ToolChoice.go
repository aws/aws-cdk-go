package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines how the model should choose which tool to use.
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
type ToolChoice interface {
	// Configuration for ANY tool choice.
	// Experimental.
	Any() interface{}
	// Configuration for AUTO tool choice.
	// Experimental.
	Auto() interface{}
	// The specific tool name if using specific tool choice.
	// Experimental.
	Tool() *string
}

// The jsii proxy struct for ToolChoice
type jsiiProxy_ToolChoice struct {
	_ byte // padding
}

func (j *jsiiProxy_ToolChoice) Any() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"any",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolChoice) Auto() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolChoice) Tool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}


// Experimental.
func NewToolChoice(any interface{}, auto interface{}, tool *string) ToolChoice {
	_init_.Initialize()

	if err := validateNewToolChoiceParameters(any, auto); err != nil {
		panic(err)
	}
	j := jsiiProxy_ToolChoice{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.ToolChoice",
		[]interface{}{any, auto, tool},
		&j,
	)

	return &j
}

// Experimental.
func NewToolChoice_Override(t ToolChoice, any interface{}, auto interface{}, tool *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.ToolChoice",
		[]interface{}{any, auto, tool},
		t,
	)
}

// The Model must request the specified tool.
//
// Only supported by some models like Anthropic Claude 3 models.
//
// Returns: A ToolChoice instance configured for the specific tool.
// Experimental.
func ToolChoice_SpecificTool(toolName *string) ToolChoice {
	_init_.Initialize()

	if err := validateToolChoice_SpecificToolParameters(toolName); err != nil {
		panic(err)
	}
	var returns ToolChoice

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.ToolChoice",
		"specificTool",
		[]interface{}{toolName},
		&returns,
	)

	return returns
}

func ToolChoice_ANY() ToolChoice {
	_init_.Initialize()
	var returns ToolChoice
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.ToolChoice",
		"ANY",
		&returns,
	)
	return returns
}

func ToolChoice_AUTO() ToolChoice {
	_init_.Initialize()
	var returns ToolChoice
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.ToolChoice",
		"AUTO",
		&returns,
	)
	return returns
}

