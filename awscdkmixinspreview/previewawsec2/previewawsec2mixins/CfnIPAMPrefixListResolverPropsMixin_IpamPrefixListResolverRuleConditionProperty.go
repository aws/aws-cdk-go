package previewawsec2mixins


// Two of the rule types allow you to add conditions to the rules.
//
// (1) For IPAM Pool CIDR rules, you can specify an ipamPoolId; if not specified, the rule will apply to all IPAM Pool CIDRs in the scope.  (2) For IPAM Resource CIDR rules, you can specify resourceId, resourceOwner, resourceRegion, cidr, or resourceTag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ipamPrefixListResolverRuleConditionProperty := &IpamPrefixListResolverRuleConditionProperty{
//   	Cidr: jsii.String("cidr"),
//   	IpamPoolId: jsii.String("ipamPoolId"),
//   	Operation: jsii.String("operation"),
//   	ResourceId: jsii.String("resourceId"),
//   	ResourceOwner: jsii.String("resourceOwner"),
//   	ResourceRegion: jsii.String("resourceRegion"),
//   	ResourceTag: &CfnTag{
//   		Key: jsii.String("key"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html
//
type CfnIPAMPrefixListResolverPropsMixin_IpamPrefixListResolverRuleConditionProperty struct {
	// Condition for the IPAM Resource CIDR rule type.
	//
	// CIDR (like 10.24.34.0/23).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Condition for the IPAM Pool CIDR rule type.
	//
	// If not chosen, the resolver applies to all IPAM Pool CIDRs in the scope.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition-ipampoolid
	//
	IpamPoolId *string `field:"optional" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Equals, Not equals, or Subnet Of.
	//
	// The subnet-of operation only applies to cidr conditions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition-operation
	//
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Condition for the IPAM Resource CIDR rule type.
	//
	// The unique ID of a resource (like vpc-1234567890abcdef0).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Condition for the IPAM Resource CIDR rule type.
	//
	// Resource owner (like 111122223333).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition-resourceowner
	//
	ResourceOwner *string `field:"optional" json:"resourceOwner" yaml:"resourceOwner"`
	// Condition for the IPAM Resource CIDR rule type.
	//
	// Resource region (like us-east-1).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition-resourceregion
	//
	ResourceRegion *string `field:"optional" json:"resourceRegion" yaml:"resourceRegion"`
	// A key-value pair to associate with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition.html#cfn-ec2-ipamprefixlistresolver-ipamprefixlistresolverrulecondition-resourcetag
	//
	ResourceTag interface{} `field:"optional" json:"resourceTag" yaml:"resourceTag"`
}

