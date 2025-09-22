package awscdkgluealpha


// Identifies if the file contains less or more values for a row than the number of columns specified in the external table definition.
//
// This property is only available for an uncompressed text file format.
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"column_count_mismatch_handling"_
//
// Experimental.
type ColumnCountMismatchHandlingAction string

const (
	// Column count mismatch handling is turned off.
	// Experimental.
	ColumnCountMismatchHandlingAction_DISABLED ColumnCountMismatchHandlingAction = "DISABLED"
	// Fail the query if the column count mismatch is detected.
	// Experimental.
	ColumnCountMismatchHandlingAction_FAIL ColumnCountMismatchHandlingAction = "FAIL"
	// Fill missing values with NULL and ignore the additional values in each row.
	// Experimental.
	ColumnCountMismatchHandlingAction_SET_TO_NULL ColumnCountMismatchHandlingAction = "SET_TO_NULL"
	// Drop all rows that contain column count mismatch error from the scan.
	// Experimental.
	ColumnCountMismatchHandlingAction_DROP_ROW ColumnCountMismatchHandlingAction = "DROP_ROW"
)

