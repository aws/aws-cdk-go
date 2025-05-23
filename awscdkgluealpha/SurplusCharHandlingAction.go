package awscdkgluealpha


// Specifies how to handle data being loaded that exceeds the length of the data type defined for columns containing VARCHAR, CHAR, or string data.
//
// By default, Redshift Spectrum sets the value to null for data that exceeds the width of the column.
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"surplus_char_handling"_
//
// Experimental.
type SurplusCharHandlingAction string

const (
	// Replaces data that exceeds the column width with null.
	// Experimental.
	SurplusCharHandlingAction_SET_TO_NULL SurplusCharHandlingAction = "SET_TO_NULL"
	// Doesn't perform surplus character handling.
	// Experimental.
	SurplusCharHandlingAction_DISABLED SurplusCharHandlingAction = "DISABLED"
	// Cancels queries that return data exceeding the column width.
	// Experimental.
	SurplusCharHandlingAction_FAIL SurplusCharHandlingAction = "FAIL"
	// Replaces each value in the row with null.
	// Experimental.
	SurplusCharHandlingAction_DROP_ROW SurplusCharHandlingAction = "DROP_ROW"
	// Removes the characters that exceed the maximum number of characters defined for the column.
	// Experimental.
	SurplusCharHandlingAction_TRUNCATE SurplusCharHandlingAction = "TRUNCATE"
)

