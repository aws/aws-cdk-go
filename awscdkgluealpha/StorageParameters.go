package awscdkgluealpha


// The storage parameter keys that are currently known, this list is not exhaustive and other keys may be used.
//
// Example:
//   var glueDatabase IDatabase
//
//   table := glue.NewTable(this, jsii.String("Table"), &S3TableProps{
//   	StorageParameters: []StorageParameter{
//   		glue.StorageParameter_SkipHeaderLineCount(jsii.Number(1)),
//   		glue.StorageParameter_CompressionType(glue.CompressionType_GZIP),
//   		glue.StorageParameter_Custom(jsii.String("foo"), jsii.String("bar")),
//   		glue.StorageParameter_*Custom(jsii.String("separatorChar"), jsii.String(",")),
//   		glue.StorageParameter_*Custom(glue.StorageParameters_WRITE_PARALLEL, jsii.String("off")),
//   	},
//   	// ...
//   	Database: glueDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_CSV(),
//   })
//
// Experimental.
type StorageParameters string

const (
	// The number of rows to skip at the top of a CSV file when the table is being created.
	// Experimental.
	StorageParameters_SKIP_HEADER_LINE_COUNT StorageParameters = "SKIP_HEADER_LINE_COUNT"
	// Determines whether data handling is on for the table.
	// Experimental.
	StorageParameters_DATA_CLEANSING_ENABLED StorageParameters = "DATA_CLEANSING_ENABLED"
	// The type of compression used on the table, when the file name does not contain an extension.
	//
	// This value overrides the compression type specified through the extension.
	// Experimental.
	StorageParameters_COMPRESSION_TYPE StorageParameters = "COMPRESSION_TYPE"
	// Specifies the action to perform when query results contain invalid UTF-8 character values.
	// Experimental.
	StorageParameters_INVALID_CHAR_HANDLING StorageParameters = "INVALID_CHAR_HANDLING"
	// Specifies the replacement character to use when you set `INVALID_CHAR_HANDLING` to `REPLACE`.
	// Experimental.
	StorageParameters_REPLACEMENT_CHAR StorageParameters = "REPLACEMENT_CHAR"
	// Specifies the action to perform when ORC data contains an integer (for example, BIGINT or int64) that is larger than the column definition (for example, SMALLINT or int16).
	// Experimental.
	StorageParameters_NUMERIC_OVERFLOW_HANDLING StorageParameters = "NUMERIC_OVERFLOW_HANDLING"
	// Specifies how to handle data being loaded that exceeds the length of the data type defined for columns containing VARBYTE data.
	//
	// By default, Redshift Spectrum sets the value to null for data that exceeds the width of the column.
	// Experimental.
	StorageParameters_SURPLUS_BYTES_HANDLING StorageParameters = "SURPLUS_BYTES_HANDLING"
	// Specifies how to handle data being loaded that exceeds the length of the data type defined for columns containing VARCHAR, CHAR, or string data.
	//
	// By default, Redshift Spectrum sets the value to null for data that exceeds the width of the column.
	// Experimental.
	StorageParameters_SURPLUS_CHAR_HANDLING StorageParameters = "SURPLUS_CHAR_HANDLING"
	// Identifies if the file contains less or more values for a row than the number of columns specified in the external table definition.
	//
	// This property is only available for an uncompressed text file format.
	// Experimental.
	StorageParameters_COLUMN_COUNT_MISMATCH_HANDLING StorageParameters = "COLUMN_COUNT_MISMATCH_HANDLING"
	// A property that sets the numRows value for the table definition.
	//
	// To explicitly update an external table's statistics, set the numRows property to indicate the size of the table. Amazon Redshift doesn't analyze external tables to generate the table statistics that the query optimizer uses to generate a query plan. If table statistics aren't set for an external table, Amazon Redshift generates a query execution plan based on an assumption that external tables are the larger tables and local tables are the smaller tables.
	// Experimental.
	StorageParameters_NUM_ROWS StorageParameters = "NUM_ROWS"
	// A property that sets number of rows to skip at the beginning of each source file.
	// Experimental.
	StorageParameters_SERIALIZATION_NULL_FORMAT StorageParameters = "SERIALIZATION_NULL_FORMAT"
	// A property that sets the column mapping type for tables that use ORC data format.
	//
	// This property is ignored for other data formats.
	// Experimental.
	StorageParameters_ORC_SCHEMA_RESOLUTION StorageParameters = "ORC_SCHEMA_RESOLUTION"
	// A property that sets whether CREATE EXTERNAL TABLE AS should write data in parallel.
	//
	// When 'write.parallel' is set to off, CREATE EXTERNAL TABLE AS writes to one or more data files serially onto Amazon S3. This table property also applies to any subsequent INSERT statement into the same external table.
	// Experimental.
	StorageParameters_WRITE_PARALLEL StorageParameters = "WRITE_PARALLEL"
	// A property that sets the maximum size (in MB) of each file written to Amazon S3 by CREATE EXTERNAL TABLE AS.
	//
	// The size must be a valid integer between 5 and 6200. The default maximum file size is 6,200 MB. This table property also applies to any subsequent INSERT statement into the same external table.
	// Experimental.
	StorageParameters_WRITE_MAX_FILESIZE_MB StorageParameters = "WRITE_MAX_FILESIZE_MB"
	// You can specify an AWS Key Management Service key to enable Server–Side Encryption (SSE) for Amazon S3 objects.
	// Experimental.
	StorageParameters_WRITE_KMS_KEY_ID StorageParameters = "WRITE_KMS_KEY_ID"
)

