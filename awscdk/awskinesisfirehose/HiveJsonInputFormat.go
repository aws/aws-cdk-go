package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// This class specifies properties for Hive JSON input format for record format conversion.
//
// You should only need to specify an instance of this class if the default configuration does not suit your needs.
//
// Example:
//   inputFormat := firehose.NewHiveJsonInputFormat(&HiveJsonInputFormatProps{
//   	TimestampParsers: []TimestampParser{
//   		firehose.TimestampParser_FromFormatString(jsii.String("yyyy-MM-dd")),
//   		firehose.TimestampParser_EPOCH_MILLIS(),
//   	},
//   })
//
type HiveJsonInputFormat interface {
	IInputFormat
	// Properties for Hive JSON input format.
	Props() *HiveJsonInputFormatProps
	// Renders the cloudformation properties for the input format.
	CreateInputFormatConfig() *CfnDeliveryStream_InputFormatConfigurationProperty
}

// The jsii proxy struct for HiveJsonInputFormat
type jsiiProxy_HiveJsonInputFormat struct {
	jsiiProxy_IInputFormat
}

func (j *jsiiProxy_HiveJsonInputFormat) Props() *HiveJsonInputFormatProps {
	var returns *HiveJsonInputFormatProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewHiveJsonInputFormat(props *HiveJsonInputFormatProps) HiveJsonInputFormat {
	_init_.Initialize()

	if err := validateNewHiveJsonInputFormatParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HiveJsonInputFormat{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.HiveJsonInputFormat",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewHiveJsonInputFormat_Override(h HiveJsonInputFormat, props *HiveJsonInputFormatProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.HiveJsonInputFormat",
		[]interface{}{props},
		h,
	)
}

func (h *jsiiProxy_HiveJsonInputFormat) CreateInputFormatConfig() *CfnDeliveryStream_InputFormatConfigurationProperty {
	var returns *CfnDeliveryStream_InputFormatConfigurationProperty

	_jsii_.Invoke(
		h,
		"createInputFormatConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

