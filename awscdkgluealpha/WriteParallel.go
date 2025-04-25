package awscdkgluealpha


// Specifies how to handle data being loaded that exceeds the length of the data type defined for columns containing VARCHAR, CHAR, or string data.
//
// By default, Redshift Spectrum sets the value to null for data that exceeds the width of the column.
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"surplus_char_handling"_
//
// Experimental.
type WriteParallel string

const (
	// Write data in parallel.
	// Experimental.
	WriteParallel_ON WriteParallel = "ON"
	// Write data serially.
	// Experimental.
	WriteParallel_OFF WriteParallel = "OFF"
)

