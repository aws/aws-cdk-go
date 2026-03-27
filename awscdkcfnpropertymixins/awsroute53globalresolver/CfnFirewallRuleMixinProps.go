package awsroute53globalresolver


// Properties for CfnFirewallRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFirewallRuleMixinProps := &CfnFirewallRuleMixinProps{
//   	Action: jsii.String("action"),
//   	BlockOverrideDnsType: jsii.String("blockOverrideDnsType"),
//   	BlockOverrideDomain: jsii.String("blockOverrideDomain"),
//   	BlockOverrideTtl: jsii.Number(123),
//   	BlockResponse: jsii.String("blockResponse"),
//   	ClientToken: jsii.String("clientToken"),
//   	ConfidenceThreshold: jsii.String("confidenceThreshold"),
//   	Description: jsii.String("description"),
//   	DnsAdvancedProtection: jsii.String("dnsAdvancedProtection"),
//   	DnsViewId: jsii.String("dnsViewId"),
//   	FirewallDomainListId: jsii.String("firewallDomainListId"),
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	QType: jsii.String("qType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html
//
type CfnFirewallRuleMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-blockoverridednstype
	//
	BlockOverrideDnsType *string `field:"optional" json:"blockOverrideDnsType" yaml:"blockOverrideDnsType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-blockoverridedomain
	//
	BlockOverrideDomain *string `field:"optional" json:"blockOverrideDomain" yaml:"blockOverrideDomain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-blockoverridettl
	//
	BlockOverrideTtl *float64 `field:"optional" json:"blockOverrideTtl" yaml:"blockOverrideTtl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-blockresponse
	//
	BlockResponse *string `field:"optional" json:"blockResponse" yaml:"blockResponse"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-confidencethreshold
	//
	ConfidenceThreshold *string `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-dnsadvancedprotection
	//
	DnsAdvancedProtection *string `field:"optional" json:"dnsAdvancedProtection" yaml:"dnsAdvancedProtection"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-dnsviewid
	//
	DnsViewId *string `field:"optional" json:"dnsViewId" yaml:"dnsViewId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-firewalldomainlistid
	//
	FirewallDomainListId *string `field:"optional" json:"firewallDomainListId" yaml:"firewallDomainListId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-firewallrule.html#cfn-route53globalresolver-firewallrule-qtype
	//
	QType *string `field:"optional" json:"qType" yaml:"qType"`
}

