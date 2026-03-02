package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnRuntimeApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnRuntimeApplicationLogsOutputFormat()
//
// Experimental.
type CfnRuntimeApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnRuntimeApplicationLogsOutputFormat
type jsiiProxy_CfnRuntimeApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnRuntimeApplicationLogsOutputFormat() CfnRuntimeApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnRuntimeApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRuntimeApplicationLogsOutputFormat_Override(c CfnRuntimeApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

