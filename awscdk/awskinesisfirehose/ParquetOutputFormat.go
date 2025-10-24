package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// This class specifies properties for Parquet output format for record format conversion.
//
// You should only need to specify an instance of this class if the default configuration does not suit your needs.
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
type ParquetOutputFormat interface {
	IOutputFormat
	// Properties for the Parquet output format.
	Props() *ParquetOutputFormatProps
	// Renders the cloudformation properties for the output format.
	CreateOutputFormatConfig() *CfnDeliveryStream_OutputFormatConfigurationProperty
}

// The jsii proxy struct for ParquetOutputFormat
type jsiiProxy_ParquetOutputFormat struct {
	jsiiProxy_IOutputFormat
}

func (j *jsiiProxy_ParquetOutputFormat) Props() *ParquetOutputFormatProps {
	var returns *ParquetOutputFormatProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewParquetOutputFormat(props *ParquetOutputFormatProps) ParquetOutputFormat {
	_init_.Initialize()

	if err := validateNewParquetOutputFormatParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ParquetOutputFormat{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetOutputFormat",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewParquetOutputFormat_Override(p ParquetOutputFormat, props *ParquetOutputFormatProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetOutputFormat",
		[]interface{}{props},
		p,
	)
}

func (p *jsiiProxy_ParquetOutputFormat) CreateOutputFormatConfig() *CfnDeliveryStream_OutputFormatConfigurationProperty {
	var returns *CfnDeliveryStream_OutputFormatConfigurationProperty

	_jsii_.Invoke(
		p,
		"createOutputFormatConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

