package previewawsrummixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnAppMonitorRumOtelSpans.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorRumOtelSpansOutputFormat := awscdkmixinspreview.Mixins.NewCfnAppMonitorRumOtelSpansOutputFormat()
//
// Experimental.
type CfnAppMonitorRumOtelSpansOutputFormat interface {
}

// The jsii proxy struct for CfnAppMonitorRumOtelSpansOutputFormat
type jsiiProxy_CfnAppMonitorRumOtelSpansOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnAppMonitorRumOtelSpansOutputFormat() CfnAppMonitorRumOtelSpansOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnAppMonitorRumOtelSpansOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAppMonitorRumOtelSpansOutputFormat_Override(c CfnAppMonitorRumOtelSpansOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorRumOtelSpansOutputFormat",
		nil, // no parameters
		c,
	)
}

