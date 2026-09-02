package previewawsrtbfabricmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnLinkApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLinkApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnLinkApplicationLogsOutputFormat()
//
// Experimental.
type CfnLinkApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnLinkApplicationLogsOutputFormat
type jsiiProxy_CfnLinkApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnLinkApplicationLogsOutputFormat() CfnLinkApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnLinkApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLinkApplicationLogsOutputFormat_Override(c CfnLinkApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rtbfabric.mixins.CfnLinkApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

