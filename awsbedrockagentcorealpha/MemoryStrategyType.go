package awsbedrockagentcorealpha


// Long-term memory extraction strategy types.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type MemoryStrategyType string

const (
	// Summarization strategy - extracts concise summaries to preserve critical context and key insights.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MemoryStrategyType_SUMMARIZATION MemoryStrategyType = "SUMMARIZATION"
	// Semantic memory strategy - extracts general factual knowledge, concepts and meanings from raw conversations using vector embeddings for similarity search.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MemoryStrategyType_SEMANTIC MemoryStrategyType = "SEMANTIC"
	// User preferences strategy - extracts user behavior patterns from raw conversations.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MemoryStrategyType_USER_PREFERENCE MemoryStrategyType = "USER_PREFERENCE"
	// Episodic memory strategy - captures meaningful slices of user and system interactions.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MemoryStrategyType_EPISODIC MemoryStrategyType = "EPISODIC"
	// Customize memory processing through custom foundation model and prompt templates.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	MemoryStrategyType_CUSTOM MemoryStrategyType = "CUSTOM"
)

