package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnVPCRoute53ResolverQueryLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCRoute53ResolverQueryLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnVPCRoute53ResolverQueryLogsOutputFormat()
//
// Experimental.
type CfnVPCRoute53ResolverQueryLogsOutputFormat interface {
}

// The jsii proxy struct for CfnVPCRoute53ResolverQueryLogsOutputFormat
type jsiiProxy_CfnVPCRoute53ResolverQueryLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnVPCRoute53ResolverQueryLogsOutputFormat() CfnVPCRoute53ResolverQueryLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnVPCRoute53ResolverQueryLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnVPCRoute53ResolverQueryLogsOutputFormat_Override(c CfnVPCRoute53ResolverQueryLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCRoute53ResolverQueryLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

