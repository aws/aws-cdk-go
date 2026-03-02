package previewawswisdommixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnAssistantEventLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssistantEventLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnAssistantEventLogsOutputFormat()
//
// Experimental.
type CfnAssistantEventLogsOutputFormat interface {
}

// The jsii proxy struct for CfnAssistantEventLogsOutputFormat
type jsiiProxy_CfnAssistantEventLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnAssistantEventLogsOutputFormat() CfnAssistantEventLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnAssistantEventLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnAssistantEventLogsOutputFormat_Override(c CfnAssistantEventLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

