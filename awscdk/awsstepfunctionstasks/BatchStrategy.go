package awsstepfunctionstasks


// Specifies the number of records to include in a mini-batch for an HTTP inference request.
type BatchStrategy string

const (
	// Fits multiple records in a mini-batch.
	BatchStrategy_MULTI_RECORD BatchStrategy = "MULTI_RECORD"
	// Use a single record when making an invocation request.
	BatchStrategy_SINGLE_RECORD BatchStrategy = "SINGLE_RECORD"
)

