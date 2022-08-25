package awsec2


// Describes a network access control (ACL) rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisAclRuleProperty := &analysisAclRuleProperty{
//   	cidr: jsii.String("cidr"),
//   	egress: jsii.Boolean(false),
//   	portRange: &portRangeProperty{
//   		from: jsii.Number(123),
//   		to: jsii.Number(123),
//   	},
//   	protocol: jsii.String("protocol"),
//   	ruleAction: jsii.String("ruleAction"),
//   	ruleNumber: jsii.Number(123),
//   }
//
type CfnNetworkInsightsAnalysis_AnalysisAclRuleProperty struct {
	// The IPv4 address range, in CIDR notation.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Indicates whether the rule is an outbound rule.
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// The range of ports.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// The protocol.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether to allow or deny traffic that matches the rule.
	RuleAction *string `field:"optional" json:"ruleAction" yaml:"ruleAction"`
	// The rule number.
	RuleNumber *float64 `field:"optional" json:"ruleNumber" yaml:"ruleNumber"`
}

