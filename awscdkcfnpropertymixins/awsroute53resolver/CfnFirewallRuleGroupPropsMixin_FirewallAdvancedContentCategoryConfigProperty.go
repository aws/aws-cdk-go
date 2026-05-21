package awsroute53resolver


// Configuration for an advanced content category rule type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   firewallAdvancedContentCategoryConfigProperty := &FirewallAdvancedContentCategoryConfigProperty{
//   	Category: jsii.String("category"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewalladvancedcontentcategoryconfig.html
//
type CfnFirewallRuleGroupPropsMixin_FirewallAdvancedContentCategoryConfigProperty struct {
	// The content category value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-firewallrulegroup-firewalladvancedcontentcategoryconfig.html#cfn-route53resolver-firewallrulegroup-firewalladvancedcontentcategoryconfig-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
}

