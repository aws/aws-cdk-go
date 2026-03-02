package previewawsbackupgatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnHypervisorDataAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHypervisorDataAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnHypervisorDataAccessLogsOutputFormat()
//
// Experimental.
type CfnHypervisorDataAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnHypervisorDataAccessLogsOutputFormat
type jsiiProxy_CfnHypervisorDataAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnHypervisorDataAccessLogsOutputFormat() CfnHypervisorDataAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnHypervisorDataAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnHypervisorDataAccessLogsOutputFormat_Override(c CfnHypervisorDataAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

