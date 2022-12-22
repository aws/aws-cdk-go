package awsroute53resolver


// A Firewall Rule.
//
// Example:
//   var myBlockList firewallDomainList
//   var ruleGroup firewallRuleGroup
//
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(10),
//   	firewallDomainList: myBlockList,
//   	// block and reply with NXDOMAIN
//   	action: route53resolver.firewallRuleAction.block(route53resolver.dnsBlockResponse.nxDomain()),
//   })
//
//   ruleGroup.addRule(&firewallRule{
//   	priority: jsii.Number(20),
//   	firewallDomainList: myBlockList,
//   	// block and override DNS response with a custom domain
//   	action: route53resolver.*firewallRuleAction.block(route53resolver.*dnsBlockResponse.override(jsii.String("amazon.com"))),
//   })
//
// Experimental.
type FirewallRule struct {
	// The action for this rule.
	// Experimental.
	Action FirewallRuleAction `field:"required" json:"action" yaml:"action"`
	// The domain list for this rule.
	// Experimental.
	FirewallDomainList IFirewallDomainList `field:"required" json:"firewallDomainList" yaml:"firewallDomainList"`
	// The priority of the rule in the rule group.
	//
	// This value must be unique within
	// the rule group.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}

