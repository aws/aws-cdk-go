package awsbedrockalpha


// Attributes needed to create an import.
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
type AgentAliasAttributes struct {
	// The underlying agent for this alias.
	// Experimental.
	Agent IAgent `field:"required" json:"agent" yaml:"agent"`
	// The agent version for this alias.
	// Experimental.
	AgentVersion *string `field:"required" json:"agentVersion" yaml:"agentVersion"`
	// The Id of the agent alias.
	// Experimental.
	AliasId *string `field:"required" json:"aliasId" yaml:"aliasId"`
	// The name of the agent alias.
	// Default: undefined - No alias name is provided.
	//
	// Experimental.
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
}

