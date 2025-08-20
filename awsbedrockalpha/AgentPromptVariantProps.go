package awsbedrockalpha


// Properties for creating an agent prompt variant.
//
// Example:
//   cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
//   })
//
//   // Assuming you have an existing agent and alias
//   agent := bedrock.Agent_FromAgentAttributes(this, jsii.String("ImportedAgent"), &AgentAttributes{
//   	AgentArn: jsii.String("arn:aws:bedrock:region:account:agent/agent-id"),
//   	RoleArn: jsii.String("arn:aws:iam::account:role/agent-role"),
//   })
//
//   agentAlias := bedrock.AgentAlias_FromAttributes(this, jsii.String("ImportedAlias"), &AgentAliasAttributes{
//   	AliasId: jsii.String("alias-id"),
//   	AliasName: jsii.String("my-alias"),
//   	AgentVersion: jsii.String("1"),
//   	Agent: agent,
//   })
//
//   agentVariant := bedrock.PromptVariant_Agent(&AgentPromptVariantProps{
//   	VariantName: jsii.String("agent-variant"),
//   	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
//   	AgentAlias: agentAlias,
//   	PromptText: jsii.String("Use the agent to help with: {{task}}. Please be thorough and provide detailed explanations."),
//   	PromptVariables: []*string{
//   		jsii.String("task"),
//   	},
//   })
//
//   bedrock.NewPrompt(this, jsii.String("agentPrompt"), &PromptProps{
//   	PromptName: jsii.String("agent-prompt"),
//   	Description: jsii.String("Prompt for agent interactions"),
//   	DefaultVariant: agentVariant,
//   	Variants: []iPromptVariant{
//   		agentVariant,
//   	},
//   	KmsKey: cmk,
//   })
//
// Experimental.
type AgentPromptVariantProps struct {
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
	// An alias pointing to the agent version to be used.
	// Experimental.
	AgentAlias IAgentAlias `field:"required" json:"agentAlias" yaml:"agentAlias"`
	// The text prompt.
	//
	// Variables are used by enclosing its name with double curly braces
	// as in `{{variable_name}}`.
	// Experimental.
	PromptText *string `field:"required" json:"promptText" yaml:"promptText"`
}

