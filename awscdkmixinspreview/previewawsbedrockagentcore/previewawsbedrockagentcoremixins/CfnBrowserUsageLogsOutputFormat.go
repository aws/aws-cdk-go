package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnBrowserUsageLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserUsageLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnBrowserUsageLogsOutputFormat()
//
// Experimental.
type CfnBrowserUsageLogsOutputFormat interface {
}

// The jsii proxy struct for CfnBrowserUsageLogsOutputFormat
type jsiiProxy_CfnBrowserUsageLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnBrowserUsageLogsOutputFormat() CfnBrowserUsageLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnBrowserUsageLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBrowserUsageLogsOutputFormat_Override(c CfnBrowserUsageLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

