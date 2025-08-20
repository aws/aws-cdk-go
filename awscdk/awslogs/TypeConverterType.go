package awslogs


// Valid data types for type conversion in the TypeConverter processor.
//
// Used to specify the target data type for field conversion.
type TypeConverterType string

const (
	// Convert value to boolean type.
	TypeConverterType_BOOLEAN TypeConverterType = "BOOLEAN"
	// Convert value to integer type.
	TypeConverterType_INTEGER TypeConverterType = "INTEGER"
	// Convert value to double (floating point) type.
	TypeConverterType_DOUBLE TypeConverterType = "DOUBLE"
	// Convert value to string type.
	TypeConverterType_STRING TypeConverterType = "STRING"
)

