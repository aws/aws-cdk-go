package awslogs


// Valid key-value delimiters for ParseKeyValue processor.
//
// Defines the delimiter string to use between the key and value in each pair.
type KeyValueDelimiter string

const (
	// Equal sign (default).
	KeyValueDelimiter_EQUAL KeyValueDelimiter = "EQUAL"
	// Colon character.
	KeyValueDelimiter_COLON KeyValueDelimiter = "COLON"
)

