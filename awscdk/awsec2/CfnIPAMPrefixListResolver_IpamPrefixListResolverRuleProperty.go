package awsec2


// CIDR selection rules define the business logic for selecting CIDRs from IPAM.
//
// If a CIDR matches any of the rules, it will be included. If a rule has multiple conditions, the CIDR has to match every condition of that rule. You can create a prefix list resolver without rules, but you'll need to add at least one rule before it can actually automate your prefix list updates.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipamPrefixListResolverRuleProperty := &IpamPrefixListResolverRuleProperty{
//   	RuleType: jsii.String("ruleType"),
//
//   	// the properties below are optional
//   	Conditions: []interface{}{
//   		&IpamPrefixListResolverRuleConditionProperty{
//   			Operation: jsii.String("operation"),
//
//   			// the properties below are optional
//   			Cidr: jsii.String("cidr"),
//   			IpamPoolId: jsii.String("ipamPoolId"),
//   			ResourceId: jsii.String("resourceId"),
//   			ResourceOwner: jsii.String("resourceOwner"),
//   			ResourceRegion: jsii.String("resourceRegion"),
//   			ResourceTag: &CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	IpamScopeId: jsii.String("ipamScopeId"),
//   	ResourceType: jsii.String("resourceType"),
//   	StaticCidr: jsii.String("staticCidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule.html
//
type CfnIPAMPrefixListResolver_IpamPrefixListResolverRuleProperty struct {
	// There are three rule types: (1) Static CIDR: A fixed list of CIDRs that don't change (like a manual list replicated across Regions).
	//
	// (2) IPAM pool CIDR: CIDRs from specific IPAM pools (like all CIDRs from your IPAM production pool).  (3) IPAM resource CIDR: CIDRs for AWS resources like VPCs, subnets, and EIPs within a specific IPAM scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule-ruletype
	//
	RuleType *string `field:"required" json:"ruleType" yaml:"ruleType"`
	// Two of the rule types allow you to add conditions to the rules.
	//
	// (1) For IPAM Pool CIDR rules, you can specify an ipamPoolId; if not specified, the rule will apply to all IPAM Pool CIDRs in the scope.  (2) For IPAM Resource CIDR rules, you can specify resourceId, resourceOwner, resourceRegion, cidr, or resourceTag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// This rule will only match resources that are in this IPAM Scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule-ipamscopeid
	//
	IpamScopeId *string `field:"optional" json:"ipamScopeId" yaml:"ipamScopeId"`
	// The resourceType property only applies to ipam-resource-cidr rules;
	//
	// this property specifies what type of resources this rule will apply to, such as VPCs or Subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// A fixed CIDR that doesn't change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrule-staticcidr
	//
	StaticCidr *string `field:"optional" json:"staticCidr" yaml:"staticCidr"`
}

