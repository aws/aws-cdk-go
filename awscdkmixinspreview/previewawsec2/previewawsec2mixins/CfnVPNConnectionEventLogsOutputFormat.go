package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnVPNConnectionEventLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionEventLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnVPNConnectionEventLogsOutputFormat()
//
// Experimental.
type CfnVPNConnectionEventLogsOutputFormat interface {
}

// The jsii proxy struct for CfnVPNConnectionEventLogsOutputFormat
type jsiiProxy_CfnVPNConnectionEventLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnVPNConnectionEventLogsOutputFormat() CfnVPNConnectionEventLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnVPNConnectionEventLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnVPNConnectionEventLogsOutputFormat_Override(c CfnVPNConnectionEventLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionEventLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

