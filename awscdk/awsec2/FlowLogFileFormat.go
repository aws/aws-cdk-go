package awsec2


// The file format for flow logs written to an S3 bucket destination.
type FlowLogFileFormat string

const (
	// File will be written as plain text.
	//
	// This is the default value.
	FlowLogFileFormat_PLAIN_TEXT FlowLogFileFormat = "PLAIN_TEXT"
	// File will be written in parquet format.
	FlowLogFileFormat_PARQUET FlowLogFileFormat = "PARQUET"
)

