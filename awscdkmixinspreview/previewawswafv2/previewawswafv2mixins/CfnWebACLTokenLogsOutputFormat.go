package previewawswafv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnWebACLTokenLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLTokenLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnWebACLTokenLogsOutputFormat()
//
// Experimental.
type CfnWebACLTokenLogsOutputFormat interface {
}

// The jsii proxy struct for CfnWebACLTokenLogsOutputFormat
type jsiiProxy_CfnWebACLTokenLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnWebACLTokenLogsOutputFormat() CfnWebACLTokenLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnWebACLTokenLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnWebACLTokenLogsOutputFormat_Override(c CfnWebACLTokenLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

