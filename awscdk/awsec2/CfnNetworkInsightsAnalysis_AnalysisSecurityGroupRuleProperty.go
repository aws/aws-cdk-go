package awsec2


// Describes a security group rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSecurityGroupRuleProperty := &AnalysisSecurityGroupRuleProperty{
//   	Cidr: jsii.String("cidr"),
//   	Direction: jsii.String("direction"),
//   	PortRange: &PortRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   	PrefixListId: jsii.String("prefixListId"),
//   	Protocol: jsii.String("protocol"),
//   	SecurityGroupId: jsii.String("securityGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysissecuritygrouprule.html
//
type CfnNetworkInsightsAnalysis_AnalysisSecurityGroupRuleProperty struct {
	// The IPv4 address range, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysissecuritygrouprule.html#cfn-ec2-networkinsightsanalysis-analysissecuritygrouprule-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The direction. The following are the possible values:.
	//
	// - egress
	// - ingress.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysissecuritygrouprule.html#cfn-ec2-networkinsightsanalysis-analysissecuritygrouprule-direction
	//
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// The port range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysissecuritygrouprule.html#cfn-ec2-networkinsightsanalysis-analysissecuritygrouprule-portrange
	//
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// The prefix list ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysissecuritygrouprule.html#cfn-ec2-networkinsightsanalysis-analysissecuritygrouprule-prefixlistid
	//
	PrefixListId *string `field:"optional" json:"prefixListId" yaml:"prefixListId"`
	// The protocol name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysissecuritygrouprule.html#cfn-ec2-networkinsightsanalysis-analysissecuritygrouprule-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The security group ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysissecuritygrouprule.html#cfn-ec2-networkinsightsanalysis-analysissecuritygrouprule-securitygroupid
	//
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
}

