package awsstepfunctionstasks


// Specifies the number of records to include in a mini-batch for an HTTP inference request.
// Experimental.
type BatchStrategy string

const (
	// Fits multiple records in a mini-batch.
	// Experimental.
	BatchStrategy_MULTI_RECORD BatchStrategy = "MULTI_RECORD"
	// Use a single record when making an invocation request.
	// Experimental.
	BatchStrategy_SINGLE_RECORD BatchStrategy = "SINGLE_RECORD"
)

