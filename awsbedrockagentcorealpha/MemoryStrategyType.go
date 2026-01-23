package awsbedrockagentcorealpha


// Long-term memory extraction strategy types.
// Experimental.
type MemoryStrategyType string

const (
	// Summarization strategy - extracts concise summaries to preserve critical context and key insights.
	// Experimental.
	MemoryStrategyType_SUMMARIZATION MemoryStrategyType = "SUMMARIZATION"
	// Semantic memory strategy - extracts general factual knowledge, concepts and meanings from raw conversations using vector embeddings for similarity search.
	// Experimental.
	MemoryStrategyType_SEMANTIC MemoryStrategyType = "SEMANTIC"
	// User preferences strategy - extracts user behavior patterns from raw conversations.
	// Experimental.
	MemoryStrategyType_USER_PREFERENCE MemoryStrategyType = "USER_PREFERENCE"
	// Episodic memory strategy - captures meaningful slices of user and system interactions.
	// Experimental.
	MemoryStrategyType_EPISODIC MemoryStrategyType = "EPISODIC"
	// Customize memory processing through custom foundation model and prompt templates.
	// Experimental.
	MemoryStrategyType_CUSTOM MemoryStrategyType = "CUSTOM"
)

