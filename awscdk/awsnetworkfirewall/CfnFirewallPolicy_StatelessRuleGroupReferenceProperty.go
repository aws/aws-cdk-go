package awsnetworkfirewall


// Identifier for a single stateless rule group, used in a firewall policy to refer to the rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statelessRuleGroupReferenceProperty := &StatelessRuleGroupReferenceProperty{
//   	Priority: jsii.Number(123),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statelessrulegroupreference.html
//
type CfnFirewallPolicy_StatelessRuleGroupReferenceProperty struct {
	// An integer setting that indicates the order in which to run the stateless rule groups in a single `FirewallPolicy` .
	//
	// Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting. You must ensure that the priority settings are unique within each policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statelessrulegroupreference.html#cfn-networkfirewall-firewallpolicy-statelessrulegroupreference-priority
	//
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The Amazon Resource Name (ARN) of the stateless rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statelessrulegroupreference.html#cfn-networkfirewall-firewallpolicy-statelessrulegroupreference-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

