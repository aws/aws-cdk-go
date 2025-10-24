package awscdkroute53resolveralpha


// Properties for a Firewall Rule Group.
//
// Example:
//   var myBlockList FirewallDomainList
//
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &FirewallRuleGroupProps{
//   	Rules: []FirewallRule{
//   		&FirewallRule{
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
	// Default: - a CloudFormation generated name.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of rules for this group.
	// Default: - no rules.
	//
	// Experimental.
	Rules *[]*FirewallRule `field:"optional" json:"rules" yaml:"rules"`
}

