package awsstepfunctionstasks


// How to assemble the results of the transform job as a single S3 object.
type AssembleWith string

const (
	// Concatenate the results in binary format.
	AssembleWith_NONE AssembleWith = "NONE"
	// Add a newline character at the end of every transformed record.
	AssembleWith_LINE AssembleWith = "LINE"
)

