package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCodeInterpreterUsageLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterUsageLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCodeInterpreterUsageLogsOutputFormat()
//
// Experimental.
type CfnCodeInterpreterUsageLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCodeInterpreterUsageLogsOutputFormat
type jsiiProxy_CfnCodeInterpreterUsageLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCodeInterpreterUsageLogsOutputFormat() CfnCodeInterpreterUsageLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeInterpreterUsageLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterUsageLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCodeInterpreterUsageLogsOutputFormat_Override(c CfnCodeInterpreterUsageLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterUsageLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

