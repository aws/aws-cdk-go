package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnMailManagerIngressPointApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnMailManagerIngressPointApplicationLogsOutputFormat()
//
// Experimental.
type CfnMailManagerIngressPointApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnMailManagerIngressPointApplicationLogsOutputFormat
type jsiiProxy_CfnMailManagerIngressPointApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnMailManagerIngressPointApplicationLogsOutputFormat() CfnMailManagerIngressPointApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnMailManagerIngressPointApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMailManagerIngressPointApplicationLogsOutputFormat_Override(c CfnMailManagerIngressPointApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

