package previewawsdsqlmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterAuditLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAuditLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterAuditLogsOutputFormat()
//
// Experimental.
type CfnClusterAuditLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterAuditLogsOutputFormat
type jsiiProxy_CfnClusterAuditLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterAuditLogsOutputFormat() CfnClusterAuditLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterAuditLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterAuditLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterAuditLogsOutputFormat_Override(c CfnClusterAuditLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterAuditLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

