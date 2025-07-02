package awsbedrockalpha


// The data type for the vectors when using a model to convert text into vector embeddings.
//
// The model must support the specified data type for vector embeddings. Floating-point (float32)
// is the default data type, and is supported by most models for vector embeddings. See Supported
// embeddings models for information on the available models and their vector data types.
// Experimental.
type VectorType string

const (
	// `FLOATING_POINT` convert the data to floating-point (float32) vector embeddings (more precise, but more costly).
	// Experimental.
	VectorType_FLOATING_POINT VectorType = "FLOATING_POINT"
	// `BINARY` convert the data to binary vector embeddings (less precise, but less costly).
	// Experimental.
	VectorType_BINARY VectorType = "BINARY"
)

