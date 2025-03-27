package awscdkgluealpha


// Specifies the action to perform when ORC data contains an integer (for example, BIGINT or int64) that is larger than the column definition (for example, SMALLINT or int16).
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"numeric_overflow_handling"_
//
// Experimental.
type NumericOverflowHandlingAction string

const (
	// Invalid character handling is turned off.
	// Experimental.
	NumericOverflowHandlingAction_DISABLED NumericOverflowHandlingAction = "DISABLED"
	// Cancel the query when the data includes invalid characters.
	// Experimental.
	NumericOverflowHandlingAction_FAIL NumericOverflowHandlingAction = "FAIL"
	// Set invalid characters to null.
	// Experimental.
	NumericOverflowHandlingAction_SET_TO_NULL NumericOverflowHandlingAction = "SET_TO_NULL"
	// Set each value in the row to null.
	// Experimental.
	NumericOverflowHandlingAction_DROP_ROW NumericOverflowHandlingAction = "DROP_ROW"
)

