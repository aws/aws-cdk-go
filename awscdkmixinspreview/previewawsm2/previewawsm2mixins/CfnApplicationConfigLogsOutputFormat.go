package previewawsm2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnApplicationConfigLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationConfigLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnApplicationConfigLogsOutputFormat()
//
// Experimental.
type CfnApplicationConfigLogsOutputFormat interface {
}

// The jsii proxy struct for CfnApplicationConfigLogsOutputFormat
type jsiiProxy_CfnApplicationConfigLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnApplicationConfigLogsOutputFormat() CfnApplicationConfigLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationConfigLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnApplicationConfigLogsOutputFormat_Override(c CfnApplicationConfigLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationConfigLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

