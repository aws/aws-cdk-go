package awsbedrockagentcorealpha

import (
	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2"
)

// Configuration for overriding model and prompt template.
//
// Example:
//   // Create a custom semantic memory strategy
//   customSemanticStrategy := agentcore.MemoryStrategy_UsingSemantic(&ManagedStrategyProps{
//   	Name: jsii.String("customSemanticStrategy"),
//   	Description: jsii.String("Custom semantic memory strategy"),
//   	Namespaces: []*string{
//   		jsii.String("/custom/strategies/{memoryStrategyId}/actors/{actorId}"),
//   	},
//   	CustomConsolidation: &OverrideConfig{
//   		Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
//   		AppendToPrompt: jsii.String("Custom consolidation prompt for semantic memory"),
//   	},
//   	CustomExtraction: &OverrideConfig{
//   		Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
//   		AppendToPrompt: jsii.String("Custom extraction prompt for semantic memory"),
//   	},
//   })
//
//   // Create memory with custom strategy
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my-custom-memory"),
//   	Description: jsii.String("Memory with custom strategy"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	MemoryStrategies: []IMemoryStrategy{
//   		customSemanticStrategy,
//   	},
//   })
//
// Experimental.
type OverrideConfig struct {
	// The prompt that will be appended to the system prompt to define the model's memory consolidation/extraction strategy.
	//
	// This configuration provides customization to how the model identifies and extracts
	// relevant information for memory storage. You can use the default user prompt or create a customized one.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/system-prompts.html
	//
	// Experimental.
	AppendToPrompt *string `field:"required" json:"appendToPrompt" yaml:"appendToPrompt"`
	// The model to use for consolidation/extraction.
	// Experimental.
	Model awsbedrockalpha.IBedrockInvokable `field:"required" json:"model" yaml:"model"`
}

