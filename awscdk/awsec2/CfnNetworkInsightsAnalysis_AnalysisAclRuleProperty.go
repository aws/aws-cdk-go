package awsec2


// Describes a network access control (ACL) rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisAclRuleProperty := &AnalysisAclRuleProperty{
//   	Cidr: jsii.String("cidr"),
//   	Egress: jsii.Boolean(false),
//   	PortRange: &PortRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   	Protocol: jsii.String("protocol"),
//   	RuleAction: jsii.String("ruleAction"),
//   	RuleNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisaclrule.html
//
type CfnNetworkInsightsAnalysis_AnalysisAclRuleProperty struct {
	// The IPv4 address range, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisaclrule.html#cfn-ec2-networkinsightsanalysis-analysisaclrule-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Indicates whether the rule is an outbound rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisaclrule.html#cfn-ec2-networkinsightsanalysis-analysisaclrule-egress
	//
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// The range of ports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisaclrule.html#cfn-ec2-networkinsightsanalysis-analysisaclrule-portrange
	//
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// The protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisaclrule.html#cfn-ec2-networkinsightsanalysis-analysisaclrule-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether to allow or deny traffic that matches the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisaclrule.html#cfn-ec2-networkinsightsanalysis-analysisaclrule-ruleaction
	//
	RuleAction *string `field:"optional" json:"ruleAction" yaml:"ruleAction"`
	// The rule number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysisaclrule.html#cfn-ec2-networkinsightsanalysis-analysisaclrule-rulenumber
	//
	RuleNumber *float64 `field:"optional" json:"ruleNumber" yaml:"ruleNumber"`
}

