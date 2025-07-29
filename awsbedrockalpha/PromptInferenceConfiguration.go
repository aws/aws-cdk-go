package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract base class for prompt inference configurations.
//
// This provides a high-level abstraction over the underlying CloudFormation
// inference configuration properties, offering a more developer-friendly interface
// while maintaining full compatibility with the underlying AWS Bedrock service.
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
//   	Variants: []iPromptVariant{
//   		variant1,
//   	},
//   	KmsKey: cmk,
//   })
//
// Experimental.
type PromptInferenceConfiguration interface {
}

// The jsii proxy struct for PromptInferenceConfiguration
type jsiiProxy_PromptInferenceConfiguration struct {
	_ byte // padding
}

// Experimental.
func NewPromptInferenceConfiguration_Override(p PromptInferenceConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.PromptInferenceConfiguration",
		nil, // no parameters
		p,
	)
}

// Creates a text inference configuration.
// Experimental.
func PromptInferenceConfiguration_Text(props *PromptInferenceConfigurationProps) PromptInferenceConfiguration {
	_init_.Initialize()

	if err := validatePromptInferenceConfiguration_TextParameters(props); err != nil {
		panic(err)
	}
	var returns PromptInferenceConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptInferenceConfiguration",
		"text",
		[]interface{}{props},
		&returns,
	)

	return returns
}

