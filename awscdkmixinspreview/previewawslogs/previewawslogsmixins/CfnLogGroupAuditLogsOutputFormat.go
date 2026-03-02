package previewawslogsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnLogGroupAuditLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLogGroupAuditLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnLogGroupAuditLogsOutputFormat()
//
// Experimental.
type CfnLogGroupAuditLogsOutputFormat interface {
}

// The jsii proxy struct for CfnLogGroupAuditLogsOutputFormat
type jsiiProxy_CfnLogGroupAuditLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnLogGroupAuditLogsOutputFormat() CfnLogGroupAuditLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnLogGroupAuditLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLogGroupAuditLogsOutputFormat_Override(c CfnLogGroupAuditLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.mixins.CfnLogGroupAuditLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

