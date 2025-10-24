package awsbedrockalpha


// Properties for creating a function tool.
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
type FunctionToolProps struct {
	// A description of what the function does.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The input schema for the function parameters.
	// Experimental.
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
	// The name of the function.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

