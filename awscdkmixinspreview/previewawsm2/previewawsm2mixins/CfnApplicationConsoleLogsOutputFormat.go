package previewawsm2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnApplicationConsoleLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationConsoleLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnApplicationConsoleLogsOutputFormat()
//
// Experimental.
type CfnApplicationConsoleLogsOutputFormat interface {
}

// The jsii proxy struct for CfnApplicationConsoleLogsOutputFormat
type jsiiProxy_CfnApplicationConsoleLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnApplicationConsoleLogsOutputFormat() CfnApplicationConsoleLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationConsoleLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnApplicationConsoleLogsOutputFormat_Override(c CfnApplicationConsoleLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConsoleLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

