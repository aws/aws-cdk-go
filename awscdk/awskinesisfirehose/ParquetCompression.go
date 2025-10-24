package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible compression options available for Parquet OutputFormat.
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
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-parquetserde.html#cfn-kinesisfirehose-deliverystream-parquetserde-compression
//
type ParquetCompression interface {
	// the string value of the Serde Compression.
	Value() *string
}

// The jsii proxy struct for ParquetCompression
type jsiiProxy_ParquetCompression struct {
	_ byte // padding
}

func (j *jsiiProxy_ParquetCompression) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates a new ParquetCompression instance with a custom value.
func ParquetCompression_Of(value *string) ParquetCompression {
	_init_.Initialize()

	if err := validateParquetCompression_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ParquetCompression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetCompression",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ParquetCompression_GZIP() ParquetCompression {
	_init_.Initialize()
	var returns ParquetCompression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetCompression",
		"GZIP",
		&returns,
	)
	return returns
}

func ParquetCompression_SNAPPY() ParquetCompression {
	_init_.Initialize()
	var returns ParquetCompression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetCompression",
		"SNAPPY",
		&returns,
	)
	return returns
}

func ParquetCompression_UNCOMPRESSED() ParquetCompression {
	_init_.Initialize()
	var returns ParquetCompression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetCompression",
		"UNCOMPRESSED",
		&returns,
	)
	return returns
}

