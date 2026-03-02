package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnAgentAliasEventLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasEventLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnAgentAliasEventLogsOutputFormat()
//
// Experimental.
type CfnAgentAliasEventLogsOutputFormat interface {
}

// The jsii proxy struct for CfnAgentAliasEventLogsOutputFormat
type jsiiProxy_CfnAgentAliasEventLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnAgentAliasEventLogsOutputFormat() CfnAgentAliasEventLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnAgentAliasEventLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAgentAliasEventLogsOutputFormat_Override(c CfnAgentAliasEventLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

