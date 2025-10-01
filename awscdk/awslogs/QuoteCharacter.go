package awslogs


// Valid quote characters for CSV processor.
//
// Defines the character used as a text qualifier for a single column of data.
type QuoteCharacter string

const (
	// Double quote character (default).
	QuoteCharacter_DOUBLE_QUOTE QuoteCharacter = "DOUBLE_QUOTE"
	// Single quote character.
	QuoteCharacter_SINGLE_QUOTE QuoteCharacter = "SINGLE_QUOTE"
)

