package awslogs


// Types of data conversion operations.
//
// Defines operations that can convert data from one format to another.
type DataConverterType string

const (
	// Convert data types.
	DataConverterType_TYPE_CONVERTER DataConverterType = "TYPE_CONVERTER"
	// Convert datetime formats.
	DataConverterType_DATETIME_CONVERTER DataConverterType = "DATETIME_CONVERTER"
)

