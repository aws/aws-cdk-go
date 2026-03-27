package previewawsapigatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnRestApiExecutionLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiExecutionLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnRestApiExecutionLogsOutputFormat()
//
// Experimental.
type CfnRestApiExecutionLogsOutputFormat interface {
}

// The jsii proxy struct for CfnRestApiExecutionLogsOutputFormat
type jsiiProxy_CfnRestApiExecutionLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnRestApiExecutionLogsOutputFormat() CfnRestApiExecutionLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnRestApiExecutionLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRestApiExecutionLogsOutputFormat_Override(c CfnRestApiExecutionLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

