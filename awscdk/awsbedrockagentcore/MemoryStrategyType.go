package awsbedrockagentcore


// Long-term memory extraction strategy types.
type MemoryStrategyType string

const (
	// Summarization strategy - extracts concise summaries to preserve critical context and key insights.
	MemoryStrategyType_SUMMARIZATION MemoryStrategyType = "SUMMARIZATION"
	// Semantic memory strategy - extracts general factual knowledge, concepts and meanings from raw conversations using vector embeddings for similarity search.
	MemoryStrategyType_SEMANTIC MemoryStrategyType = "SEMANTIC"
	// User preferences strategy - extracts user behavior patterns from raw conversations.
	MemoryStrategyType_USER_PREFERENCE MemoryStrategyType = "USER_PREFERENCE"
	// Episodic memory strategy - captures meaningful slices of user and system interactions.
	MemoryStrategyType_EPISODIC MemoryStrategyType = "EPISODIC"
	// Customize memory processing through custom foundation model and prompt templates.
	MemoryStrategyType_CUSTOM MemoryStrategyType = "CUSTOM"
)

