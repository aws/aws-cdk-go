package awscdkintegtestsalpha


// Set to Tail to include the execution log in the response.
//
// Applies to synchronously invoked functions only.
// Experimental.
type LogType string

const (
	// The log messages are not returned in the response.
	// Experimental.
	LogType_NONE LogType = "NONE"
	// The log messages are returned in the response.
	// Experimental.
	LogType_TAIL LogType = "TAIL"
)

