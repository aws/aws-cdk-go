package awsbedrockalpha


// The role of a message in a chat conversation.
// Experimental.
type ChatMessageRole string

const (
	// This role represents the human user in the conversation.
	//
	// Inputs from the
	// user guide the conversation and prompt responses from the assistant.
	// Experimental.
	ChatMessageRole_USER ChatMessageRole = "USER"
	// This is the role of the model itself, responding to user inputs based on the context set by the system.
	// Experimental.
	ChatMessageRole_ASSISTANT ChatMessageRole = "ASSISTANT"
)

