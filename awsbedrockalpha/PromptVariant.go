package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating prompt variants.
//
// Provides static methods to create different types of prompt variants
// with proper configuration and type safety.
//
// Example:
//   cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
//   })
//   claudeModel := bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_SONNET_V1_0()
//
//   variant1 := bedrock.PromptVariant_Text(&TextPromptVariantProps{
//   	VariantName: jsii.String("variant1"),
//   	Model: claudeModel,
//   	PromptVariables: []*string{
//   		jsii.String("topic"),
//   	},
//   	PromptText: jsii.String("This is my first text prompt. Please summarize our conversation on: {{topic}}."),
//   	InferenceConfiguration: bedrock.PromptInferenceConfiguration_Text(&PromptInferenceConfigurationProps{
//   		Temperature: jsii.Number(1),
//   		TopP: jsii.Number(0.999),
//   		MaxTokens: jsii.Number(2000),
//   	}),
//   })
//
//   prompt1 := bedrock.NewPrompt(this, jsii.String("prompt1"), &PromptProps{
//   	PromptName: jsii.String("prompt1"),
//   	Description: jsii.String("my first prompt"),
//   	DefaultVariant: variant1,
//   	Variants: []IPromptVariant{
//   		variant1,
//   	},
//   	KmsKey: cmk,
//   })
//
// Experimental.
type PromptVariant interface {
}

// The jsii proxy struct for PromptVariant
type jsiiProxy_PromptVariant struct {
	_ byte // padding
}

// Creates an agent prompt template variant.
//
// Returns: A PromptVariant configured for agent interactions.
// Experimental.
func PromptVariant_Agent(props *AgentPromptVariantProps) IPromptVariant {
	_init_.Initialize()

	if err := validatePromptVariant_AgentParameters(props); err != nil {
		panic(err)
	}
	var returns IPromptVariant

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptVariant",
		"agent",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a chat template variant.
//
// Use this template type when
// the model supports the Converse API or the Anthropic Claude Messages API.
// This allows you to include a System prompt and previous User messages
// and Assistant messages for context.
//
// Returns: A PromptVariant configured for chat interactions.
// Experimental.
func PromptVariant_Chat(props *ChatPromptVariantProps) IPromptVariant {
	_init_.Initialize()

	if err := validatePromptVariant_ChatParameters(props); err != nil {
		panic(err)
	}
	var returns IPromptVariant

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptVariant",
		"chat",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a text template variant.
//
// Returns: A PromptVariant configured for text processing.
// Experimental.
func PromptVariant_Text(props *TextPromptVariantProps) IPromptVariant {
	_init_.Initialize()

	if err := validatePromptVariant_TextParameters(props); err != nil {
		panic(err)
	}
	var returns IPromptVariant

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptVariant",
		"text",
		[]interface{}{props},
		&returns,
	)

	return returns
}

