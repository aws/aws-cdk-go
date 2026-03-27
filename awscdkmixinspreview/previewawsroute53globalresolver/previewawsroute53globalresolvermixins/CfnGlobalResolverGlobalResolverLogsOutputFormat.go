package previewawsroute53globalresolvermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnGlobalResolverGlobalResolverLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGlobalResolverGlobalResolverLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnGlobalResolverGlobalResolverLogsOutputFormat()
//
// Experimental.
type CfnGlobalResolverGlobalResolverLogsOutputFormat interface {
}

// The jsii proxy struct for CfnGlobalResolverGlobalResolverLogsOutputFormat
type jsiiProxy_CfnGlobalResolverGlobalResolverLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnGlobalResolverGlobalResolverLogsOutputFormat() CfnGlobalResolverGlobalResolverLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnGlobalResolverGlobalResolverLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnGlobalResolverGlobalResolverLogsOutputFormat_Override(c CfnGlobalResolverGlobalResolverLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverGlobalResolverLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

