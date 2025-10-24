package awsbedrockalpha


// Properties for creating a CDK managed Bedrock Prompt Version.
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
//   promptVersion := bedrock.NewPromptVersion(this, jsii.String("MyPromptVersion"), &PromptVersionProps{
//   	Prompt: prompt1,
//   	Description: jsii.String("my first version"),
//   })
//   //or alternatively:
//   // const promptVersion = prompt1.createVersion('my first version');
//   versionString := promptVersion.Version
//
// Experimental.
type PromptVersionProps struct {
	// The prompt to use for this version.
	// Experimental.
	Prompt IPrompt `field:"required" json:"prompt" yaml:"prompt"`
	// The description of the prompt version.
	// Default: - No description provided.
	// Maximum length: 200.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

