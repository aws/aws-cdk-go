package awsbedrockalpha


// Attributes for specifying an imported Bedrock Agent.
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
type AgentAttributes struct {
	// The ARN of the agent.
	// Experimental.
	AgentArn *string `field:"required" json:"agentArn" yaml:"agentArn"`
	// The ARN of the IAM role associated to the agent.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The agent version.
	//
	// If no explicit versions have been created,
	// leave this empty to use the DRAFT version. Otherwise, use the
	// version number (e.g. 1).
	// Default: 'DRAFT'.
	//
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Optional KMS encryption key associated with this agent.
	// Default: undefined - An AWS managed key is used.
	//
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// When this agent was last updated.
	// Default: undefined - No last updated timestamp is provided.
	//
	// Experimental.
	LastUpdated *string `field:"optional" json:"lastUpdated" yaml:"lastUpdated"`
}

