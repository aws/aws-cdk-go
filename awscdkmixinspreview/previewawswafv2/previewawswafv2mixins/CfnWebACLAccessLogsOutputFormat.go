package previewawswafv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnWebACLAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnWebACLAccessLogsOutputFormat()
//
// Experimental.
type CfnWebACLAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnWebACLAccessLogsOutputFormat
type jsiiProxy_CfnWebACLAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnWebACLAccessLogsOutputFormat() CfnWebACLAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACLAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnWebACLAccessLogsOutputFormat_Override(c CfnWebACLAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

