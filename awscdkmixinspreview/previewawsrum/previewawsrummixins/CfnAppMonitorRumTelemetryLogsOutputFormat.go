package previewawsrummixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnAppMonitorRumTelemetryLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorRumTelemetryLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnAppMonitorRumTelemetryLogsOutputFormat()
//
// Experimental.
type CfnAppMonitorRumTelemetryLogsOutputFormat interface {
}

// The jsii proxy struct for CfnAppMonitorRumTelemetryLogsOutputFormat
type jsiiProxy_CfnAppMonitorRumTelemetryLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnAppMonitorRumTelemetryLogsOutputFormat() CfnAppMonitorRumTelemetryLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnAppMonitorRumTelemetryLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAppMonitorRumTelemetryLogsOutputFormat_Override(c CfnAppMonitorRumTelemetryLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumTelemetryLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

