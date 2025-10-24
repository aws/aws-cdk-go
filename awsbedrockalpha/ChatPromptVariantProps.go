package awsbedrockalpha


// Properties for creating a chat prompt variant.
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
type ChatPromptVariantProps struct {
	// The model which is used to run the prompt.
	//
	// The model could be a foundation
	// model, a custom model, or a provisioned model.
	// Experimental.
	Model IBedrockInvokable `field:"required" json:"model" yaml:"model"`
	// The name of the prompt variant.
	// Experimental.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The variables in the prompt template that can be filled in at runtime.
	// Default: - No variables defined.
	//
	// Experimental.
	PromptVariables *[]*string `field:"optional" json:"promptVariables" yaml:"promptVariables"`
	// The messages in the chat prompt.
	//
	// Must include at least one User Message.
	// The messages should alternate between User and Assistant.
	// Experimental.
	Messages *[]ChatMessage `field:"required" json:"messages" yaml:"messages"`
	// Inference configuration for the Chat Prompt.
	// Default: - No inference configuration provided.
	//
	// Experimental.
	InferenceConfiguration PromptInferenceConfiguration `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Context or instructions for the model to consider before generating a response.
	// Default: - No system message provided.
	//
	// Experimental.
	System *string `field:"optional" json:"system" yaml:"system"`
	// The configuration with available tools to the model and how it must use them.
	// Default: - No tool configuration provided.
	//
	// Experimental.
	ToolConfiguration *ToolConfiguration `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

