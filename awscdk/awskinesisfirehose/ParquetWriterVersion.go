package awskinesisfirehose


// The available WriterVersions for Parquet output format.
//
// Example:
//   outputFormat := firehose.NewParquetOutputFormat(&ParquetOutputFormatProps{
//   	BlockSize: awscdk.Size_Mebibytes(jsii.Number(512)),
//   	Compression: firehose.ParquetCompression_UNCOMPRESSED(),
//   	EnableDictionaryCompression: jsii.Boolean(true),
//   	MaxPadding: awscdk.Size_Bytes(jsii.Number(10)),
//   	PageSize: awscdk.Size_*Mebibytes(jsii.Number(2)),
//   	WriterVersion: firehose.ParquetWriterVersion_V2,
//   })
//
type ParquetWriterVersion string

const (
	// Use V1 Parquet writer version when writing the output.
	ParquetWriterVersion_V1 ParquetWriterVersion = "V1"
	// Use V2 Parquet writer version when writing the output.
	ParquetWriterVersion_V2 ParquetWriterVersion = "V2"
)

