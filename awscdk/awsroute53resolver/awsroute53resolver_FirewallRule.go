package awsroute53resolver


// A Firewall Rule.
//
// Example:
//   var myBlockList firewallDomainList
//   var ruleGroup firewallRuleGroup
//
//
//   ruleGroup.AddRule(&FirewallRule{
//   	Priority: jsii.Number(10),
//   	FirewallDomainList: myBlockList,
//   	// block and reply with NXDOMAIN
//   	Action: route53resolver.FirewallRuleAction_Block(route53resolver.DnsBlockResponse_NxDomain()),
//   })
//
//   ruleGroup.AddRule(&FirewallRule{
//   	Priority: jsii.Number(20),
//   	FirewallDomainList: myBlockList,
//   	// block and override DNS response with a custom domain
//   	Action: route53resolver.FirewallRuleAction_*Block(route53resolver.DnsBlockResponse_Override(jsii.String("amazon.com"))),
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

