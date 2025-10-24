package awsbedrockalpha


// Properties for creating a prompt inference configuration.
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
type PromptInferenceConfigurationProps struct {
	// The maximum number of tokens to return in the response.
	// Default: - No limit specified.
	//
	// Experimental.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// A list of strings that define sequences after which the model will stop generating.
	// Default: - No stop sequences.
	//
	// Experimental.
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Controls the randomness of the response.
	//
	// Higher values make output more random, lower values more deterministic.
	// Valid range is 0.0 to 1.0.
	// Default: - Model default temperature.
	//
	// Experimental.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// The percentage of most-likely candidates that the model considers for the next token.
	//
	// Valid range is 0.0 to 1.0.
	// Default: - Model default topP.
	//
	// Experimental.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}

