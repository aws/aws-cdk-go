package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnMailManagerIngressPointTrafficPolicyDebugLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat()
//
// Experimental.
type CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat interface {
}

// The jsii proxy struct for CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat
type jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat() CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Override(c CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

