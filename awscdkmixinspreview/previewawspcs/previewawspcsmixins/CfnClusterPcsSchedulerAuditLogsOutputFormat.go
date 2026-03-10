package previewawspcsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterPcsSchedulerAuditLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerAuditLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterPcsSchedulerAuditLogsOutputFormat()
//
// Experimental.
type CfnClusterPcsSchedulerAuditLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterPcsSchedulerAuditLogsOutputFormat
type jsiiProxy_CfnClusterPcsSchedulerAuditLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterPcsSchedulerAuditLogsOutputFormat() CfnClusterPcsSchedulerAuditLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterPcsSchedulerAuditLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterPcsSchedulerAuditLogsOutputFormat_Override(c CfnClusterPcsSchedulerAuditLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerAuditLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

