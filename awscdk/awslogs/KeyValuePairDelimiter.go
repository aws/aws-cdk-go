package awslogs


// Valid field delimiters for ParseKeyValue processor.
//
// Defines the delimiter string used between key-value pairs in the original log events.
type KeyValuePairDelimiter string

const (
	// Ampersand character (default).
	KeyValuePairDelimiter_AMPERSAND KeyValuePairDelimiter = "AMPERSAND"
	// Semicolon character.
	KeyValuePairDelimiter_SEMICOLON KeyValuePairDelimiter = "SEMICOLON"
	// Space character.
	KeyValuePairDelimiter_SPACE KeyValuePairDelimiter = "SPACE"
	// Newline character.
	KeyValuePairDelimiter_NEWLINE KeyValuePairDelimiter = "NEWLINE"
)

