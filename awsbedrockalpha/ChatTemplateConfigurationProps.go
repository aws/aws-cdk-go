package awsbedrockalpha


// Properties for creating a chat template configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   var chatMessage chatMessage
//   var tool tool
//   var toolChoice toolChoice
//
//   chatTemplateConfigurationProps := &ChatTemplateConfigurationProps{
//   	Messages: []*chatMessage{
//   		chatMessage,
//   	},
//
//   	// the properties below are optional
//   	InputVariables: []*string{
//   		jsii.String("inputVariables"),
//   	},
//   	System: jsii.String("system"),
//   	ToolConfiguration: &ToolConfiguration{
//   		ToolChoice: toolChoice,
//   		Tools: []*tool{
//   			tool,
//   		},
//   	},
//   }
//
// Experimental.
type ChatTemplateConfigurationProps struct {
	// The messages in the chat template.
	// Experimental.
	Messages *[]ChatMessage `field:"required" json:"messages" yaml:"messages"`
	// The input variables for the template.
	// Default: - No input variables.
	//
	// Experimental.
	InputVariables *[]*string `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// The system message for the chat template.
	// Default: - No system message.
	//
	// Experimental.
	System *string `field:"optional" json:"system" yaml:"system"`
	// The tool configuration for the chat template.
	// Default: - No tool configuration.
	//
	// Experimental.
	ToolConfiguration *ToolConfiguration `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

