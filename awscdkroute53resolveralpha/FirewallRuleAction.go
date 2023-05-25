package awscdkroute53resolveralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Firewall Rule.
//
// Example:
//   var myBlockList firewallDomainList
//
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &FirewallRuleGroupProps{
//   	Rules: []firewallRule{
//   		&firewallRule{
//   			Priority: jsii.Number(10),
//   			FirewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			Action: route53resolver.FirewallRuleAction_Block(),
//   		},
//   	},
//   })
//
// Experimental.
type FirewallRuleAction interface {
	// The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule's domain list.
	// Experimental.
	Action() *string
	// The way that you want DNS Firewall to block the request.
	// Experimental.
	BlockResponse() DnsBlockResponse
}

// The jsii proxy struct for FirewallRuleAction
type jsiiProxy_FirewallRuleAction struct {
	_ byte // padding
}

func (j *jsiiProxy_FirewallRuleAction) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRuleAction) BlockResponse() DnsBlockResponse {
	var returns DnsBlockResponse
	_jsii_.Get(
		j,
		"blockResponse",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallRuleAction_Override(f FirewallRuleAction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		nil, // no parameters
		f,
	)
}

// Permit the request to go through but send an alert to the logs.
// Experimental.
func FirewallRuleAction_Alert() FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		"alert",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Permit the request to go through.
// Experimental.
func FirewallRuleAction_Allow() FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		"allow",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Disallow the request.
// Experimental.
func FirewallRuleAction_Block(response DnsBlockResponse) FirewallRuleAction {
	_init_.Initialize()

	var returns FirewallRuleAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallRuleAction",
		"block",
		[]interface{}{response},
		&returns,
	)

	return returns
}

