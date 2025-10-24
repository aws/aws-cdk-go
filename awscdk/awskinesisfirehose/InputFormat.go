package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents possible input formats when performing record data conversion.
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
type InputFormat interface {
}

// The jsii proxy struct for InputFormat
type jsiiProxy_InputFormat struct {
	_ byte // padding
}

func InputFormat_HIVE_JSON() HiveJsonInputFormat {
	_init_.Initialize()
	var returns HiveJsonInputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.InputFormat",
		"HIVE_JSON",
		&returns,
	)
	return returns
}

func InputFormat_OPENX_JSON() OpenXJsonInputFormat {
	_init_.Initialize()
	var returns OpenXJsonInputFormat
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.InputFormat",
		"OPENX_JSON",
		&returns,
	)
	return returns
}

