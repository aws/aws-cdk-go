package previewawsnetworkfirewallmixins


// Identifier for a single stateful rule group, used in a firewall policy to refer to a rule group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   statefulRuleGroupReferenceProperty := &StatefulRuleGroupReferenceProperty{
//   	DeepThreatInspection: jsii.Boolean(false),
//   	Override: &StatefulRuleGroupOverrideProperty{
//   		Action: jsii.String("action"),
//   	},
//   	Priority: jsii.Number(123),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupreference.html
//
type CfnFirewallPolicyPropsMixin_StatefulRuleGroupReferenceProperty struct {
	// AWS Network Firewall plans to augment the active threat defense managed rule group with an additional deep threat inspection capability.
	//
	// When this capability is released, AWS will analyze service logs of network traffic processed by these rule groups to identify threat indicators across customers. AWS will use these threat indicators to improve the active threat defense managed rule groups and protect the security of AWS customers and services.
	//
	// > Customers can opt-out of deep threat inspection at any time through the AWS Network Firewall console or API. When customers opt out, AWS Network Firewall will not use the network traffic processed by those customers' active threat defense rule groups for rule group improvement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupreference.html#cfn-networkfirewall-firewallpolicy-statefulrulegroupreference-deepthreatinspection
	//
	DeepThreatInspection interface{} `field:"optional" json:"deepThreatInspection" yaml:"deepThreatInspection"`
	// The action that allows the policy owner to override the behavior of the rule group within a policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupreference.html#cfn-networkfirewall-firewallpolicy-statefulrulegroupreference-override
	//
	Override interface{} `field:"optional" json:"override" yaml:"override"`
	// An integer setting that indicates the order in which to run the stateful rule groups in a single firewall policy.
	//
	// This setting only applies to firewall policies that specify the `STRICT_ORDER` rule order in the stateful engine options settings.
	//
	// Network Firewall evalutes each stateful rule group against a packet starting with the group that has the lowest priority setting. You must ensure that the priority settings are unique within each policy.
	//
	// You can change the priority settings of your rule groups at any time. To make it easier to insert rule groups later, number them so there's a wide range in between, for example use 100, 200, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupreference.html#cfn-networkfirewall-firewallpolicy-statefulrulegroupreference-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The Amazon Resource Name (ARN) of the stateful rule group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-statefulrulegroupreference.html#cfn-networkfirewall-firewallpolicy-statefulrulegroupreference-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

