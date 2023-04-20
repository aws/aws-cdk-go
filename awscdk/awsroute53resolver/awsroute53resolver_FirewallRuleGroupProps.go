package awsroute53resolver


// Properties for a Firewall Rule Group.
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
type FirewallRuleGroupProps struct {
	// The name of the rule group.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of rules for this group.
	// Experimental.
	Rules *[]*FirewallRule `field:"optional" json:"rules" yaml:"rules"`
}

