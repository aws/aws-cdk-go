package awsbedrockalpha


// The type of prompt template.
// Experimental.
type PromptTemplateType string

const (
	// Text template for simple text-based prompts.
	// Experimental.
	PromptTemplateType_TEXT PromptTemplateType = "TEXT"
	// Chat template for conversational prompts with message history.
	// Experimental.
	PromptTemplateType_CHAT PromptTemplateType = "CHAT"
)

