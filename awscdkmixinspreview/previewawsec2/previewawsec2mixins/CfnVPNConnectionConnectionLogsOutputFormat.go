package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnVPNConnectionConnectionLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionConnectionLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnVPNConnectionConnectionLogsOutputFormat()
//
// Experimental.
type CfnVPNConnectionConnectionLogsOutputFormat interface {
}

// The jsii proxy struct for CfnVPNConnectionConnectionLogsOutputFormat
type jsiiProxy_CfnVPNConnectionConnectionLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnVPNConnectionConnectionLogsOutputFormat() CfnVPNConnectionConnectionLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnVPNConnectionConnectionLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnVPNConnectionConnectionLogsOutputFormat_Override(c CfnVPNConnectionConnectionLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPNConnectionConnectionLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

