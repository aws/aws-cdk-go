package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating memory strategies If you need long-term memory for context recall across sessions, you can setup memory extraction strategies to extract the relevant memory from the raw events.
//
// Use built-in strategies for quick setup, use built-in strategies with override to specify models and prompt templates.
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
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/memory-strategies.html
//
type MemoryStrategy interface {
}

// The jsii proxy struct for MemoryStrategy
type jsiiProxy_MemoryStrategy struct {
	_ byte // padding
}

func NewMemoryStrategy() MemoryStrategy {
	_init_.Initialize()

	j := jsiiProxy_MemoryStrategy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewMemoryStrategy_Override(m MemoryStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		nil, // no parameters
		m,
	)
}

// Default strategies for organizing and extracting memory data, each optimized for specific use cases.
//
// Captures meaningful slices of user and system interactions, preserve them into compact records after summarizing.
// Extracted memory example: User first asked about pricing on Monday, then requested feature comparison on Tuesday, finally made purchase decision on Wednesday.
//
// Returns: A ManagedMemoryStrategy.
func MemoryStrategy_UsingBuiltInEpisodic() ManagedMemoryStrategy {
	_init_.Initialize()

	var returns ManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingBuiltInEpisodic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Default strategies for organizing and extracting memory data, each optimized for specific use cases.
//
// Distills general facts, concepts, and underlying meanings from raw conversational data, presenting the information in a context-independent format.
// Extracted memory example: In-context learning = task-solving via examples, no training needed.
//
// Returns: A ManagedMemoryStrategy.
func MemoryStrategy_UsingBuiltInSemantic() ManagedMemoryStrategy {
	_init_.Initialize()

	var returns ManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingBuiltInSemantic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Default strategies for organizing and extracting memory data, each optimized for specific use cases.
//
// This strategy compresses conversations into concise overviews, preserving essential context and key insights for quick recall.
// Extracted memory example: Users confused by cloud setup during onboarding.
//
// Returns: A ManagedMemoryStrategy.
func MemoryStrategy_UsingBuiltInSummarization() ManagedMemoryStrategy {
	_init_.Initialize()

	var returns ManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingBuiltInSummarization",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Default strategies for organizing and extracting memory data, each optimized for specific use cases.
//
// Captures individual preferences, interaction patterns, and personalized settings to enhance future experiences.
// Extracted memory example: User needs clear guidance on cloud storage account connection during onboarding.
//
// Returns: A ManagedMemoryStrategy.
func MemoryStrategy_UsingBuiltInUserPreference() ManagedMemoryStrategy {
	_init_.Initialize()

	var returns ManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingBuiltInUserPreference",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creates an episodic memory strategy with custom configuration.
//
// Captures meaningful slices of user and system interactions, preserve them into compact records after summarizing.
// Extracted memory example: User first asked about pricing on Monday, then requested feature comparison on Tuesday, finally made purchase decision on Wednesday.
//
// Returns: A ManagedMemoryStrategy.
func MemoryStrategy_UsingEpisodic(config *ManagedStrategyProps) ManagedMemoryStrategy {
	_init_.Initialize()

	if err := validateMemoryStrategy_UsingEpisodicParameters(config); err != nil {
		panic(err)
	}
	var returns ManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingEpisodic",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Creates a self-managed memory strategy.
//
// A self-managed strategy gives you complete control over your memory extraction and consolidation pipelines.
//
// Returns: A SelfManagedMemoryStrategy.
func MemoryStrategy_UsingSelfManaged(config *SelfManagedStrategyProps) SelfManagedMemoryStrategy {
	_init_.Initialize()

	if err := validateMemoryStrategy_UsingSelfManagedParameters(config); err != nil {
		panic(err)
	}
	var returns SelfManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingSelfManaged",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Creates a semantic memory strategy with custom configuration.
//
// Distills general facts, concepts, and underlying meanings from raw conversational data, presenting the information in a context-independent format.
// Extracted memory example: In-context learning = task-solving via examples, no training needed.
//
// Returns: A ManagedMemoryStrategy.
func MemoryStrategy_UsingSemantic(config *ManagedStrategyProps) ManagedMemoryStrategy {
	_init_.Initialize()

	if err := validateMemoryStrategy_UsingSemanticParameters(config); err != nil {
		panic(err)
	}
	var returns ManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingSemantic",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Creates a summarization memory strategy with custom configuration.
//
// This strategy compresses conversations into concise overviews, preserving essential context and key insights for quick recall.
// Extracted memory example: Users confused by cloud setup during onboarding.
//
// Returns: A ManagedMemoryStrategy.
func MemoryStrategy_UsingSummarization(config *ManagedStrategyProps) ManagedMemoryStrategy {
	_init_.Initialize()

	if err := validateMemoryStrategy_UsingSummarizationParameters(config); err != nil {
		panic(err)
	}
	var returns ManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingSummarization",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Creates a user preference memory strategy with custom configuration.
//
// Captures individual preferences, interaction patterns, and personalized settings to enhance future experiences.
// Extracted memory example: User needs clear guidance on cloud storage account connection during onboarding.
//
// Returns: A ManagedMemoryStrategy.
func MemoryStrategy_UsingUserPreference(config *ManagedStrategyProps) ManagedMemoryStrategy {
	_init_.Initialize()

	if err := validateMemoryStrategy_UsingUserPreferenceParameters(config); err != nil {
		panic(err)
	}
	var returns ManagedMemoryStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.MemoryStrategy",
		"usingUserPreference",
		[]interface{}{config},
		&returns,
	)

	return returns
}

