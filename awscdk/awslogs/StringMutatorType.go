package awslogs


// Types of string mutation operations.
//
// Defines various operations that can be performed to modify string values in log events.
type StringMutatorType string

const (
	// Convert strings to lowercase.
	StringMutatorType_LOWER_CASE StringMutatorType = "LOWER_CASE"
	// Convert strings to uppercase.
	StringMutatorType_UPPER_CASE StringMutatorType = "UPPER_CASE"
	// Trim whitespace from strings.
	StringMutatorType_TRIM StringMutatorType = "TRIM"
	// Split strings by delimiter.
	StringMutatorType_SPLIT StringMutatorType = "SPLIT"
	// Replace substrings in strings.
	StringMutatorType_SUBSTITUTE StringMutatorType = "SUBSTITUTE"
)

