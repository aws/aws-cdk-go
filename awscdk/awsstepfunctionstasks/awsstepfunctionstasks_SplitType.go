package awsstepfunctionstasks


// Method to use to split the transform job's data files into smaller batches.
// Experimental.
type SplitType string

const (
	// Input data files are not split,.
	// Experimental.
	SplitType_NONE SplitType = "NONE"
	// Split records on a newline character boundary.
	// Experimental.
	SplitType_LINE SplitType = "LINE"
	// Split using MXNet RecordIO format.
	// Experimental.
	SplitType_RECORD_IO SplitType = "RECORD_IO"
	// Split using TensorFlow TFRecord format.
	// Experimental.
	SplitType_TF_RECORD SplitType = "TF_RECORD"
)

