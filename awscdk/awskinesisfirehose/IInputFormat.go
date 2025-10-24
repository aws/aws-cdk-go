package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An input format to be used in Firehose record format conversion.
type IInputFormat interface {
	// Renders the cloudformation properties for the input format.
	CreateInputFormatConfig() *CfnDeliveryStream_InputFormatConfigurationProperty
}

// The jsii proxy for IInputFormat
type jsiiProxy_IInputFormat struct {
	_ byte // padding
}

func (i *jsiiProxy_IInputFormat) CreateInputFormatConfig() *CfnDeliveryStream_InputFormatConfigurationProperty {
	var returns *CfnDeliveryStream_InputFormatConfigurationProperty

	_jsii_.Invoke(
		i,
		"createInputFormatConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

