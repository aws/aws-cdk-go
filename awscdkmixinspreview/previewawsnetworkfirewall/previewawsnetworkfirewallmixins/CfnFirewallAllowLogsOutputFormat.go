package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnFirewallAllowLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallAllowLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnFirewallAllowLogsOutputFormat()
//
// Experimental.
type CfnFirewallAllowLogsOutputFormat interface {
}

// The jsii proxy struct for CfnFirewallAllowLogsOutputFormat
type jsiiProxy_CfnFirewallAllowLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnFirewallAllowLogsOutputFormat() CfnFirewallAllowLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallAllowLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnFirewallAllowLogsOutputFormat_Override(c CfnFirewallAllowLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

