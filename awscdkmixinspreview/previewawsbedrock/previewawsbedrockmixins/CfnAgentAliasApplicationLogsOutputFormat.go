package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnAgentAliasApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnAgentAliasApplicationLogsOutputFormat()
//
// Experimental.
type CfnAgentAliasApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnAgentAliasApplicationLogsOutputFormat
type jsiiProxy_CfnAgentAliasApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnAgentAliasApplicationLogsOutputFormat() CfnAgentAliasApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnAgentAliasApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAgentAliasApplicationLogsOutputFormat_Override(c CfnAgentAliasApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

