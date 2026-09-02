package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnFirewallDenyLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallDenyLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnFirewallDenyLogsOutputFormat()
//
// Experimental.
type CfnFirewallDenyLogsOutputFormat interface {
}

// The jsii proxy struct for CfnFirewallDenyLogsOutputFormat
type jsiiProxy_CfnFirewallDenyLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnFirewallDenyLogsOutputFormat() CfnFirewallDenyLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallDenyLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnFirewallDenyLogsOutputFormat_Override(c CfnFirewallDenyLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

