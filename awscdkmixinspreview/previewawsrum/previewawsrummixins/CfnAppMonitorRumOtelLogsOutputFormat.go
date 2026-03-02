package previewawsrummixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnAppMonitorRumOtelLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorRumOtelLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnAppMonitorRumOtelLogsOutputFormat()
//
// Experimental.
type CfnAppMonitorRumOtelLogsOutputFormat interface {
}

// The jsii proxy struct for CfnAppMonitorRumOtelLogsOutputFormat
type jsiiProxy_CfnAppMonitorRumOtelLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnAppMonitorRumOtelLogsOutputFormat() CfnAppMonitorRumOtelLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnAppMonitorRumOtelLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAppMonitorRumOtelLogsOutputFormat_Override(c CfnAppMonitorRumOtelLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

