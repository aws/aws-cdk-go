package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnetworkfirewall/previewawsnetworkfirewallmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the firewall to provide stateful, managed, network firewall and intrusion detection and prevention filtering for your VPCs in Amazon VPC .
//
// The firewall defines the configuration settings for an AWS Network Firewall firewall. The settings include the firewall policy, the subnets in your VPC to use for the firewall endpoints, and any tags that are attached to the firewall AWS resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnFirewallLogsMixin := awscdkmixinspreview.Mixins.NewCfnFirewallLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-firewall.html
//
type CfnFirewallLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFirewallLogsMixin
type jsiiProxy_CfnFirewallLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFirewallLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFirewallLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::NetworkFirewall::Firewall`.
func NewCfnFirewallLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnFirewallLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnFirewallLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFirewallLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::NetworkFirewall::Firewall`.
func NewCfnFirewallLogsMixin_Override(c CfnFirewallLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFirewallLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFirewallLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFirewallLogsMixin_ALERT_LOGS() CfnFirewallAlertLogs {
	_init_.Initialize()
	var returns CfnFirewallAlertLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallLogsMixin",
		"ALERT_LOGS",
		&returns,
	)
	return returns
}

func CfnFirewallLogsMixin_FLOW_LOGS() CfnFirewallFlowLogs {
	_init_.Initialize()
	var returns CfnFirewallFlowLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallLogsMixin",
		"FLOW_LOGS",
		&returns,
	)
	return returns
}

func CfnFirewallLogsMixin_TLS_LOGS() CfnFirewallTlsLogs {
	_init_.Initialize()
	var returns CfnFirewallTlsLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallLogsMixin",
		"TLS_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFirewallLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFirewallLogsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

