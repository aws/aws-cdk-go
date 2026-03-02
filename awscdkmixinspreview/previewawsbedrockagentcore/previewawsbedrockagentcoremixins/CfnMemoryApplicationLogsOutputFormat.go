package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnMemoryApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnMemoryApplicationLogsOutputFormat()
//
// Experimental.
type CfnMemoryApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnMemoryApplicationLogsOutputFormat
type jsiiProxy_CfnMemoryApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnMemoryApplicationLogsOutputFormat() CfnMemoryApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnMemoryApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMemoryApplicationLogsOutputFormat_Override(c CfnMemoryApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

