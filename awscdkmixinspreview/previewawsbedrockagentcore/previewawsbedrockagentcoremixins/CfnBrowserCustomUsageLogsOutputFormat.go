package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnBrowserCustomUsageLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomUsageLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnBrowserCustomUsageLogsOutputFormat()
//
// Experimental.
type CfnBrowserCustomUsageLogsOutputFormat interface {
}

// The jsii proxy struct for CfnBrowserCustomUsageLogsOutputFormat
type jsiiProxy_CfnBrowserCustomUsageLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnBrowserCustomUsageLogsOutputFormat() CfnBrowserCustomUsageLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnBrowserCustomUsageLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBrowserCustomUsageLogsOutputFormat_Override(c CfnBrowserCustomUsageLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

