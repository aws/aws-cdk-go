package awsbedrockalpha


// Configuration for tools available to the model.
//
// Example:
//   cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
//   })
//
//   variantChat := bedrock.PromptVariant_Chat(&ChatPromptVariantProps{
//   	VariantName: jsii.String("variant1"),
//   	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
//   	Messages: []chatMessage{
//   		bedrock.*chatMessage_User(jsii.String("From now on, you speak Japanese!")),
//   		bedrock.*chatMessage_Assistant(jsii.String("Konnichiwa!")),
//   		bedrock.*chatMessage_*User(jsii.String("From now on, you speak {{language}}!")),
//   	},
//   	System: jsii.String("You are a helpful assistant that only speaks the language you`re told."),
//   	PromptVariables: []*string{
//   		jsii.String("language"),
//   	},
//   	ToolConfiguration: &ToolConfiguration{
//   		ToolChoice: bedrock.ToolChoice_AUTO(),
//   		Tools: []tool{
//   			bedrock.*tool_Function(&FunctionToolProps{
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
//   	Variants: []iPromptVariant{
//   		variantChat,
//   	},
//   	KmsKey: cmk,
//   })
//
// Experimental.
type ToolConfiguration struct {
	// How the model should choose which tool to use.
	// Experimental.
	ToolChoice ToolChoice `field:"required" json:"toolChoice" yaml:"toolChoice"`
	// The tools available to the model.
	// Experimental.
	Tools *[]Tool `field:"required" json:"tools" yaml:"tools"`
}

