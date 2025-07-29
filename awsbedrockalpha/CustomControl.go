package awsbedrockalpha


// The type of custom control for the action group executor.
// Experimental.
type CustomControl string

const (
	// Returns the action group invocation results directly in the InvokeAgent response.
	// Experimental.
	CustomControl_RETURN_CONTROL CustomControl = "RETURN_CONTROL"
)

