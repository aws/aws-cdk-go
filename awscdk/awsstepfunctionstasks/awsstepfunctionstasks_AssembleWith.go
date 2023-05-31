package awsstepfunctionstasks


// How to assemble the results of the transform job as a single S3 object.
// Experimental.
type AssembleWith string

const (
	// Concatenate the results in binary format.
	// Experimental.
	AssembleWith_NONE AssembleWith = "NONE"
	// Add a newline character at the end of every transformed record.
	// Experimental.
	AssembleWith_LINE AssembleWith = "LINE"
)

