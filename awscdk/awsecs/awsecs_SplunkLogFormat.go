package awsecs


// Log Message Format.
// Experimental.
type SplunkLogFormat string

const (
	// Experimental.
	SplunkLogFormat_INLINE SplunkLogFormat = "INLINE"
	// Experimental.
	SplunkLogFormat_JSON SplunkLogFormat = "JSON"
	// Experimental.
	SplunkLogFormat_RAW SplunkLogFormat = "RAW"
)

