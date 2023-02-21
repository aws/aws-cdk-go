package awsstepfunctionstasks


// Define the format of the input data.
type RecordWrapperType string

const (
	// None record wrapper type.
	RecordWrapperType_NONE RecordWrapperType = "NONE"
	// RecordIO record wrapper type.
	RecordWrapperType_RECORD_IO RecordWrapperType = "RECORD_IO"
)

