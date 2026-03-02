package previewawsroute53profilesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnProfileRoute53ProfilesResolverQueryLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProfileRoute53ProfilesResolverQueryLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnProfileRoute53ProfilesResolverQueryLogsOutputFormat()
//
// Experimental.
type CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat interface {
}

// The jsii proxy struct for CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat
type jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnProfileRoute53ProfilesResolverQueryLogsOutputFormat() CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_Override(c CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

