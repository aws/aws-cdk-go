package awsnetworkmanager


// A NAT gateway rule on the attachment.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-directconnectgatewayattachment-natrule.html
//
type CfnDirectConnectGatewayAttachment_NatRuleProperty struct {
	// The IPv4 CIDR that source traffic is matched against for translation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-directconnectgatewayattachment-natrule.html#cfn-networkmanager-directconnectgatewayattachment-natrule-matchsourcecidr
	//
	MatchSourceCidr *string `field:"required" json:"matchSourceCidr" yaml:"matchSourceCidr"`
	// The IPv4 CIDR that matched source traffic is translated into.
	//
	// Must have the same prefix length as MatchSourceCidr.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-directconnectgatewayattachment-natrule.html#cfn-networkmanager-directconnectgatewayattachment-natrule-natcidr
	//
	NatCidr *string `field:"required" json:"natCidr" yaml:"natCidr"`
	// The rule number that uniquely identifies the NAT rule on the attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-directconnectgatewayattachment-natrule.html#cfn-networkmanager-directconnectgatewayattachment-natrule-rulenumber
	//
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
}

