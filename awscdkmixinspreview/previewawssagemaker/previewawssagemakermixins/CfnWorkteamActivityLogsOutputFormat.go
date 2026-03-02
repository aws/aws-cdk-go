package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnWorkteamActivityLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkteamActivityLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnWorkteamActivityLogsOutputFormat()
//
// Experimental.
type CfnWorkteamActivityLogsOutputFormat interface {
}

// The jsii proxy struct for CfnWorkteamActivityLogsOutputFormat
type jsiiProxy_CfnWorkteamActivityLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnWorkteamActivityLogsOutputFormat() CfnWorkteamActivityLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkteamActivityLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnWorkteamActivityLogsOutputFormat_Override(c CfnWorkteamActivityLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

