package awsnetworkfirewall


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   route53ResolvedIPValidationProperty := &Route53ResolvedIPValidationProperty{
//   	Action: jsii.String("action"),
//   	BypassDestinationValidationIpRanges: []interface{}{
//   		&AddressProperty{
//   			AddressDefinition: jsii.String("addressDefinition"),
//   		},
//   	},
//   	DnsExpiryExceptionPolicy: jsii.String("dnsExpiryExceptionPolicy"),
//   	DnsFirewallRuleGroupId: jsii.String("dnsFirewallRuleGroupId"),
//   	SourceIpRanges: []interface{}{
//   		&AddressProperty{
//   			AddressDefinition: jsii.String("addressDefinition"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html
//
type CfnFirewallPolicyPropsMixin_Route53ResolvedIPValidationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-bypassdestinationvalidationipranges
	//
	BypassDestinationValidationIpRanges interface{} `field:"optional" json:"bypassDestinationValidationIpRanges" yaml:"bypassDestinationValidationIpRanges"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-dnsexpiryexceptionpolicy
	//
	DnsExpiryExceptionPolicy *string `field:"optional" json:"dnsExpiryExceptionPolicy" yaml:"dnsExpiryExceptionPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-dnsfirewallrulegroupid
	//
	DnsFirewallRuleGroupId *string `field:"optional" json:"dnsFirewallRuleGroupId" yaml:"dnsFirewallRuleGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-firewallpolicy-route53resolvedipvalidation.html#cfn-networkfirewall-firewallpolicy-route53resolvedipvalidation-sourceipranges
	//
	SourceIpRanges interface{} `field:"optional" json:"sourceIpRanges" yaml:"sourceIpRanges"`
}

