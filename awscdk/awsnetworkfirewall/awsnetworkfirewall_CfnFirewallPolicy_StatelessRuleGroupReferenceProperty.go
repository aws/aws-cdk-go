package awsnetworkfirewall


// Identifier for a single stateless rule group, used in a firewall policy to refer to the rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statelessRuleGroupReferenceProperty := &statelessRuleGroupReferenceProperty{
//   	priority: jsii.Number(123),
//   	resourceArn: jsii.String("resourceArn"),
//   }
//
type CfnFirewallPolicy_StatelessRuleGroupReferenceProperty struct {
	// An integer setting that indicates the order in which to run the stateless rule groups in a single `FirewallPolicy` .
	//
	// Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting. You must ensure that the priority settings are unique within each policy.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The Amazon Resource Name (ARN) of the stateless rule group.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

