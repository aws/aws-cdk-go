package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCodeInterpreterCustomUsageLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomUsageLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCodeInterpreterCustomUsageLogsOutputFormat()
//
// Experimental.
type CfnCodeInterpreterCustomUsageLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCodeInterpreterCustomUsageLogsOutputFormat
type jsiiProxy_CfnCodeInterpreterCustomUsageLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCodeInterpreterCustomUsageLogsOutputFormat() CfnCodeInterpreterCustomUsageLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeInterpreterCustomUsageLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCodeInterpreterCustomUsageLogsOutputFormat_Override(c CfnCodeInterpreterCustomUsageLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

