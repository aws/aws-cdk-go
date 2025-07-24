package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating prompt variants.
//
// Provides static methods to create different types of prompt variants
// with proper configuration and type safety.
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
type PromptVariant interface {
}

// The jsii proxy struct for PromptVariant
type jsiiProxy_PromptVariant struct {
	_ byte // padding
}

// Creates an agent prompt template variant.
//
// Returns: A PromptVariant configured for agent interactions.
// Experimental.
func PromptVariant_Agent(props *AgentPromptVariantProps) IPromptVariant {
	_init_.Initialize()

	if err := validatePromptVariant_AgentParameters(props); err != nil {
		panic(err)
	}
	var returns IPromptVariant

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptVariant",
		"agent",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a chat template variant.
//
// Use this template type when
// the model supports the Converse API or the Anthropic Claude Messages API.
// This allows you to include a System prompt and previous User messages
// and Assistant messages for context.
//
// Returns: A PromptVariant configured for chat interactions.
// Experimental.
func PromptVariant_Chat(props *ChatPromptVariantProps) IPromptVariant {
	_init_.Initialize()

	if err := validatePromptVariant_ChatParameters(props); err != nil {
		panic(err)
	}
	var returns IPromptVariant

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptVariant",
		"chat",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a text template variant.
//
// Returns: A PromptVariant configured for text processing.
// Experimental.
func PromptVariant_Text(props *TextPromptVariantProps) IPromptVariant {
	_init_.Initialize()

	if err := validatePromptVariant_TextParameters(props); err != nil {
		panic(err)
	}
	var returns IPromptVariant

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptVariant",
		"text",
		[]interface{}{props},
		&returns,
	)

	return returns
}

