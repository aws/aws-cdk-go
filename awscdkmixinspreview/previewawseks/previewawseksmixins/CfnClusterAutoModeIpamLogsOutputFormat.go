package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterAutoModeIpamLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeIpamLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterAutoModeIpamLogsOutputFormat()
//
// Experimental.
type CfnClusterAutoModeIpamLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterAutoModeIpamLogsOutputFormat
type jsiiProxy_CfnClusterAutoModeIpamLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterAutoModeIpamLogsOutputFormat() CfnClusterAutoModeIpamLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterAutoModeIpamLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterAutoModeIpamLogsOutputFormat_Override(c CfnClusterAutoModeIpamLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeIpamLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

