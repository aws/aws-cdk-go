package previewawspipesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnPipeExecutionLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipeExecutionLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnPipeExecutionLogsOutputFormat()
//
// Experimental.
type CfnPipeExecutionLogsOutputFormat interface {
}

// The jsii proxy struct for CfnPipeExecutionLogsOutputFormat
type jsiiProxy_CfnPipeExecutionLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnPipeExecutionLogsOutputFormat() CfnPipeExecutionLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnPipeExecutionLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnPipeExecutionLogsOutputFormat_Override(c CfnPipeExecutionLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

