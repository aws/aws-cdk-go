package awscdk


// Type hints for resolved values.
type ResolutionTypeHint string

const (
	// This value is expected to resolve to a String.
	ResolutionTypeHint_STRING ResolutionTypeHint = "STRING"
	// This value is expected to resolve to a Number.
	ResolutionTypeHint_NUMBER ResolutionTypeHint = "NUMBER"
	// This value is expected to resolve to a String List.
	ResolutionTypeHint_STRING_LIST ResolutionTypeHint = "STRING_LIST"
)

