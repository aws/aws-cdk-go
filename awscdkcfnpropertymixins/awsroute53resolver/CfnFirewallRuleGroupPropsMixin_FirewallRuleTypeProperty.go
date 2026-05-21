package awsroute53resolver


// Firewall rule type union.
//
// Exactly one member must be set. Mutually exclusive with FirewallDomainListId and DnsThreatProtection/ConfidenceThreshold.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   firewallRuleTypeProperty := &FirewallRuleTypeProperty{
//   	FirewallAdvancedContentCategory: &FirewallAdvancedContentCategoryConfigProperty{
//   		Category: jsii.String("category"),
//   	},
//   	FirewallAdvancedThreatCategory: &FirewallAdvancedThreatCategoryConfigProperty{
//   		Category: jsii.String("category"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallruletype.html
//
type CfnFirewallRuleGroupPropsMixin_FirewallRuleTypeProperty struct {
	// Configuration for an advanced content category rule type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallruletype.html#cfn-route53resolver-firewallrulegroup-firewallruletype-firewalladvancedcontentcategory
	//
	FirewallAdvancedContentCategory interface{} `field:"optional" json:"firewallAdvancedContentCategory" yaml:"firewallAdvancedContentCategory"`
	// Configuration for an advanced threat category rule type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewallruletype.html#cfn-route53resolver-firewallrulegroup-firewallruletype-firewalladvancedthreatcategory
	//
	FirewallAdvancedThreatCategory interface{} `field:"optional" json:"firewallAdvancedThreatCategory" yaml:"firewallAdvancedThreatCategory"`
}

