package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnFirewallAlertLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallAlertLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnFirewallAlertLogsOutputFormat()
//
// Experimental.
type CfnFirewallAlertLogsOutputFormat interface {
}

// The jsii proxy struct for CfnFirewallAlertLogsOutputFormat
type jsiiProxy_CfnFirewallAlertLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnFirewallAlertLogsOutputFormat() CfnFirewallAlertLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallAlertLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnFirewallAlertLogsOutputFormat_Override(c CfnFirewallAlertLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

