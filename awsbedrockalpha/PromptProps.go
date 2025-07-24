package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a CDK managed Bedrock Prompt.
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

