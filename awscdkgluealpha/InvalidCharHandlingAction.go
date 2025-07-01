package awscdkgluealpha


// Specifies the action to perform when query results contain invalid UTF-8 character values.
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"invalid_char_handling"_
//
// Experimental.
type InvalidCharHandlingAction string

const (
	// Doesn't perform invalid character handling.
	// Experimental.
	InvalidCharHandlingAction_DISABLED InvalidCharHandlingAction = "DISABLED"
	// Cancels queries that return data containing invalid UTF-8 values.
	// Experimental.
	InvalidCharHandlingAction_FAIL InvalidCharHandlingAction = "FAIL"
	// Replaces invalid UTF-8 values with null.
	// Experimental.
	InvalidCharHandlingAction_SET_TO_NULL InvalidCharHandlingAction = "SET_TO_NULL"
	// Replaces each value in the row with null.
	// Experimental.
	InvalidCharHandlingAction_DROP_ROW InvalidCharHandlingAction = "DROP_ROW"
	// Replaces the invalid character with the replacement character you specify using `REPLACEMENT_CHAR`.
	// Experimental.
	InvalidCharHandlingAction_REPLACE InvalidCharHandlingAction = "REPLACE"
)

