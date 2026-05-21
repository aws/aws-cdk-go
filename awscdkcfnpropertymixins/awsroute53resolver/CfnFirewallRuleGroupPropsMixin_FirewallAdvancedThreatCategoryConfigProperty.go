package awsroute53resolver


// Configuration for an advanced threat category rule type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   firewallAdvancedThreatCategoryConfigProperty := &FirewallAdvancedThreatCategoryConfigProperty{
//   	Category: jsii.String("category"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewalladvancedthreatcategoryconfig.html
//
type CfnFirewallRuleGroupPropsMixin_FirewallAdvancedThreatCategoryConfigProperty struct {
	// The threat category value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewalladvancedthreatcategoryconfig.html#cfn-route53resolver-firewallrulegroup-firewalladvancedthreatcategoryconfig-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
}

