package awsappsync


// Enumerated type for the handler behavior for a channel namespace.
type HandlerBehavior string

const (
	// Code handler.
	HandlerBehavior_CODE HandlerBehavior = "CODE"
	// Direct integration handler.
	HandlerBehavior_DIRECT HandlerBehavior = "DIRECT"
)

