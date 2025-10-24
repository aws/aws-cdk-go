package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// This class specifies properties for ORC output format for record format conversion.
//
// You should only need to specify an instance of this class if the default configuration does not suit your needs.
//
// Example:
//   outputFormat := firehose.NewOrcOutputFormat(&OrcOutputFormatProps{
//   	FormatVersion: firehose.OrcFormatVersion_V0_11,
//   	BlockSize: awscdk.Size_Mebibytes(jsii.Number(256)),
//   	Compression: firehose.OrcCompression_NONE(),
//   	BloomFilterColumns: []*string{
//   		jsii.String("columnA"),
//   	},
//   	BloomFilterFalsePositiveProbability: jsii.Number(0.1),
//   	DictionaryKeyThreshold: jsii.Number(0.7),
//   	EnablePadding: jsii.Boolean(true),
//   	PaddingTolerance: jsii.Number(0.2),
//   	RowIndexStride: jsii.Number(9000),
//   	StripeSize: awscdk.Size_*Mebibytes(jsii.Number(32)),
//   })
//
type OrcOutputFormat interface {
	IOutputFormat
	// Properties for the ORC output format.
	Props() *OrcOutputFormatProps
	// Renders the cloudformation properties for the output format.
	CreateOutputFormatConfig() *CfnDeliveryStream_OutputFormatConfigurationProperty
}

// The jsii proxy struct for OrcOutputFormat
type jsiiProxy_OrcOutputFormat struct {
	jsiiProxy_IOutputFormat
}

func (j *jsiiProxy_OrcOutputFormat) Props() *OrcOutputFormatProps {
	var returns *OrcOutputFormatProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewOrcOutputFormat(props *OrcOutputFormatProps) OrcOutputFormat {
	_init_.Initialize()

	if err := validateNewOrcOutputFormatParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrcOutputFormat{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.OrcOutputFormat",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewOrcOutputFormat_Override(o OrcOutputFormat, props *OrcOutputFormatProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.OrcOutputFormat",
		[]interface{}{props},
		o,
	)
}

func (o *jsiiProxy_OrcOutputFormat) CreateOutputFormatConfig() *CfnDeliveryStream_OutputFormatConfigurationProperty {
	var returns *CfnDeliveryStream_OutputFormatConfigurationProperty

	_jsii_.Invoke(
		o,
		"createOutputFormatConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

