package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents possible output formats when performing record data conversion.
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
type OutputFormat interface {
}

// The jsii proxy struct for OutputFormat
type jsiiProxy_OutputFormat struct {
	_ byte // padding
}

func OutputFormat_ORC() OrcOutputFormat {
	_init_.Initialize()
	var returns OrcOutputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.OutputFormat",
		"ORC",
		&returns,
	)
	return returns
}

func OutputFormat_PARQUET() ParquetOutputFormat {
	_init_.Initialize()
	var returns ParquetOutputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.OutputFormat",
		"PARQUET",
		&returns,
	)
	return returns
}

