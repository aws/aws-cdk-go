package previewawsbackupgatewaymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnHypervisorBgwHypervisorLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHypervisorBgwHypervisorLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnHypervisorBgwHypervisorLogsOutputFormat()
//
// Experimental.
type CfnHypervisorBgwHypervisorLogsOutputFormat interface {
}

// The jsii proxy struct for CfnHypervisorBgwHypervisorLogsOutputFormat
type jsiiProxy_CfnHypervisorBgwHypervisorLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnHypervisorBgwHypervisorLogsOutputFormat() CfnHypervisorBgwHypervisorLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnHypervisorBgwHypervisorLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnHypervisorBgwHypervisorLogsOutputFormat_Override(c CfnHypervisorBgwHypervisorLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

