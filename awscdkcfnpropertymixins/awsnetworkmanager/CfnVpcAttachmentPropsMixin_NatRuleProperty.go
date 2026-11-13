package awsnetworkmanager


// A NAT gateway rule on the attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   natRuleProperty := &NatRuleProperty{
//   	MatchSourceCidr: jsii.String("matchSourceCidr"),
//   	NatCidr: jsii.String("natCidr"),
//   	RuleNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-natrule.html
//
type CfnVpcAttachmentPropsMixin_NatRuleProperty struct {
	// The IPv4 CIDR that source traffic is matched against for translation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-natrule.html#cfn-networkmanager-vpcattachment-natrule-matchsourcecidr
	//
	MatchSourceCidr *string `field:"optional" json:"matchSourceCidr" yaml:"matchSourceCidr"`
	// The IPv4 CIDR that matched source traffic is translated into.
	//
	// Must have the same prefix length as MatchSourceCidr.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-natrule.html#cfn-networkmanager-vpcattachment-natrule-natcidr
	//
	NatCidr *string `field:"optional" json:"natCidr" yaml:"natCidr"`
	// The rule number that uniquely identifies the NAT rule on the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-natrule.html#cfn-networkmanager-vpcattachment-natrule-rulenumber
	//
	RuleNumber *float64 `field:"optional" json:"ruleNumber" yaml:"ruleNumber"`
}

