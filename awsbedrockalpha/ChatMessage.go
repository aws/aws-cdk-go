package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a message in a chat conversation.
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
type ChatMessage interface {
	// The role of the message sender.
	// Experimental.
	Role() ChatMessageRole
	// The text content of the message.
	// Experimental.
	Text() *string
}

// The jsii proxy struct for ChatMessage
type jsiiProxy_ChatMessage struct {
	_ byte // padding
}

func (j *jsiiProxy_ChatMessage) Role() ChatMessageRole {
	var returns ChatMessageRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatMessage) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}


// Experimental.
func NewChatMessage(role ChatMessageRole, text *string) ChatMessage {
	_init_.Initialize()

	if err := validateNewChatMessageParameters(role, text); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChatMessage{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.ChatMessage",
		[]interface{}{role, text},
		&j,
	)

	return &j
}

// Experimental.
func NewChatMessage_Override(c ChatMessage, role ChatMessageRole, text *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.ChatMessage",
		[]interface{}{role, text},
		c,
	)
}

// Creates an assistant message.
//
// Returns: A ChatMessage instance representing an assistant message.
// Experimental.
func ChatMessage_Assistant(text *string) ChatMessage {
	_init_.Initialize()

	if err := validateChatMessage_AssistantParameters(text); err != nil {
		panic(err)
	}
	var returns ChatMessage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.ChatMessage",
		"assistant",
		[]interface{}{text},
		&returns,
	)

	return returns
}

// Creates a user message.
//
// Returns: A ChatMessage instance representing a user message.
// Experimental.
func ChatMessage_User(text *string) ChatMessage {
	_init_.Initialize()

	if err := validateChatMessage_UserParameters(text); err != nil {
		panic(err)
	}
	var returns ChatMessage

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.ChatMessage",
		"user",
		[]interface{}{text},
		&returns,
	)

	return returns
}

