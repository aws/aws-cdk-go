package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a CDK managed Bedrock Prompt.
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
type PromptProps struct {
	// The name of the prompt.
	//
	// This will be used as the physical name of the prompt.
	// Allowed Pattern: ^([0-9a-zA-Z][_-]?){1,100}$
	// Experimental.
	PromptName *string `field:"required" json:"promptName" yaml:"promptName"`
	// The Prompt Variant that will be used by default.
	// Default: - No default variant provided.
	//
	// Experimental.
	DefaultVariant IPromptVariant `field:"optional" json:"defaultVariant" yaml:"defaultVariant"`
	// A description of what the prompt does.
	// Default: - No description provided.
	// Maximum Length: 200.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key that the prompt is encrypted with.
	// Default: - AWS owned and managed key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Tags to apply to the prompt.
	// Default: - No tags applied.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The variants of your prompt.
	//
	// Variants can use different messages, models,
	// or configurations so that you can compare their outputs to decide the best
	// variant for your use case. Maximum of 1 variants.
	// Default: - No additional variants provided.
	//
	// Experimental.
	Variants *[]IPromptVariant `field:"optional" json:"variants" yaml:"variants"`
}

