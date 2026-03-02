package previewawsvpclatticemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnServiceAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnServiceAccessLogsOutputFormat()
//
// Experimental.
type CfnServiceAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnServiceAccessLogsOutputFormat
type jsiiProxy_CfnServiceAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnServiceAccessLogsOutputFormat() CfnServiceAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnServiceAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnServiceAccessLogsOutputFormat_Override(c CfnServiceAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

