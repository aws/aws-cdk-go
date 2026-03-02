package previewawsshieldmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnProtectionFlowLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProtectionFlowLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnProtectionFlowLogsOutputFormat()
//
// Experimental.
type CfnProtectionFlowLogsOutputFormat interface {
}

// The jsii proxy struct for CfnProtectionFlowLogsOutputFormat
type jsiiProxy_CfnProtectionFlowLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnProtectionFlowLogsOutputFormat() CfnProtectionFlowLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnProtectionFlowLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnProtectionFlowLogsOutputFormat_Override(c CfnProtectionFlowLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionFlowLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

