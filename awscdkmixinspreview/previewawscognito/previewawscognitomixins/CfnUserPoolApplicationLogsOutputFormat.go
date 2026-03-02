package previewawscognitomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnUserPoolApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserPoolApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnUserPoolApplicationLogsOutputFormat()
//
// Experimental.
type CfnUserPoolApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnUserPoolApplicationLogsOutputFormat
type jsiiProxy_CfnUserPoolApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnUserPoolApplicationLogsOutputFormat() CfnUserPoolApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnUserPoolApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnUserPoolApplicationLogsOutputFormat_Override(c CfnUserPoolApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

