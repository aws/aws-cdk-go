package previewawseventsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnEventBusTraceLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusTraceLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnEventBusTraceLogsOutputFormat()
//
// Experimental.
type CfnEventBusTraceLogsOutputFormat interface {
}

// The jsii proxy struct for CfnEventBusTraceLogsOutputFormat
type jsiiProxy_CfnEventBusTraceLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnEventBusTraceLogsOutputFormat() CfnEventBusTraceLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnEventBusTraceLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnEventBusTraceLogsOutputFormat_Override(c CfnEventBusTraceLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusTraceLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

