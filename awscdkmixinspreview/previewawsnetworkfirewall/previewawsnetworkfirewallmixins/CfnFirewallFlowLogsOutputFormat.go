package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnFirewallFlowLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallFlowLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnFirewallFlowLogsOutputFormat()
//
// Experimental.
type CfnFirewallFlowLogsOutputFormat interface {
}

// The jsii proxy struct for CfnFirewallFlowLogsOutputFormat
type jsiiProxy_CfnFirewallFlowLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnFirewallFlowLogsOutputFormat() CfnFirewallFlowLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnFirewallFlowLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnFirewallFlowLogsOutputFormat_Override(c CfnFirewallFlowLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

