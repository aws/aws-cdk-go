package previewawsstepfunctionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnStateMachineStandardLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineStandardLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnStateMachineStandardLogsOutputFormat()
//
// Experimental.
type CfnStateMachineStandardLogsOutputFormat interface {
}

// The jsii proxy struct for CfnStateMachineStandardLogsOutputFormat
type jsiiProxy_CfnStateMachineStandardLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnStateMachineStandardLogsOutputFormat() CfnStateMachineStandardLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnStateMachineStandardLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnStateMachineStandardLogsOutputFormat_Override(c CfnStateMachineStandardLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

