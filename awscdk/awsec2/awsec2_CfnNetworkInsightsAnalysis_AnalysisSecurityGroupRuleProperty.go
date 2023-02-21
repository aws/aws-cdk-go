package awsec2


// Describes a security group rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSecurityGroupRuleProperty := &analysisSecurityGroupRuleProperty{
//   	cidr: jsii.String("cidr"),
//   	direction: jsii.String("direction"),
//   	portRange: &portRangeProperty{
//   		from: jsii.Number(123),
//   		to: jsii.Number(123),
//   	},
//   	prefixListId: jsii.String("prefixListId"),
//   	protocol: jsii.String("protocol"),
//   	securityGroupId: jsii.String("securityGroupId"),
//   }
//
type CfnNetworkInsightsAnalysis_AnalysisSecurityGroupRuleProperty struct {
	// The IPv4 address range, in CIDR notation.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The direction. The following are the possible values:.
	//
	// - egress
	// - ingress.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// The port range.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// The prefix list ID.
	PrefixListId *string `field:"optional" json:"prefixListId" yaml:"prefixListId"`
	// The protocol name.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The security group ID.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
}

