package previewawsvpclatticemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnResourceConfigurationResourceAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceConfigurationResourceAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnResourceConfigurationResourceAccessLogsOutputFormat()
//
// Experimental.
type CfnResourceConfigurationResourceAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnResourceConfigurationResourceAccessLogsOutputFormat
type jsiiProxy_CfnResourceConfigurationResourceAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnResourceConfigurationResourceAccessLogsOutputFormat() CfnResourceConfigurationResourceAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnResourceConfigurationResourceAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnResourceConfigurationResourceAccessLogsOutputFormat_Override(c CfnResourceConfigurationResourceAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

