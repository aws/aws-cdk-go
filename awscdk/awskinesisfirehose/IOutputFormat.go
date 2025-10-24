package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An output format to be used in Firehose record format conversion.
type IOutputFormat interface {
	// Renders the cloudformation properties for the output format.
	CreateOutputFormatConfig() *CfnDeliveryStream_OutputFormatConfigurationProperty
}

// The jsii proxy for IOutputFormat
type jsiiProxy_IOutputFormat struct {
	_ byte // padding
}

func (i *jsiiProxy_IOutputFormat) CreateOutputFormatConfig() *CfnDeliveryStream_OutputFormatConfigurationProperty {
	var returns *CfnDeliveryStream_OutputFormatConfigurationProperty

	_jsii_.Invoke(
		i,
		"createOutputFormatConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

