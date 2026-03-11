package previewawseventsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnEventBusInfoLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusInfoLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnEventBusInfoLogsOutputFormat()
//
// Experimental.
type CfnEventBusInfoLogsOutputFormat interface {
}

// The jsii proxy struct for CfnEventBusInfoLogsOutputFormat
type jsiiProxy_CfnEventBusInfoLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnEventBusInfoLogsOutputFormat() CfnEventBusInfoLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnEventBusInfoLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnEventBusInfoLogsOutputFormat_Override(c CfnEventBusInfoLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusInfoLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

