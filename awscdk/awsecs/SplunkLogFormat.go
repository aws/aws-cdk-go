package awsecs


// Log Message Format.
type SplunkLogFormat string

const (
	SplunkLogFormat_INLINE SplunkLogFormat = "INLINE"
	SplunkLogFormat_JSON SplunkLogFormat = "JSON"
	SplunkLogFormat_RAW SplunkLogFormat = "RAW"
)

