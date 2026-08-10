package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCodeInterpreterApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCodeInterpreterApplicationLogsOutputFormat()
//
// Experimental.
type CfnCodeInterpreterApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCodeInterpreterApplicationLogsOutputFormat
type jsiiProxy_CfnCodeInterpreterApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCodeInterpreterApplicationLogsOutputFormat() CfnCodeInterpreterApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeInterpreterApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCodeInterpreterApplicationLogsOutputFormat_Override(c CfnCodeInterpreterApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

