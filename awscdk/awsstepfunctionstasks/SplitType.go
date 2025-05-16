package awsstepfunctionstasks


// Method to use to split the transform job's data files into smaller batches.
type SplitType string

const (
	// Input data files are not split,.
	SplitType_NONE SplitType = "NONE"
	// Split records on a newline character boundary.
	SplitType_LINE SplitType = "LINE"
	// Split using MXNet RecordIO format.
	SplitType_RECORD_IO SplitType = "RECORD_IO"
	// Split using TensorFlow TFRecord format.
	SplitType_TF_RECORD SplitType = "TF_RECORD"
)

