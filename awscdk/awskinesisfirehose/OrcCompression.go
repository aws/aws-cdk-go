package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible compression options available for ORC OutputFormat.
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
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-orcserde.html#cfn-kinesisfirehose-deliverystream-orcserde-compression
//
type OrcCompression interface {
	// the string value of the Serde Compression.
	Value() *string
}

// The jsii proxy struct for OrcCompression
type jsiiProxy_OrcCompression struct {
	_ byte // padding
}

func (j *jsiiProxy_OrcCompression) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates a new OrcCompression instance with a custom value.
func OrcCompression_Of(value *string) OrcCompression {
	_init_.Initialize()

	if err := validateOrcCompression_OfParameters(value); err != nil {
		panic(err)
	}
	var returns OrcCompression

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.OrcCompression",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func OrcCompression_NONE() OrcCompression {
	_init_.Initialize()
	var returns OrcCompression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.OrcCompression",
		"NONE",
		&returns,
	)
	return returns
}

func OrcCompression_SNAPPY() OrcCompression {
	_init_.Initialize()
	var returns OrcCompression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.OrcCompression",
		"SNAPPY",
		&returns,
	)
	return returns
}

func OrcCompression_ZLIB() OrcCompression {
	_init_.Initialize()
	var returns OrcCompression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.OrcCompression",
		"ZLIB",
		&returns,
	)
	return returns
}

