package previewawsapigatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnRestApiAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnRestApiAccessLogsOutputFormat()
//
// Experimental.
type CfnRestApiAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnRestApiAccessLogsOutputFormat
type jsiiProxy_CfnRestApiAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnRestApiAccessLogsOutputFormat() CfnRestApiAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnRestApiAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRestApiAccessLogsOutputFormat_Override(c CfnRestApiAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

