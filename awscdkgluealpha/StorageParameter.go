package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A storage parameter. The list of storage parameters available is not exhaustive and other keys may be used.
//
// If you would like to specify a storage parameter that is not available as a static member of this class, use the `StorageParameter.custom` method.
//
// The list of storage parameters currently known within the CDK is listed.
//
// Example:
//   var myDatabase database
//
//   glue.NewTable(this, jsii.String("MyTable"), &TableProps{
//   	StorageParameters: []storageParameter{
//   		glue.*storageParameter_SkipHeaderLineCount(jsii.Number(1)),
//   		glue.*storageParameter_CompressionType(glue.CompressionType_GZIP),
//   		glue.*storageParameter_Custom(jsii.String("separatorChar"), jsii.String(",")),
//   	},
//   	// ...
//   	Database: myDatabase,
//   	Columns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   })
//
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_
//
// Experimental.
type StorageParameter interface {
	// Experimental.
	Key() *string
	// Experimental.
	Value() *string
}

// The jsii proxy struct for StorageParameter
type jsiiProxy_StorageParameter struct {
	_ byte // padding
}

func (j *jsiiProxy_StorageParameter) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageParameter) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewStorageParameter(key *string, value *string) StorageParameter {
	_init_.Initialize()

	if err := validateNewStorageParameterParameters(key, value); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageParameter{}

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		[]interface{}{key, value},
		&j,
	)

	return &j
}

// Experimental.
func NewStorageParameter_Override(s StorageParameter, key *string, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		[]interface{}{key, value},
		s,
	)
}

// Identifies if the file contains less or more values for a row than the number of columns specified in the external table definition.
//
// This property is only available for an uncompressed text file format.
// Experimental.
func StorageParameter_ColumnCountMismatchHandling(value ColumnCountMismatchHandlingAction) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_ColumnCountMismatchHandlingParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"columnCountMismatchHandling",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The type of compression used on the table, when the file name does not contain an extension.
//
// This value overrides the compression type specified through the extension.
// Experimental.
func StorageParameter_CompressionType(value CompressionType) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_CompressionTypeParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"compressionType",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A custom storage parameter.
// Experimental.
func StorageParameter_Custom(key *string, value interface{}) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_CustomParameters(key, value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"custom",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Determines whether data handling is on for the table.
// Experimental.
func StorageParameter_DataCleansingEnabled(value *bool) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_DataCleansingEnabledParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"dataCleansingEnabled",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Specifies the action to perform when query results contain invalid UTF-8 character values.
// Experimental.
func StorageParameter_InvalidCharHandling(value InvalidCharHandlingAction) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_InvalidCharHandlingParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"invalidCharHandling",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Specifies the action to perform when ORC data contains an integer (for example, BIGINT or int64) that is larger than the column definition (for example, SMALLINT or int16).
// Experimental.
func StorageParameter_NumericOverflowHandling(value NumericOverflowHandlingAction) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_NumericOverflowHandlingParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"numericOverflowHandling",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A property that sets the numRows value for the table definition.
//
// To explicitly update an external table's statistics, set the numRows property to indicate the size of the table. Amazon Redshift doesn't analyze external tables to generate the table statistics that the query optimizer uses to generate a query plan. If table statistics aren't set for an external table, Amazon Redshift generates a query execution plan based on an assumption that external tables are the larger tables and local tables are the smaller tables.
// Experimental.
func StorageParameter_NumRows(value *float64) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_NumRowsParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"numRows",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A property that sets the column mapping type for tables that use ORC data format.
//
// This property is ignored for other data formats. If this property is omitted, columns are mapped by `OrcColumnMappingType.NAME` by default.
// Experimental.
func StorageParameter_OrcSchemaResolution(value OrcColumnMappingType) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_OrcSchemaResolutionParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"orcSchemaResolution",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Specifies the replacement character to use when you set `INVALID_CHAR_HANDLING` to `REPLACE`.
// Experimental.
func StorageParameter_ReplacementChar(value *string) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_ReplacementCharParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"replacementChar",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A property that sets number of rows to skip at the beginning of each source file.
// Experimental.
func StorageParameter_SerializationNullFormat(value *string) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_SerializationNullFormatParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"serializationNullFormat",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The number of rows to skip at the top of a CSV file when the table is being created.
// Experimental.
func StorageParameter_SkipHeaderLineCount(value *float64) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_SkipHeaderLineCountParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"skipHeaderLineCount",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Specifies how to handle data being loaded that exceeds the length of the data type defined for columns containing VARBYTE data.
//
// By default, Redshift Spectrum sets the value to null for data that exceeds the width of the column.
// Experimental.
func StorageParameter_SurplusBytesHandling(value SurplusBytesHandlingAction) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_SurplusBytesHandlingParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"surplusBytesHandling",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Specifies how to handle data being loaded that exceeds the length of the data type defined for columns containing VARCHAR, CHAR, or string data.
//
// By default, Redshift Spectrum sets the value to null for data that exceeds the width of the column.
// Experimental.
func StorageParameter_SurplusCharHandling(value SurplusCharHandlingAction) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_SurplusCharHandlingParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"surplusCharHandling",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// You can specify an AWS Key Management Service key to enable Server–Side Encryption (SSE) for Amazon S3 objects.
// Experimental.
func StorageParameter_WriteKmsKeyId(value *string) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_WriteKmsKeyIdParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"writeKmsKeyId",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A property that sets the maximum size (in MB) of each file written to Amazon S3 by CREATE EXTERNAL TABLE AS.
//
// The size must be a valid integer between 5 and 6200. The default maximum file size is 6,200 MB. This table property also applies to any subsequent INSERT statement into the same external table.
// Experimental.
func StorageParameter_WriteMaxFileSizeMb(value *float64) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_WriteMaxFileSizeMbParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"writeMaxFileSizeMb",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A property that sets whether CREATE EXTERNAL TABLE AS should write data in parallel.
//
// When 'write.parallel' is set to off, CREATE EXTERNAL TABLE AS writes to one or more data files serially onto Amazon S3. This table property also applies to any subsequent INSERT statement into the same external table.
// Experimental.
func StorageParameter_WriteParallel(value WriteParallel) StorageParameter {
	_init_.Initialize()

	if err := validateStorageParameter_WriteParallelParameters(value); err != nil {
		panic(err)
	}
	var returns StorageParameter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.StorageParameter",
		"writeParallel",
		[]interface{}{value},
		&returns,
	)

	return returns
}

