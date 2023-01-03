package awsnetworkfirewall


// Identifier for a single stateful rule group, used in a firewall policy to refer to a rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statefulRuleGroupReferenceProperty := &statefulRuleGroupReferenceProperty{
//   	resourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	override: &statefulRuleGroupOverrideProperty{
//   		action: jsii.String("action"),
//   	},
//   	priority: jsii.Number(123),
//   }
//
type CfnFirewallPolicy_StatefulRuleGroupReferenceProperty struct {
	// The Amazon Resource Name (ARN) of the stateful rule group.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// `CfnFirewallPolicy.StatefulRuleGroupReferenceProperty.Override`.
	Override interface{} `field:"optional" json:"override" yaml:"override"`
	// An integer setting that indicates the order in which to run the stateful rule groups in a single `FirewallPolicy` .
	//
	// This setting only applies to firewall policies that specify the `STRICT_ORDER` rule order in the stateful engine options settings.
	//
	// Network Firewall evalutes each stateful rule group against a packet starting with the group that has the lowest priority setting. You must ensure that the priority settings are unique within each policy.
	//
	// You can change the priority settings of your rule groups at any time. To make it easier to insert rule groups later, number them so there's a wide range in between, for example use 100, 200, and so on.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

