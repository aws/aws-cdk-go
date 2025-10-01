package awslogs


// Valid delimiter characters for CSV processor.
//
// Defines the character used to separate each column in CSV data.
type DelimiterCharacter string

const (
	// Comma character.
	DelimiterCharacter_COMMA DelimiterCharacter = "COMMA"
	// Tab character.
	DelimiterCharacter_TAB DelimiterCharacter = "TAB"
	// Space character.
	DelimiterCharacter_SPACE DelimiterCharacter = "SPACE"
	// Semicolon character.
	DelimiterCharacter_SEMICOLON DelimiterCharacter = "SEMICOLON"
	// Pipe character.
	DelimiterCharacter_PIPE DelimiterCharacter = "PIPE"
)

