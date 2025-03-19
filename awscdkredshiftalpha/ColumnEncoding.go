package awscdkredshiftalpha


// The compression encoding of a column.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//
//
//   awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
//   	TableColumns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			DataType: jsii.String("varchar(4)"),
//   			Encoding: awscdkredshiftalpha.ColumnEncoding_TEXT32K,
//   		},
//   		&column{
//   			Name: jsii.String("col2"),
//   			DataType: jsii.String("float"),
//   			Encoding: awscdkredshiftalpha.ColumnEncoding_DELTA32K,
//   		},
//   	},
//   	Cluster: cluster,
//   	DatabaseName: jsii.String("databaseName"),
//   })
//
// See: https://docs.aws.amazon.com/redshift/latest/dg/c_Compression_encodings.html
//
// Experimental.
type ColumnEncoding string

const (
	// Amazon Redshift assigns an optimal encoding based on the column data.
	//
	// This is the default.
	// Experimental.
	ColumnEncoding_AUTO ColumnEncoding = "AUTO"
	// The column is not compressed.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_Raw_encoding.html
	//
	// Experimental.
	ColumnEncoding_RAW ColumnEncoding = "RAW"
	// The column is compressed using the AZ64 algorithm.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/az64-encoding.html
	//
	// Experimental.
	ColumnEncoding_AZ64 ColumnEncoding = "AZ64"
	// The column is compressed using a separate dictionary for each block column value on disk.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_Byte_dictionary_encoding.html
	//
	// Experimental.
	ColumnEncoding_BYTEDICT ColumnEncoding = "BYTEDICT"
	// The column is compressed based on the difference between values in the column.
	//
	// This records differences as 1-byte values.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_Delta_encoding.html
	//
	// Experimental.
	ColumnEncoding_DELTA ColumnEncoding = "DELTA"
	// The column is compressed based on the difference between values in the column.
	//
	// This records differences as 2-byte values.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_Delta_encoding.html
	//
	// Experimental.
	ColumnEncoding_DELTA32K ColumnEncoding = "DELTA32K"
	// The column is compressed using the LZO algorithm.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/lzo-encoding.html
	//
	// Experimental.
	ColumnEncoding_LZO ColumnEncoding = "LZO"
	// The column is compressed to a smaller storage size than the original data type.
	//
	// The compressed storage size is 1 byte.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_MostlyN_encoding.html
	//
	// Experimental.
	ColumnEncoding_MOSTLY8 ColumnEncoding = "MOSTLY8"
	// The column is compressed to a smaller storage size than the original data type.
	//
	// The compressed storage size is 2 bytes.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_MostlyN_encoding.html
	//
	// Experimental.
	ColumnEncoding_MOSTLY16 ColumnEncoding = "MOSTLY16"
	// The column is compressed to a smaller storage size than the original data type.
	//
	// The compressed storage size is 4 bytes.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_MostlyN_encoding.html
	//
	// Experimental.
	ColumnEncoding_MOSTLY32 ColumnEncoding = "MOSTLY32"
	// The column is compressed by recording the number of occurrences of each value in the column.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_Runlength_encoding.html
	//
	// Experimental.
	ColumnEncoding_RUNLENGTH ColumnEncoding = "RUNLENGTH"
	// The column is compressed by recording the first 245 unique words and then using a 1-byte index to represent each word.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_Text255_encoding.html
	//
	// Experimental.
	ColumnEncoding_TEXT255 ColumnEncoding = "TEXT255"
	// The column is compressed by recording the first 32K unique words and then using a 2-byte index to represent each word.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/c_Text255_encoding.html
	//
	// Experimental.
	ColumnEncoding_TEXT32K ColumnEncoding = "TEXT32K"
	// The column is compressed using the ZSTD algorithm.
	// See: https://docs.aws.amazon.com/redshift/latest/dg/zstd-encoding.html
	//
	// Experimental.
	ColumnEncoding_ZSTD ColumnEncoding = "ZSTD"
)

