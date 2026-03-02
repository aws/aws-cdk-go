package previewawsb2bimixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnTransformerB2biExecutionLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransformerB2biExecutionLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnTransformerB2biExecutionLogsOutputFormat()
//
// Experimental.
type CfnTransformerB2biExecutionLogsOutputFormat interface {
}

// The jsii proxy struct for CfnTransformerB2biExecutionLogsOutputFormat
type jsiiProxy_CfnTransformerB2biExecutionLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnTransformerB2biExecutionLogsOutputFormat() CfnTransformerB2biExecutionLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnTransformerB2biExecutionLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnTransformerB2biExecutionLogsOutputFormat_Override(c CfnTransformerB2biExecutionLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

