package awscdkgluealpha


// The compression type.
//
// Example:
//   var myDatabase Database
//
//   glue.NewS3Table(this, jsii.String("MyTable"), &S3TableProps{
//   	StorageParameters: []StorageParameter{
//   		glue.StorageParameter_SkipHeaderLineCount(jsii.Number(1)),
//   		glue.StorageParameter_CompressionType(glue.CompressionType_GZIP),
//   		glue.StorageParameter_Custom(jsii.String("separatorChar"), jsii.String(",")),
//   	},
//   	// ...
//   	Database: myDatabase,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("col1"),
//   			Type: glue.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: glue.DataFormat_JSON(),
//   })
//
// See: https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_EXTERNAL_TABLE.html#r_CREATE_EXTERNAL_TABLE-parameters - under _"TABLE PROPERTIES"_ > _"compression_type"_
//
// Experimental.
type CompressionType string

const (
	// No compression.
	// Experimental.
	CompressionType_NONE CompressionType = "NONE"
	// Burrows-Wheeler compression.
	// Experimental.
	CompressionType_BZIP2 CompressionType = "BZIP2"
	// Deflate compression.
	// Experimental.
	CompressionType_GZIP CompressionType = "GZIP"
	// Compression algorithm focused on high compression and decompression speeds, rather than the maximum possible compression.
	// Experimental.
	CompressionType_SNAPPY CompressionType = "SNAPPY"
)

