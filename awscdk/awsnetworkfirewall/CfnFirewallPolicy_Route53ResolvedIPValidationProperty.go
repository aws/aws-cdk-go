package awsnetworkfirewall


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   route53ResolvedIPValidationProperty := &Route53ResolvedIPValidationProperty{
//   	Action: jsii.String("action"),
//   	DnsExpiryExceptionPolicy: jsii.String("dnsExpiryExceptionPolicy"),
//   	DnsFirewallRuleGroupId: jsii.String("dnsFirewallRuleGroupId"),
//
//   	// the properties below are optional
//   	BypassDestinationValidationIpRanges: []interface{}{
//   		&AddressProperty{
//   			AddressDefinition: jsii.String("addressDefinition"),
//   		},
//   	},
//   	SourceIpRanges: []interface{}{
//   		&AddressProperty{
//   			AddressDefinition: jsii.String("addressDefinition"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html
//
type CfnFirewallPolicy_Route53ResolvedIPValidationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-dnsexpiryexceptionpolicy
	//
	DnsExpiryExceptionPolicy *string `field:"required" json:"dnsExpiryExceptionPolicy" yaml:"dnsExpiryExceptionPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-dnsfirewallrulegroupid
	//
	DnsFirewallRuleGroupId *string `field:"required" json:"dnsFirewallRuleGroupId" yaml:"dnsFirewallRuleGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-bypassdestinationvalidationipranges
	//
	BypassDestinationValidationIpRanges interface{} `field:"optional" json:"bypassDestinationValidationIpRanges" yaml:"bypassDestinationValidationIpRanges"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-sourceipranges
	//
	SourceIpRanges interface{} `field:"optional" json:"sourceIpRanges" yaml:"sourceIpRanges"`
}

