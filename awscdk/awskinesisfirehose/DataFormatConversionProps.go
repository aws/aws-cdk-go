package awskinesisfirehose


// Props for specifying data format conversion for Firehose.
//
// Example:
//   var bucket Bucket
//   var schemaGlueTable CfnTable
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	DataFormatConversion: &DataFormatConversionProps{
//   		SchemaConfiguration: firehose.SchemaConfiguration_FromCfnTable(schemaGlueTable),
//   		InputFormat: firehose.InputFormat_OPENX_JSON(),
//   		OutputFormat: firehose.OutputFormat_PARQUET(),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/firehose/latest/dev/record-format-conversion.html
//
type DataFormatConversionProps struct {
	// The input format to convert from for record format conversion.
	InputFormat IInputFormat `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// The output format to convert to for record format conversion.
	OutputFormat IOutputFormat `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// The schema configuration to use in converting the input format to output format.
	SchemaConfiguration SchemaConfiguration `field:"required" json:"schemaConfiguration" yaml:"schemaConfiguration"`
	// Whether data format conversion is enabled or not.
	// Default: `true`.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

