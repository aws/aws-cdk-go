package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// Configuration for overriding model and prompt template.
//
// Example:
//   // Create a custom semantic memory strategy
//   customSemanticStrategy := agentcore.MemoryStrategy_UsingSemantic(&ManagedStrategyProps{
//   	StrategyName: jsii.String("customSemanticStrategy"),
//   	Description: jsii.String("Custom semantic memory strategy"),
//   	Namespaces: []*string{
//   		jsii.String("/custom/strategies/{memoryStrategyId}/actors/{actorId}"),
//   	},
//   	CustomConsolidation: &OverrideConfig{
//   		Model: bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("ConsolidationModel"), bedrock.FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_5_SONNET_20241022_V2_0()),
//   		AppendToPrompt: jsii.String("Custom consolidation prompt for semantic memory"),
//   	},
//   	CustomExtraction: &OverrideConfig{
//   		Model: bedrock.FoundationModel_*FromFoundationModelId(this, jsii.String("ExtractionModel"), bedrock.FoundationModelIdentifier_ANTHROPIC_CLAUDE_3_5_SONNET_20241022_V2_0()),
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
type OverrideConfig struct {
	// The prompt that will be appended to the system prompt to define the model's memory consolidation/extraction strategy.
	//
	// This configuration provides customization to how the model identifies and extracts
	// relevant information for memory storage. You can use the default user prompt or create a customized one.
	// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/system-prompts.html
	//
	AppendToPrompt *string `field:"required" json:"appendToPrompt" yaml:"appendToPrompt"`
	// The model to use for consolidation/extraction.
	Model awsbedrock.IModel `field:"required" json:"model" yaml:"model"`
}

