package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnApplicationEventLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationEventLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnApplicationEventLogsOutputFormat()
//
// Experimental.
type CfnApplicationEventLogsOutputFormat interface {
}

// The jsii proxy struct for CfnApplicationEventLogsOutputFormat
type jsiiProxy_CfnApplicationEventLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnApplicationEventLogsOutputFormat() CfnApplicationEventLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationEventLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnApplicationEventLogsOutputFormat_Override(c CfnApplicationEventLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationEventLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

