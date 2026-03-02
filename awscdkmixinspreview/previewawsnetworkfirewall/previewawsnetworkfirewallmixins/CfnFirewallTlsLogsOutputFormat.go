package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnFirewallTlsLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallTlsLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnFirewallTlsLogsOutputFormat()
//
// Experimental.
type CfnFirewallTlsLogsOutputFormat interface {
}

// The jsii proxy struct for CfnFirewallTlsLogsOutputFormat
type jsiiProxy_CfnFirewallTlsLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnFirewallTlsLogsOutputFormat() CfnFirewallTlsLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallTlsLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnFirewallTlsLogsOutputFormat_Override(c CfnFirewallTlsLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

