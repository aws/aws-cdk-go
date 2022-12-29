// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha


// Properties for a Firewall Rule Group.
//
// Example:
//   var myBlockList firewallDomainList
//
//   route53resolver.NewFirewallRuleGroup(this, jsii.String("RuleGroup"), &firewallRuleGroupProps{
//   	rules: []firewallRule{
//   		&firewallRule{
//   			priority: jsii.Number(10),
//   			firewallDomainList: myBlockList,
//   			// block and reply with NODATA
//   			action: route53resolver.firewallRuleAction.block(),
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

