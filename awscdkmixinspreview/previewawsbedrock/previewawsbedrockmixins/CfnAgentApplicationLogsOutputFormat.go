package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnAgentApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnAgentApplicationLogsOutputFormat()
//
// Experimental.
type CfnAgentApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnAgentApplicationLogsOutputFormat
type jsiiProxy_CfnAgentApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnAgentApplicationLogsOutputFormat() CfnAgentApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnAgentApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAgentApplicationLogsOutputFormat_Override(c CfnAgentApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

