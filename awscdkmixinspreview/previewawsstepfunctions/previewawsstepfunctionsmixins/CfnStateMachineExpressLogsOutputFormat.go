package previewawsstepfunctionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnStateMachineExpressLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineExpressLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnStateMachineExpressLogsOutputFormat()
//
// Experimental.
type CfnStateMachineExpressLogsOutputFormat interface {
}

// The jsii proxy struct for CfnStateMachineExpressLogsOutputFormat
type jsiiProxy_CfnStateMachineExpressLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnStateMachineExpressLogsOutputFormat() CfnStateMachineExpressLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnStateMachineExpressLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnStateMachineExpressLogsOutputFormat_Override(c CfnStateMachineExpressLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

