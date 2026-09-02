package previewawsecsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterActionLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterActionLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterActionLogsOutputFormat()
//
// Experimental.
type CfnClusterActionLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterActionLogsOutputFormat
type jsiiProxy_CfnClusterActionLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterActionLogsOutputFormat() CfnClusterActionLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterActionLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterActionLogsOutputFormat_Override(c CfnClusterActionLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnClusterActionLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

