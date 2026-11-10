package awsec2


// A static 1:1 NAT rule.
//
// Traffic whose source address falls within MatchSourceCidr is translated to the corresponding address in NatCidr.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   natRuleProperty := &NatRuleProperty{
//   	MatchSourceCidr: jsii.String("matchSourceCidr"),
//   	NatCidr: jsii.String("natCidr"),
//   	RuleNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachmentnatconfiguration-natrule.html
//
type CfnTransitGatewayAttachmentNatConfiguration_NatRuleProperty struct {
	// The source CIDR that the rule matches on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachmentnatconfiguration-natrule.html#cfn-ec2-transitgatewayattachmentnatconfiguration-natrule-matchsourcecidr
	//
	MatchSourceCidr *string `field:"required" json:"matchSourceCidr" yaml:"matchSourceCidr"`
	// The CIDR that the matched source addresses are translated to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachmentnatconfiguration-natrule.html#cfn-ec2-transitgatewayattachmentnatconfiguration-natrule-natcidr
	//
	NatCidr *string `field:"required" json:"natCidr" yaml:"natCidr"`
	// The rule number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachmentnatconfiguration-natrule.html#cfn-ec2-transitgatewayattachmentnatconfiguration-natrule-rulenumber
	//
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
}

