package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnInstanceConsoleLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceConsoleLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnInstanceConsoleLogsOutputFormat()
//
// Experimental.
type CfnInstanceConsoleLogsOutputFormat interface {
}

// The jsii proxy struct for CfnInstanceConsoleLogsOutputFormat
type jsiiProxy_CfnInstanceConsoleLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnInstanceConsoleLogsOutputFormat() CfnInstanceConsoleLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnInstanceConsoleLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnInstanceConsoleLogsOutputFormat_Override(c CfnInstanceConsoleLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnInstanceConsoleLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

