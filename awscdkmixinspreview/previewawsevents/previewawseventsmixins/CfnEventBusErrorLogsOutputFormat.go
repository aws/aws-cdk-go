package previewawseventsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnEventBusErrorLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusErrorLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnEventBusErrorLogsOutputFormat()
//
// Experimental.
type CfnEventBusErrorLogsOutputFormat interface {
}

// The jsii proxy struct for CfnEventBusErrorLogsOutputFormat
type jsiiProxy_CfnEventBusErrorLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnEventBusErrorLogsOutputFormat() CfnEventBusErrorLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnEventBusErrorLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnEventBusErrorLogsOutputFormat_Override(c CfnEventBusErrorLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusErrorLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

