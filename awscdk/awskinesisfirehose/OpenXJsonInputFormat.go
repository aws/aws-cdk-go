package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// This class specifies properties for OpenX JSON input format for record format conversion.
//
// You should only need to specify an instance of this class if the default configuration does not suit your needs.
//
// Example:
//   inputFormat := firehose.NewOpenXJsonInputFormat(&OpenXJsonInputFormatProps{
//   	LowercaseColumnNames: jsii.Boolean(false),
//   	ColumnToJsonKeyMappings: map[string]*string{
//   		"ts": jsii.String("timestamp"),
//   	},
//   	ConvertDotsInJsonKeysToUnderscores: jsii.Boolean(true),
//   })
//
type OpenXJsonInputFormat interface {
	IInputFormat
	// Properties for OpenX JSON input format.
	Props() *OpenXJsonInputFormatProps
	// Renders the cloudformation properties for the input format.
	CreateInputFormatConfig() *CfnDeliveryStream_InputFormatConfigurationProperty
}

// The jsii proxy struct for OpenXJsonInputFormat
type jsiiProxy_OpenXJsonInputFormat struct {
	jsiiProxy_IInputFormat
}

func (j *jsiiProxy_OpenXJsonInputFormat) Props() *OpenXJsonInputFormatProps {
	var returns *OpenXJsonInputFormatProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewOpenXJsonInputFormat(props *OpenXJsonInputFormatProps) OpenXJsonInputFormat {
	_init_.Initialize()

	if err := validateNewOpenXJsonInputFormatParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenXJsonInputFormat{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.OpenXJsonInputFormat",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewOpenXJsonInputFormat_Override(o OpenXJsonInputFormat, props *OpenXJsonInputFormatProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.OpenXJsonInputFormat",
		[]interface{}{props},
		o,
	)
}

func (o *jsiiProxy_OpenXJsonInputFormat) CreateInputFormatConfig() *CfnDeliveryStream_InputFormatConfigurationProperty {
	var returns *CfnDeliveryStream_InputFormatConfigurationProperty

	_jsii_.Invoke(
		o,
		"createInputFormatConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

