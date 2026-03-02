package previewawsdevopsagentmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnAgentSpaceApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentSpaceApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnAgentSpaceApplicationLogsOutputFormat()
//
// Experimental.
type CfnAgentSpaceApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnAgentSpaceApplicationLogsOutputFormat
type jsiiProxy_CfnAgentSpaceApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnAgentSpaceApplicationLogsOutputFormat() CfnAgentSpaceApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnAgentSpaceApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAgentSpaceApplicationLogsOutputFormat_Override(c CfnAgentSpaceApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

